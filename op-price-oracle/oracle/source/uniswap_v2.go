package source

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/source/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var type112Max = new(big.Int).Exp(big.NewInt(2), big.NewInt(112), nil)
var type144Max = new(big.Int).Exp(big.NewInt(2), big.NewInt(144), nil)

type UniswapV2PriceSource struct {
	rpc        *ethclient.Client
	contract   *bindings.UniswapV2PairCaller
	quoteAsset common.Address
	blockRange uint64

	baseScalar *big.Int
	token0     common.Address
	token1     common.Address
}

func NewUniswapV2PriceSource(rpcUrl string, address common.Address, quoteAsset common.Address, blockRange uint64) (*UniswapV2PriceSource, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	caller, err := bindings.NewUniswapV2PairCaller(address, client)
	if err != nil {
		return nil, err
	}

	opts := new(bind.CallOpts)
	opts.Context = context.Background()
	token0, err := caller.Token0(opts)
	if err != nil {
		return nil, err
	}

	token1, err := caller.Token1(opts)
	if err != nil {
		return nil, err
	}
	var baseScalar *big.Int
	if token0 == quoteAsset {
		erc20, err := bindings.NewERC20Caller(token1, client)
		if err != nil {
			return nil, err
		}
		decimals, err := erc20.Decimals(nil)
		if err != nil {
			return nil, err
		}

		baseScalar = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	} else {
		erc20, err := bindings.NewERC20Caller(token0, client)
		if err != nil {
			return nil, err
		}
		decimals, err := erc20.Decimals(nil)
		if err != nil {
			return nil, err
		}

		baseScalar = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	}

	return &UniswapV2PriceSource{
		client,
		caller,
		quoteAsset,
		blockRange,
		baseScalar,
		token0,
		token1,
	}, nil

}

func (s *UniswapV2PriceSource) GetPriceInEth(ctx context.Context) (*big.Int, error) {
	twap, err := s.getTwap(ctx)
	if err != nil {
		return nil, err
	}

	quoteAmountQ112 := new(big.Int).Mul(twap, s.baseScalar)
	return new(big.Int).Div(quoteAmountQ112, type112Max), nil
}

func (s *UniswapV2PriceSource) getTwap(ctx context.Context) (*big.Int, error) {
	endBlock, err := s.rpc.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	startBlock, err := s.rpc.HeaderByNumber(ctx, new(big.Int).SetUint64(endBlock.Number.Uint64()-s.blockRange))
	if err != nil {
		return nil, err
	}

	startPriceCumulative, err := s.getPriceCumulative(ctx, startBlock)
	if err != nil {
		return nil, err
	}

	endPriceCumulative, err := s.getPriceCumulative(ctx, endBlock)
	if err != nil {
		return nil, err
	}

	priceCumulativeDiff := new(big.Int).Sub(endPriceCumulative, startPriceCumulative)
	timeDiff := new(big.Int).Sub(new(big.Int).SetUint64(endBlock.Time), new(big.Int).SetUint64(startBlock.Time))

	return new(big.Int).Div(priceCumulativeDiff, timeDiff), nil
}

func (s *UniswapV2PriceSource) getPriceCumulative(ctx context.Context, block *types.Header) (*big.Int, error) {
	opts := callOptions(ctx, block)
	reserves, err := s.contract.GetReserves(opts)
	if err != nil {
		return nil, err
	}

	var price *big.Int
	if s.token0 == s.quoteAsset {
		// price1Cumulative = token0 / token1
		price, err = s.getPrice1Cumulative(opts)
	} else {
		// price0Cumulative = token1 / token0
		price, err = s.getPrice0Cumulative(opts)
	}

	if err != nil {
		return nil, err
	}

	if block.Time == uint64(reserves.BlockTimestampLast) {
		return price, nil
	}

	// interpolate the cumulative price
	timeDiff := new(big.Int).Sub(new(big.Int).SetUint64(block.Time), new(big.Int).SetUint64(uint64(reserves.BlockTimestampLast)))
	toAdd := new(big.Int).Mul(fixedPointFraction(reserves.Reserve1, reserves.Reserve0), timeDiff)
	if s.token0 == s.quoteAsset {
		toAdd = new(big.Int).Mul(fixedPointFraction(reserves.Reserve0, reserves.Reserve1), timeDiff)
	}

	return new(big.Int).Add(price, toAdd), nil
}

func (s *UniswapV2PriceSource) getPrice0Cumulative(opts *bind.CallOpts) (*big.Int, error) {
	priceCumulative, err := s.contract.Price0CumulativeLast(opts)
	if err != nil {
		return nil, err
	}

	return priceCumulative, nil
}

func (s *UniswapV2PriceSource) getPrice1Cumulative(opts *bind.CallOpts) (*big.Int, error) {
	priceCumulative, err := s.contract.Price1CumulativeLast(opts)
	if err != nil {
		return nil, err
	}

	return priceCumulative, nil
}

// returns a UQ112x112 which represents the ratio of the numerator to the denominator
func fixedPointFraction(numerator *big.Int, denominator *big.Int) *big.Int {
	if numerator.Cmp(common.Big0) == 0 {
		return big.NewInt(0)
	}

	if numerator.Cmp(type144Max) <= 0 {
		result := new(big.Int).Mul(numerator, type112Max)
		return result.Div(result, denominator)
	}

	result := new(big.Int).Mul(numerator, type112Max)
	return result.Div(result, denominator)
}

func callOptions(ctx context.Context, block *types.Header) *bind.CallOpts {
	opts := new(bind.CallOpts)
	opts.Context = ctx
	opts.BlockNumber = block.Number
	return opts
}
