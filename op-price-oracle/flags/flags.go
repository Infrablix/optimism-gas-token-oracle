package flags

import (
	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	oprpc "github.com/ethereum-optimism/optimism/op-service/rpc"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli/v2"
)

const EnvVarPrefix = "OP_PRICE_ORACLE"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var Flags []cli.Flag

var (
	// Required flags
	L1EthRpcFlag = &cli.StringFlag{
		Name:    "l1-eth-rpc",
		Usage:   "HTTP provider URL for L1",
		EnvVars: prefixEnvVars("L1_ETH_RPC"),
	}
	L2EthRpcFlag = &cli.StringFlag{
		Name:    "l2-eth-rpc",
		Usage:   "HTTP provider URL for L2 execution engine. A comma-separated list enables the active L2 endpoint provider. Such a list needs to match the number of rollup-rpcs provided.",
		EnvVars: prefixEnvVars("L2_ETH_RPC"),
	}
	// Optional flags
	MinPriceChangeFlag = &cli.StringFlag{
		Name:    "min-price-change-percent",
		Usage:   "Min price change percent to trigger price update tx, where 1 == 1% change",
		EnvVars: prefixEnvVars("MIN_PRICE_CHANGE_PERCENT"),
	}
	MaxElapsedTimeFlag = &cli.DurationFlag{
		Name:    "max-elapsed-time",
		Usage:   "Max time elapsed to trigger price update tx, even if min price change percent not reached",
		EnvVars: prefixEnvVars("MAX_ELAPSED_TIME"),
	}
)

func init() {
	Flags = append(Flags, L1EthRpcFlag)
	Flags = append(Flags, L2EthRpcFlag)
	Flags = append(Flags, MinPriceChangeFlag)
	Flags = append(Flags, MaxElapsedTimeFlag)

	Flags = append(Flags, oprpc.CLIFlags(EnvVarPrefix)...)
	Flags = append(Flags, txmgr.CLIFlags(EnvVarPrefix)...)
	Flags = append(Flags, oplog.CLIFlags(EnvVarPrefix)...)
	Flags = append(Flags, opmetrics.CLIFlags(EnvVarPrefix)...)
}
