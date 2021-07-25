package eth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3.go/consts"
	"web3.go/providers"
)

func TestGetBalance(t *testing.T) {
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("address err", func(t *testing.T) {
		_, err := web3.GetBalance("", "")
		assert.NotNil(t, err)
	})
	t.Run("eth_getBalance", func(t *testing.T) {
		balance, err := web3.GetBalance("", "")
		assert.Nil(t, err)
		assert.Equal(t, "0", balance.String())
		fmt.Println(balance)
	})
	t.Run("eth_gasPrice", func(t *testing.T) {
		v, err := web3.GasPrice()
		assert.Nil(t, err)
		assert.NotEqual(t, "0", v.String())
	})
}
