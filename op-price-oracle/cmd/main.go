package main

import (
	"context"
	"os"

	"github.com/ethereum-optimism/optimism/op-price-oracle/flags"
	"github.com/ethereum-optimism/optimism/op-price-oracle/metrics"
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum-optimism/optimism/op-service/ctxinterrupt"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/metrics/doc"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Name = "op-price-oracle"
	app.Usage = "Custom Gas Token Price Oracle"
	app.Description = "Fetches custom gas token price in ETH and submits price update txs on L2"
	app.Action = cliapp.LifecycleCmd(oracle.Main())
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.NewSubcommands(metrics.NewMetrics("")),
		},
	}

	ctx := ctxinterrupt.WithSignalWaiterMain(context.Background())
	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
