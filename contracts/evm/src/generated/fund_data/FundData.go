// Code generated — DO NOT EDIT.

package fund_data

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
)

var FundDataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"}],\"name\":\"UnauthorizedWorkflowOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proofOfReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netAssetValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"currencyType\",\"type\":\"string\"}],\"name\":\"FundDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proofOfReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netAssetValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"currencyType\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrencyType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proofOfReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netAssetValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"currencyType\",\"type\":\"string\"}],\"internalType\":\"structFundData.FundInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetAssetValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProofOfReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proofOfReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netAssetValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rating\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"currencyType\",\"type\":\"string\"}],\"internalType\":\"structFundData.FundInfo\",\"name\":\"newFundInfo\",\"type\":\"tuple\"}],\"name\":\"updateFundData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRating\",\"type\":\"uint256\"}],\"name\":\"updateRating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Structs
type FundInfo struct {
	ProofOfReserve *big.Int
	NetAssetValue  *big.Int
	LastUpdate     *big.Int
	Rating         *big.Int
	Open           bool
	CurrencyType   string
}

// Contract Method Inputs
type OnReportInput struct {
	Metadata []byte
	Report   []byte
}

type SupportsInterfaceInput struct {
	InterfaceId [4]byte
}

type TransferAdminInput struct {
	NewAdmin common.Address
}

type UpdateFundDataInput struct {
	NewFundInfo FundInfo
}

type UpdateRatingInput struct {
	NewRating *big.Int
}

// Errors
type Unauthorized struct {
}

type UnauthorizedWorkflowOwner struct {
	WorkflowOwner common.Address
}

// Events
type FundDataUpdated struct {
	ProofOfReserve *big.Int
	NetAssetValue  *big.Int
	LastUpdate     *big.Int
	Rating         *big.Int
	Open           bool
	CurrencyType   string
}

// Main Binding Type for FundData
type FundData struct {
	Address []byte
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   FundDataCodec
}

type FundDataCodec interface {
	EncodeAdminMethodCall() ([]byte, error)
	DecodeAdminMethodOutput(data []byte) (common.Address, error)
	EncodeFundInfoMethodCall() ([]byte, error)
	DecodeFundInfoMethodOutput(data []byte) (*big.Int, error)
	EncodeGetCurrencyTypeMethodCall() ([]byte, error)
	DecodeGetCurrencyTypeMethodOutput(data []byte) (string, error)
	EncodeGetFundInfoMethodCall() ([]byte, error)
	DecodeGetFundInfoMethodOutput(data []byte) (FundInfo, error)
	EncodeGetLastUpdateMethodCall() ([]byte, error)
	DecodeGetLastUpdateMethodOutput(data []byte) (*big.Int, error)
	EncodeGetNetAssetValueMethodCall() ([]byte, error)
	DecodeGetNetAssetValueMethodOutput(data []byte) (*big.Int, error)
	EncodeGetProofOfReserveMethodCall() ([]byte, error)
	DecodeGetProofOfReserveMethodOutput(data []byte) (*big.Int, error)
	EncodeGetRatingMethodCall() ([]byte, error)
	DecodeGetRatingMethodOutput(data []byte) (*big.Int, error)
	EncodeIsOpenMethodCall() ([]byte, error)
	DecodeIsOpenMethodOutput(data []byte) (bool, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error)
	DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error)
	EncodeTransferAdminMethodCall(in TransferAdminInput) ([]byte, error)
	EncodeUpdateFundDataMethodCall(in UpdateFundDataInput) ([]byte, error)
	EncodeUpdateRatingMethodCall(in UpdateRatingInput) ([]byte, error)
	EncodeFundInfoStruct(in FundInfo) ([]byte, error)
	FundDataUpdatedLogHash() []byte
	EncodeFundDataUpdatedTopics(evt abi.Event, values []FundDataUpdated) ([]*evm.TopicValues, error)
	DecodeFundDataUpdated(log *evm.Log) (*FundDataUpdated, error)
}

func NewFundData(
	client *evm.Client,
	address []byte,
	options *bindings.ContractInitOptions,
) (*FundData, error) {
	parsed, err := abi.JSON(strings.NewReader(FundDataMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &FundData{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (FundDataCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(FundDataMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeAdminMethodCall() ([]byte, error) {
	return c.abi.Pack("admin")
}

func (c *Codec) DecodeAdminMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["admin"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeFundInfoMethodCall() ([]byte, error) {
	return c.abi.Pack("fundInfo")
}

func (c *Codec) DecodeFundInfoMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["fundInfo"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetCurrencyTypeMethodCall() ([]byte, error) {
	return c.abi.Pack("getCurrencyType")
}

func (c *Codec) DecodeGetCurrencyTypeMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["getCurrencyType"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetFundInfoMethodCall() ([]byte, error) {
	return c.abi.Pack("getFundInfo")
}

func (c *Codec) DecodeGetFundInfoMethodOutput(data []byte) (FundInfo, error) {
	vals, err := c.abi.Methods["getFundInfo"].Outputs.Unpack(data)
	if err != nil {
		return *new(FundInfo), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(FundInfo), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result FundInfo
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(FundInfo), fmt.Errorf("failed to unmarshal to FundInfo: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetLastUpdateMethodCall() ([]byte, error) {
	return c.abi.Pack("getLastUpdate")
}

func (c *Codec) DecodeGetLastUpdateMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getLastUpdate"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetNetAssetValueMethodCall() ([]byte, error) {
	return c.abi.Pack("getNetAssetValue")
}

func (c *Codec) DecodeGetNetAssetValueMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getNetAssetValue"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetProofOfReserveMethodCall() ([]byte, error) {
	return c.abi.Pack("getProofOfReserve")
}

func (c *Codec) DecodeGetProofOfReserveMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getProofOfReserve"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetRatingMethodCall() ([]byte, error) {
	return c.abi.Pack("getRating")
}

func (c *Codec) DecodeGetRatingMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getRating"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeIsOpenMethodCall() ([]byte, error) {
	return c.abi.Pack("isOpen")
}

func (c *Codec) DecodeIsOpenMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["isOpen"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Metadata, in.Report)
}

func (c *Codec) EncodeSupportsInterfaceMethodCall(in SupportsInterfaceInput) ([]byte, error) {
	return c.abi.Pack("supportsInterface", in.InterfaceId)
}

func (c *Codec) DecodeSupportsInterfaceMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["supportsInterface"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTransferAdminMethodCall(in TransferAdminInput) ([]byte, error) {
	return c.abi.Pack("transferAdmin", in.NewAdmin)
}

func (c *Codec) EncodeUpdateFundDataMethodCall(in UpdateFundDataInput) ([]byte, error) {
	return c.abi.Pack("updateFundData", in.NewFundInfo)
}

func (c *Codec) EncodeUpdateRatingMethodCall(in UpdateRatingInput) ([]byte, error) {
	return c.abi.Pack("updateRating", in.NewRating)
}

func (c *Codec) EncodeFundInfoStruct(in FundInfo) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "proofOfReserve", Type: "uint256"},
			{Name: "netAssetValue", Type: "uint256"},
			{Name: "lastUpdate", Type: "uint256"},
			{Name: "rating", Type: "uint256"},
			{Name: "open", Type: "bool"},
			{Name: "currencyType", Type: "string"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for FundInfo: %w", err)
	}
	args := abi.Arguments{
		{Name: "fundInfo", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *Codec) FundDataUpdatedLogHash() []byte {
	return c.abi.Events["FundDataUpdated"].ID.Bytes()
}

func (c *Codec) EncodeFundDataUpdatedTopics(
	evt abi.Event,
	values []FundDataUpdated,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeFundDataUpdated decodes a log into a FundDataUpdated struct.
func (c *Codec) DecodeFundDataUpdated(log *evm.Log) (*FundDataUpdated, error) {
	event := new(FundDataUpdated)
	if err := c.abi.UnpackIntoInterface(event, "FundDataUpdated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["FundDataUpdated"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c FundData) Admin(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeAdminMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) FundInfo(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeFundInfoMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetCurrencyType(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetCurrencyTypeMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetFundInfo(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetFundInfoMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetLastUpdate(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetLastUpdateMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetNetAssetValue(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetNetAssetValueMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetProofOfReserve(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetProofOfReserveMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) GetRating(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetRatingMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) IsOpen(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeIsOpenMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(rpc.FinalizedBlockNumber.Int64())),
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c FundData) WriteReportFromFundInfo(
	runtime cre.Runtime,
	input FundInfo,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeFundInfoStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address,
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c FundData) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address,
		Report:    report,
		GasConfig: gasConfig,
	})
}

// DecodeUnauthorizedError decodes a Unauthorized error from revert data.
func (c *FundData) DecodeUnauthorizedError(data []byte) (*Unauthorized, error) {
	args := c.ABI.Errors["Unauthorized"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &Unauthorized{}, nil
}

// Error implements the error interface for Unauthorized.
func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Unauthorized error:")
}

// DecodeUnauthorizedWorkflowOwnerError decodes a UnauthorizedWorkflowOwner error from revert data.
func (c *FundData) DecodeUnauthorizedWorkflowOwnerError(data []byte) (*UnauthorizedWorkflowOwner, error) {
	args := c.ABI.Errors["UnauthorizedWorkflowOwner"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 values, got %d", len(values))
	}

	workflowOwner, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for workflowOwner in UnauthorizedWorkflowOwner error")
	}

	return &UnauthorizedWorkflowOwner{
		WorkflowOwner: workflowOwner,
	}, nil
}

// Error implements the error interface for UnauthorizedWorkflowOwner.
func (e *UnauthorizedWorkflowOwner) Error() string {
	return fmt.Sprintf("UnauthorizedWorkflowOwner error: workflowOwner=%v;", e.WorkflowOwner)
}

func (c *FundData) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["Unauthorized"].ID.Bytes()[:4]):
		return c.DecodeUnauthorizedError(data)
	case common.Bytes2Hex(c.ABI.Errors["UnauthorizedWorkflowOwner"].ID.Bytes()[:4]):
		return c.DecodeUnauthorizedWorkflowOwnerError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

func (c *FundData) LogTriggerFundDataUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []FundDataUpdated) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["FundDataUpdated"]
	topics, err := c.Codec.EncodeFundDataUpdatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for FundDataUpdated: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *FundData) RegisterLogTrackingFundDataUpdated(runtime cre.Runtime, options *bindings.LogTrackingOptions[FundDataUpdated]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[FundDataUpdated](options)
	topics, err := c.Codec.EncodeFundDataUpdatedTopics(c.ABI.Events["FundDataUpdated"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for FundDataUpdated: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "FundDataUpdated-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.FundDataUpdatedLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *FundData) UnregisterLogTrackingFundDataUpdated(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "FundDataUpdated-" + common.Bytes2Hex(c.Address),
	})
}

func (c *FundData) FilterLogsFundDataUpdated(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.FundDataUpdatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}
