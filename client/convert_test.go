package client

import (
	"strconv"
	"testing"

	"github.com/6boris/web3-go/consts"
	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func Test_Convert_Server(t *testing.T) {
	t.Run("Start Server", func(t *testing.T) {
		app := gin.New()
		app.GET("/metrics", PromHandler(promhttp.Handler()))
		app.POST("/evm", NewGinMethodConvert(clientModel.GetDefaultConfPool())._convertGinHandler)
		app.POST("/solana", NewGinMethodConvert(clientModel.GetDefaultConfPool())._convertSolanaHandler)
		_ = app.Run(":8545")
	})
}

func Test_Convert_Unite(t *testing.T) {
	cases := []struct {
		name   string
		url    string
		params interface{}
		expect string
	}{
		{
			"EVM", "http://127.0.0.1:8545/evm",
			map[string]interface{}{"chain_id": 1, "method": consts.EvmMethodChainID, "params": []interface{}{}},
			"EVM",
		},
		{
			"EVM", "http://127.0.0.1:8545/evm",
			map[string]interface{}{"chain_id": 1, "method": consts.EvmMethodSuggestGasTipCap, "params": []interface{}{}},
			"",
		},
		{
			"EVM", "http://127.0.0.1:8545/evm",
			map[string]interface{}{"chain_id": 1, "method": consts.EvmMethodBlockNumber, "params": []interface{}{}},
			"",
		},
		{
			"EVM", "http://127.0.0.1:8545/evm",
			map[string]interface{}{"chain_id": 1, "method": consts.EvmMethodBalanceAt, "params": []interface{}{common.HexToAddress("0xcDd37Ada79F589c15bD4f8fD2083dc88E34A2af2")}},
			"",
		},
		{
			"SOLANA", "http://127.0.0.1:8545/solana",
			map[string]interface{}{"chain_env": consts.ChainEnvMainnet, "method": consts.SolanaMethodGetBalance, "params": []interface{}{"5EhGYUyQNrxgUbuYF4vbL2SZDT6RMfhq3yjeyevvULeC"}},
			"",
		},
	}
	t.Run("Unite Case", func(t *testing.T) {
		for i, c := range cases {
			t.Run(c.name+" "+strconv.Itoa(i+1), func(t *testing.T) {
				_, err := req.C().DevMode().R().SetBody(c.params).Post(c.url)
				assert.Nil(t, err)
			})
		}
	})
}
