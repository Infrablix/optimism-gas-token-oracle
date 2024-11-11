package source

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/source/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

// based on https://docs.uniswap.org/sdk/v3/guides/advanced/price-oracle
type UniswapV3PriceSource struct {
	rpc        *ethclient.Client
	contract   *bindings.UniswapV3PoolCaller
	quoteAsset common.Address
	blockRange uint64

	baseAmount *entities.CurrencyAmount
	token0     *entities.Token
	token1     *entities.Token
}

func NewUniswapV3PriceSource(rpcUrl string, address common.Address, quoteAsset common.Address, blockRange uint64) (*UniswapV3PriceSource, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	caller, err := bindings.NewUniswapV3PoolCaller(address, client)
	if err != nil {
		return nil, err
	}

	opts := new(bind.CallOpts)
	opts.Context = context.Background()
	token0, err := caller.Token0(opts)
	if err != nil {
		return nil, err
	}

	token0ERC20, err := bindings.NewERC20Caller(token0, client)
	if err != nil {
		return nil, err
	}

	tokenODecimals, err := token0ERC20.Decimals(opts)
	if err != nil {
		return nil, err
	}

	token1, err := caller.Token1(opts)
	if err != nil {
		return nil, err
	}

	token1ERC20, err := bindings.NewERC20Caller(token1, client)
	if err != nil {
		return nil, err
	}

	token1Decimals, err := token1ERC20.Decimals(opts)
	if err != nil {
		return nil, err
	}

	token0Entity := entities.NewToken(1, token0, uint(tokenODecimals), "", "")
	token1Entity := entities.NewToken(1, token1, uint(token1Decimals), "", "")

	var baseAmount *entities.CurrencyAmount
	if quoteAsset == token0 {
		baseAmount = entities.FromRawAmount(token1Entity, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token1Entity.Decimals())), nil))
	} else {
		baseAmount = entities.FromRawAmount(token0Entity, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token0Entity.Decimals())), nil))
	}

	return &UniswapV3PriceSource{
		client,
		caller,
		quoteAsset,
		blockRange,
		baseAmount,
		token0Entity,
		token1Entity,
	}, nil

}

func (s *UniswapV3PriceSource) GetPriceInEth(ctx context.Context) (*big.Int, error) {
	endBlock, err := s.rpc.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	startBlock, err := s.rpc.HeaderByNumber(ctx, new(big.Int).SetUint64(endBlock.Number.Uint64()-s.blockRange))
	if err != nil {
		return nil, err
	}

	timestamps := []uint32{0, uint32(endBlock.Time) - uint32(startBlock.Time)}
	opts := new(bind.CallOpts)
	opts.Context = ctx
	opts.BlockNumber = endBlock.Number

	observations, err := s.contract.Observe(opts, timestamps)
	if err != nil {
		return nil, err
	}

	tickDiffCumulative := new(big.Int).Sub(observations.TickCumulatives[0], observations.TickCumulatives[1])
	secondsBetweenTicks := new(big.Int).SetUint64(uint64(timestamps[1]))

	averageTick := new(big.Int).Div(tickDiffCumulative, secondsBetweenTicks)
	baseAsset := s.token0
	quoteAsset := s.token1
	if s.quoteAsset == baseAsset.Address {
		baseAsset = s.token1
		quoteAsset = s.token0
	}

	priceX192, err := utils.TickToPrice(baseAsset, quoteAsset, int(averageTick.Int64()))
	if err != nil {
		return nil, err
	}

	price, err := priceX192.Quote(s.baseAmount)
	if err != nil {
		return nil, err
	}

	return price.Quotient(), nil
}
