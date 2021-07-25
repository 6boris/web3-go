package web3

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web3.go/consts"
	"web3.go/providers"
)

func Test_Web3_GetBalance(t *testing.T) {
	web3 := NewWeb3(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))

	t.Run("web3_clientVersion", func(t *testing.T) {
		v, err := web3.ClientVersion()
		assert.Nil(t, err)
		assert.NotNil(t, v)
	})
}
