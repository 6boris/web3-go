package eth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"web3.go/consts"
	"web3.go/entity/req"
	"web3.go/providers"
)

func TestEth_GetTransactionCount(t *testing.T) {
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("eth_getTransactionCount", func(t *testing.T) {
		count, err := web3.GetTransactionCount("", "")
		assert.Nil(t, err)
		assert.Equal(t, "0", count.String())
		fmt.Println(count)
	})
}

func TestEth_GetBlockTransactionCountByHash(t *testing.T) {
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("eth_getTransactionCount", func(t *testing.T) {
		count, err := web3.GetBlockTransactionCountByHash("")
		assert.Nil(t, err)
		assert.Equal(t, "0", count.String())
		fmt.Println(count)
	})
}

func TestEth_SendTransaction(t *testing.T) {
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("eth_sendTransaction", func(t *testing.T) {
		count, err := web3.SendTransaction([]req.SendTransactionReq{
			{
				From:     "",
				To:       "",
				Gas:      "0x76c0",
				GasPrice: "0x9184e72a000",
				Value:    "0x9184e72a",
				Data:     "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
			},
		})
		assert.Nil(t, err)
		fmt.Println(count)
	})
}
