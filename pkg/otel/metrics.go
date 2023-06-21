package otel

import (
	otelProm "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricSdk "go.opentelemetry.io/otel/sdk/metric"
)

var MetricsWeb3RequestCounter metric.Int64Counter
var MetricsWeb3RequestHistogram metric.Int64Histogram

func init() {
	exporter, err := otelProm.New()
	if err != nil {
		panic(err)
	}
	provider := metricSdk.NewMeterProvider(metricSdk.WithReader(exporter))
	meter := provider.Meter("Web3 Go")
	m1, err := meter.Int64Counter("web3_abi_call", metric.WithDescription("Web3 Gateway abi call counter"))
	if err != nil {
		panic(err)
	}
	m2, err := meter.Int64Histogram("web3_abi_call", metric.WithDescription("Web3 Gateway abi call hist"))
	if err != nil {
		panic(err)
	}
	MetricsWeb3RequestCounter = m1
	MetricsWeb3RequestHistogram = m2
}
