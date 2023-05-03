package client

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"testing"
)

func Test_Unite_Convert(t *testing.T) {
	t.Run("", func(t *testing.T) {
		app := gin.New()
		app.GET("/metrics", PromHandler(promhttp.Handler()))
		app.POST("/eth_proxy/:chain_id", NewGinMethodConvert(GetDefaultConfPool()).ConvertGinHandler)
		_ = app.Run(":20003")
	})
}
