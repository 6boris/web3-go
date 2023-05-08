package client

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	"testing"
)

func Test_Unite_Convert(t *testing.T) {
	t.Run("", func(t *testing.T) {
		exporter, err := prometheus.New()
		if err != nil {
			panic(err)
		}
		meterProvider := metricSdk.NewMeterProvider(metricSdk.WithReader(exporter))
		global.SetMeterProvider(meterProvider)
		app := gin.New()
		app.GET("/metrics", PromHandler(promhttp.Handler()))
		app.POST("/eth_proxy/:chain_id", NewGinMethodConvert(GetDefaultConfPool()).ConvertGinHandler)
		_ = app.Run(":20005")
	})
}
