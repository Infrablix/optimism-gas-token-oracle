// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import { ISemver } from "src/universal/interfaces/ISemver.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

/// @title CustomGasTokenPriceOracle
/// @notice This contract holds the price of a custom fee token in ETH. The price is fetched by the L1 Cost function
///         and is used to calculate the L1 portion of the L2 transaction cost in the custom fee token denomination.
contract CustomGasTokenPriceOracle is ISemver, Ownable {
    /// @notice Semantic version.
    /// @custom:semver 0.1.0
    string public constant version = "0.1.0";

    /// @notice The price of the custom fee token in ETH.
    uint256 public priceInEth;

    /// @notice The block timestamp of the last price update.
    uint256 public lastUpdateTimestamp;

    /// @notice Update the price of the custom fee token in ETH, and sets the last update timestamp as the current
    ///         block timestamp. Can only be called by the owner.
    function update(uint256 _priceInEth) external onlyOwner {
        priceInEth = _priceInEth;
        lastUpdateTimestamp = uint64(block.timestamp);
    }

    /// @notice Fetche the price of the custom fee token in ETH and the last update timestamp.
    function getPrice() external view returns (uint256 _priceInEth, uint256 _lastUpdateTimestamp) {
        return (priceInEth, lastUpdateTimestamp);
    }
}
