package client

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Test_Unite_Convert(t *testing.T) {
	t.Run("TestServer", func(t *testing.T) {
		app := gin.New()
		app.GET("/metrics", PromHandler(promhttp.Handler()))
		app.POST("/", NewGinMethodConvert(GetDefaultConfPool()).ConvertGinHandler)
		_ = app.Run(":8545")
	})
}
