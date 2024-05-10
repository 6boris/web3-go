package client

import (
	"testing"

	"github.com/6boris/web3-go/model/solana"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestSolana_Unite_GetAccountInfo(t *testing.T) {
	t.Run("Contract_USDT", func(t *testing.T) {
		reply, err := testSolanaClient.GetAccountInfo(testCtx,
			&solana.GetAccountInfoRequest{Account: "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB"},
		)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
	t.Run("User_Common", func(t *testing.T) {
		reply, err := testSolanaClient.GetAccountInfo(testCtx,
			&solana.GetAccountInfoRequest{Account: "Hw161dCAE9VBtbTo3EgbzwLdvVxNnEpiqqAJw7q38BYE"},
		)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetBalance(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		reply, err := testSolanaClient.GetBalance(testCtx,
			&solana.GetBalanceRequest{Account: "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB"},
		)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetVersion(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		reply, err := testSolanaClient.GetVersion(testCtx)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetTokenAccountBalance(t *testing.T) {
	t.Run("SLND", func(t *testing.T) {
		reply, err := testSolanaClient.GetTokenAccountBalance(testCtx,
			&solana.GetTokenAccountBalanceRequest{Account: "3Lz6rCrXdLybFiuJGJnEjv6Z2XtCh5n4proPGP2aBkA1"})
		assert.Nil(t, err)
		spew.Dump(reply)
	})
	t.Run("TNSR", func(t *testing.T) {
		reply, err := testSolanaClient.GetTokenAccountBalance(testCtx,
			&solana.GetTokenAccountBalanceRequest{Account: "5FhnYa75QKfMkPjBCrM7iucf2wMBNzHE2chyyTUfJEqj"})
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetBlockHeight(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		reply, err := testSolanaClient.GetBlockHeight(testCtx)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetBlockTime(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		blockHeight, err := testSolanaClient.GetBlockHeight(testCtx)
		assert.Nil(t, err)
		reply, err := testSolanaClient.GetBlockTime(testCtx, blockHeight)
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetBlock(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		reply, err := testSolanaClient.GetBlock(testCtx, &solana.GetBlockRequest{
			Slot:               430,
			Encoding:           "base58",
			TransactionDetails: "full",
			Rewards:            false,
		})
		assert.Nil(t, err)
		spew.Dump(reply)
	})
}

func TestSolana_Unite_GetClusterNodes(t *testing.T) {
	t.Run("Demo", func(t *testing.T) {
		reply, err := testSolanaClient.GetClusterNodes(testCtx)
		assert.Nil(t, err)
		spew.Dump(reply)
		spew.Dump(len(reply))
	})
}
