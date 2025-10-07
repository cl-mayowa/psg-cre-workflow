//go:build wasip1

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"psg-digital/contracts/evm/src/generated/fund_data"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

// EVMConfig holds per-chain configuration.
type EVMConfig struct {
	FundDataAddress string `json:"fundDataAddress"`
	ChainSelector   uint64 `json:"chainSelector"`
	GasLimit        uint64 `json:"gasLimit"`
}

type Config struct {
	EVMs []EVMConfig `json:"evms"`
}

type HTTPTriggerPayload struct {
	ExecutionTime  time.Time `json:"executionTime"`
	ProofOfReserve uint64    `json:"proofOfReserve"`
	NetAssetValue  uint64    `json:"netAssetValue"`
	Rating         uint64    `json:"rating"`
	Open           bool      `json:"open"`
	CurrencyType   string    `json:"currencyType"`
}

func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	httpTriggerCfg := &http.Config{}

	return cre.Workflow[*Config]{
		cre.Handler(
			http.Trigger(httpTriggerCfg),
			onHTTPTrigger,
		),
	}, nil
}

func onHTTPTrigger(config *Config, runtime cre.Runtime, payload *http.Payload) (string, error) {
	logger := runtime.Logger()
	logger.Info("Raw HTTP trigger received")

	// If there's no input, read current fund data from contract
	if len(payload.Input) == 0 {
		logger.Info("No payload provided, reading current fund data from contract")
		fundInfo, err := readFundDataFromContract(config, runtime)
		if err != nil {
			return "", fmt.Errorf("failed to read fund data from contract: %w", err)
		}

		// Return the current fund data as JSON
		fundDataJSON, err := json.Marshal(fundInfo)
		if err != nil {
			return "", fmt.Errorf("failed to marshal fund data: %w", err)
		}
		return string(fundDataJSON), nil
	}

	// Log the raw JSON for debugging
	logger.Info("Payload bytes", "payloadBytes", string(payload.Input))

	//Unmarshal raw JSON bytes directly into the struct
	var req HTTPTriggerPayload
	if err := json.Unmarshal(payload.Input, &req); err != nil {
		logger.Error("failed to unmarshal http trigger payload", "err", err)
		return "", err
	}

	// Provide a sensible default if the field is missing/zero
	if req.ExecutionTime.IsZero() {
		req.ExecutionTime = time.Now().UTC()
	}

	logger.Info("Parsed HTTP trigger received", "payload", req)

	// Update fund data contract
	if err := updateFundDataContract(config, runtime, &req); err != nil {
		return "", fmt.Errorf("failed to update fund data contract: %w", err)
	}

	return "Fund data updated successfully", nil
}

func readFundDataFromContract(config *Config, runtime cre.Runtime) (map[string]interface{}, error) {
	evmCfg := config.EVMs[0]
	logger := runtime.Logger()

	logger.Info("Reading fund data from contract", "address", evmCfg.FundDataAddress)

	evmClient := &evm.Client{
		ChainSelector: evmCfg.ChainSelector,
	}

	fundDataAddress, err := hexToBytes(evmCfg.FundDataAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to decode fund data address %s: %w", evmCfg.FundDataAddress, err)
	}

	fundDataContract, err := fund_data.NewFundData(evmClient, fundDataAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create fund data contract: %w", err)
	}

	// Read individual fields from the contract using getter functions
	porResp, err := fundDataContract.GetProofOfReserve(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read proof of reserve", "error", err)
		return nil, fmt.Errorf("failed to read proof of reserve: %w", err)
	}

	navResp, err := fundDataContract.GetNetAssetValue(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read net asset value", "error", err)
		return nil, fmt.Errorf("failed to read net asset value: %w", err)
	}

	lastUpdateResp, err := fundDataContract.GetLastUpdate(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read last update", "error", err)
		return nil, fmt.Errorf("failed to read last update: %w", err)
	}

	ratingResp, err := fundDataContract.GetRating(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read rating", "error", err)
		return nil, fmt.Errorf("failed to read rating: %w", err)
	}

	openResp, err := fundDataContract.IsOpen(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read open status", "error", err)
		return nil, fmt.Errorf("failed to read open status: %w", err)
	}

	currencyResp, err := fundDataContract.GetCurrencyType(runtime, big.NewInt(8771643)).Await()
	if err != nil {
		logger.Error("Failed to read currency type", "error", err)
		return nil, fmt.Errorf("failed to read currency type: %w", err)
	}

	// Convert responses to appropriate types
	proofOfReserve := new(big.Int).SetBytes(porResp.GetData())
	netAssetValue := new(big.Int).SetBytes(navResp.GetData())
	lastUpdate := new(big.Int).SetBytes(lastUpdateResp.GetData())
	rating := new(big.Int).SetBytes(ratingResp.GetData())

	// For bool, check if the first byte is non-zero
	open := len(openResp.GetData()) > 0 && openResp.GetData()[len(openResp.GetData())-1] != 0

	// Currency type is a string, decode from bytes
	currencyType := string(currencyResp.GetData())

	fundData := map[string]interface{}{
		"proofOfReserve": proofOfReserve.String(),
		"netAssetValue":  netAssetValue.String(),
		"lastUpdate":     lastUpdate.Int64(),
		"rating":         rating.Uint64(),
		"open":           open,
		"currencyType":   currencyType,
	}

	logger.Info("Fund data read successfully", "fundData", fundData)

	return fundData, nil
}

func updateFundDataContract(config *Config, runtime cre.Runtime, payload *HTTPTriggerPayload) error {
	evmCfg := config.EVMs[0]
	logger := runtime.Logger()

	logger.Info("Updating fund data contract",
		"proofOfReserve", payload.ProofOfReserve,
		"netAssetValue", payload.NetAssetValue,
		"rating", payload.Rating,
		"open", payload.Open,
		"currencyType", payload.CurrencyType,
		"executionTime", payload.ExecutionTime,
	)

	evmClient := &evm.Client{
		ChainSelector: evmCfg.ChainSelector,
	}

	fundDataAddress, err := hexToBytes(evmCfg.FundDataAddress)
	if err != nil {
		return fmt.Errorf("failed to decode fund data address %s: %w", evmCfg.FundDataAddress, err)
	}

	fundDataContract, err := fund_data.NewFundData(evmClient, fundDataAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to create fund data contract: %w", err)
	}

	// Convert execution time to Unix timestamp
	lastUpdate := big.NewInt(payload.ExecutionTime.Unix())

	// Create FundInfo struct matching the Solidity contract
	fundInfo := fund_data.FundInfo{
		ProofOfReserve: big.NewInt(int64(payload.ProofOfReserve)),
		NetAssetValue:  big.NewInt(int64(payload.NetAssetValue)),
		LastUpdate:     lastUpdate,
		Rating:         big.NewInt(int64(payload.Rating)),
		Open:           payload.Open,
		CurrencyType:   payload.CurrencyType,
	}

	logger.Info("Writing report to fund data contract", "fundInfo", fundInfo)

	resp, err := fundDataContract.WriteReportFromFundInfo(runtime, fundInfo, nil).Await()
	if err != nil {
		logger.Error("WriteReport await failed", "error", err, "errorType", fmt.Sprintf("%T", err))
		return fmt.Errorf("failed to write report: %w", err)
	}

	logger.Info("Write report succeeded", "response", resp)
	logger.Info("Write report transaction succeeded", "txHash", common.BytesToHash(resp.TxHash).Hex())

	return nil
}

func hexToBytes(hexStr string) ([]byte, error) {
	if len(hexStr) < 2 || hexStr[:2] != "0x" {
		return nil, fmt.Errorf("invalid hex string: %s", hexStr)
	}
	return hex.DecodeString(hexStr[2:])
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
