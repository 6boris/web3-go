package eth

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"web3.go/consts"
	"web3.go/providers"
	"web3.go/utils"
)

func TestEth_Accounts(t *testing.T) {
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("eth_accounts", func(t *testing.T) {
		accounts, err := web3.Accounts()
		assert.Nil(t, err)
		fmt.Println(accounts)
	})
}

func TestListBalances(t *testing.T) {
	eth := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	accounts, err := eth.Accounts()
	assert.Nil(t, err)
	web3 := NewEth(providers.NewHTTPProvider(consts.HOST_HTTP_PROVIDER_LOCAL, 10))
	t.Run("list account ", func(t *testing.T) {
		ans := make(map[string]*big.Float)
		for _, v := range accounts {
			balance, err := web3.GetBalance(v, "")
			assert.Nil(t, err)
			ans[v] = balance
		}
		for i, v := range ans {
			fmt.Println(fmt.Sprintf("account: %s  balance: %s ETH", i, utils.FromWei(v).String()))
		}
	})
}
