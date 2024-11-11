# Optimism Custom Gas Token Oracle

This is a fork of [Optimism](https://github.com/ethereum-optimism/optimism), with support for custom gas token oracle.
While Optimism supports custom gas tokens, it lacks support for converting DA price to custom gas token denomination on
the L2. This leads to a situation where L2s with custom gas token either overcharge or undercharge for the DA costs of
transactions.

## Changes

### op-price-oracle module

- separate, independent service which fetches oracle prices from L1, and submits price updates to L2
- bindings for uniswap v2 and v3 oracles

### CustomGasTokenPriceOracle contract

- predeploy at address `0x42000000000000000000000000000000000007fE`
- stores latest price of custom gas token in ETH / L1 native asset

### Sequencer attributes derivation

- fetches latest custom gas token price in ETH / L1 native asset
- converts `L1BlockInfo#BaseFee/BlobBaseFee` to L2 custom gas token denomination
