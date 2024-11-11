package metrics

import (
	"context"
	"io"
	"time"

	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/bindings"
	"github.com/ethereum-optimism/optimism/op-service/clock"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	txmetrics "github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metricer interface {
	RecordUp()

	// Record Tx metrics
	txmetrics.TxMetricer

	UpdaterStatusIncrement(status string)
	StartBalanceMetrics(l log.Logger, client *ethclient.Client, account common.Address) io.Closer
	StartTokenPriceMetrics(l log.Logger, oracle *bindings.CustomGasTokenPriceOracle) io.Closer
}

type Metrics struct {
	namespace string
	registry  *prometheus.Registry
	factory   opmetrics.Factory

	up            prometheus.Gauge
	updaterStatus prometheus.CounterVec

	txmetrics.TxMetrics
}

var _ Metricer = (*Metrics)(nil)

// implements the Registry getter, for metrics HTTP server to hook into
var _ opmetrics.RegistryMetricer = (*Metrics)(nil)

func NewMetrics(procName string) *Metrics {
	if procName == "" {
		procName = "default"
	}
	namespace := "op_price_oracle_" + procName

	registry := opmetrics.NewRegistry()
	factory := opmetrics.With(registry)

	return &Metrics{
		namespace: namespace,
		registry:  registry,
		factory:   factory,

		up: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "1 if the op-price-oracle has finished starting up",
		}),

		updaterStatus: *factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "updater_status",
			Help:      "Metric tracking the status of the price updater",
		}, []string{
			"status",
		}),

		TxMetrics: txmetrics.MakeTxMetrics(namespace, factory),
	}
}

func (m *Metrics) Registry() *prometheus.Registry {
	return m.registry
}

func (m *Metrics) Document() []opmetrics.DocumentedMetric {
	return m.factory.Document()
}

func (m *Metrics) UpdaterStatusIncrement(status string) {
	m.updaterStatus.WithLabelValues(status).Inc()
}

func (m *Metrics) StartBalanceMetrics(l log.Logger, client *ethclient.Client, account common.Address) io.Closer {
	return opmetrics.LaunchBalanceMetrics(l, m.registry, m.namespace, client, account)
}

func (m *Metrics) StartTokenPriceMetrics(l log.Logger, oracle *bindings.CustomGasTokenPriceOracle) io.Closer {
	tokenPriceGauge := promauto.With(m.registry).NewGauge(prometheus.GaugeOpts{
		Namespace: m.namespace,
		Name:      "native_token_price_in_eth",
		Help:      "Native token price in ETH",
	})
	lastUpdateAgeGauge := promauto.With(m.registry).NewGauge(prometheus.GaugeOpts{
		Namespace: m.namespace,
		Name:      "native_token_price_update_age",
		Help:      "Duration since last native token price update",
	})

	return clock.NewLoopFn(
		clock.SystemClock,
		func(ctx context.Context) {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
			defer cancel()

			price, err := oracle.GetPrice(&bind.CallOpts{Context: ctx})
			if err != nil {
				l.Warn("failed to get native token price", "err", err)
				return
			}

			duration := time.Since(time.Unix(price.LastUpdateTimestamp.Int64(), 0))

			tokenPriceGauge.Set(eth.WeiToEther(price.PriceInEth))
			lastUpdateAgeGauge.Set(duration.Seconds())
		},
		func() error {
			log.Info("balance metrics shutting down")
			return nil
		},
		10*time.Second,
	)
}

// RecordUp sets the up metric to 1.
func (m *Metrics) RecordUp() {
	prometheus.MustRegister()
	m.up.Set(1)
}
