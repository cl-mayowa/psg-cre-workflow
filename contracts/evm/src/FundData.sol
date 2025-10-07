// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC165} from "./keystone/IERC165.sol";
import {IReceiver} from "./keystone/IReceiver.sol";

contract FundData is IReceiver {
    /// @notice Struct containing all fund data
    struct FundInfo {
        uint256 proofOfReserve;      // Proof of Reserve in base units
        uint256 netAssetValue;        // Net Asset Value in base units
        uint256 lastUpdate;           // Timestamp of last update (business datetime)
        uint256 rating;               // Fund rating (could be encoded as uint)
        bool open;                    // Whether fund is open for investment
        string currencyType;          // Currency type (e.g., "enum")
    }

    /// @notice The current fund information
    FundInfo public fundInfo;

    /// @notice Contract owner/admin
    address public admin;

    /// @notice Events
    event FundDataUpdated(
        uint256 proofOfReserve,
        uint256 netAssetValue,
        uint256 lastUpdate,
        uint256 rating,
        bool open,
        string currencyType
    );

    /// @notice Errors
    error Unauthorized();

    modifier onlyAdmin() {
        if (msg.sender != admin) revert Unauthorized();
        _;
    }

    constructor(address _admin) {
        admin = _admin;
    }

    /// @inheritdoc IReceiver
    /// @notice Receives and processes reports from Chainlink CRE
    function onReport(
        bytes calldata metadata,
        bytes calldata report
    ) external override {
        // Decode the report to get the FundInfo struct
        FundInfo memory newFundInfo = abi.decode(report, (FundInfo));

        // Update the fund information
        _updateFundData(newFundInfo);
    }

    /// @notice Updates the fund data with new information
    /// @param newFundInfo The new fund information to store
    function _updateFundData(FundInfo memory newFundInfo) internal {
        fundInfo = newFundInfo;

        emit FundDataUpdated(
            newFundInfo.proofOfReserve,
            newFundInfo.netAssetValue,
            newFundInfo.lastUpdate,
            newFundInfo.rating,
            newFundInfo.open,
            newFundInfo.currencyType
        );
    }

    /// @inheritdoc IERC165
    function supportsInterface(
        bytes4 interfaceId
    ) public pure override returns (bool) {
        return
            interfaceId == type(IReceiver).interfaceId ||
            interfaceId == type(IERC165).interfaceId;
    }

    // ========== GETTER FUNCTIONS ==========

    /// @notice Get the Proof of Reserve value
    function getProofOfReserve() external view returns (uint256) {
        return fundInfo.proofOfReserve;
    }

    /// @notice Get the Net Asset Value
    function getNetAssetValue() external view returns (uint256) {
        return fundInfo.netAssetValue;
    }

    /// @notice Get the last update timestamp
    function getLastUpdate() external view returns (uint256) {
        return fundInfo.lastUpdate;
    }

    /// @notice Get the fund rating
    function getRating() external view returns (uint256) {
        return fundInfo.rating;
    }

    /// @notice Get whether the fund is open
    function isOpen() external view returns (bool) {
        return fundInfo.open;
    }

    /// @notice Get the currency type
    function getCurrencyType() external view returns (string memory) {
        return fundInfo.currencyType;
    }

    /// @notice Get all fund information at once
    function getFundInfo() external view returns (FundInfo memory) {
        return fundInfo;
    }

    // ========== ADMIN FUNCTIONS ==========

    /// @notice Transfer admin rights
    function transferAdmin(address newAdmin) external onlyAdmin {
        admin = newAdmin;
    }
}