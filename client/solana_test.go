package client

import (
	"testing"

	"github.com/6boris/web3-go/model/solana"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestSolanaClient(t *testing.T) {
	t.Run("GetBalance", func(t *testing.T) {
		reply, err := testSolanaClient.GetBalance(testCtx, &solana.GetBalanceRequest{Account: "5EhGYUyQNrxgUbuYF4vbL2SZDT6RMfhq3yjeyevvULeC"})
		assert.Nil(t, err)
		spew.Dump(reply)
	})

	t.Run("GetTokenAccountBalance", func(t *testing.T) {
		reply, err := testSolanaClient.GetTokenAccountBalance(testCtx,
			&solana.GetTokenAccountBalanceRequest{Account: "J6feTwcYDydW71Dp9qgfW7Mu9qk3qDRrDZAWV8NMVh9x"})
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}
