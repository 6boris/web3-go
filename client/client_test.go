package client

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"testing"
)

var (
	testClient *Client
	testPool   *Pool
	testCtx    context.Context
)

func TestMain(m *testing.M) {
	var err error
	testClient, err = NewClient(&ConfClient{TransportURL: "https://eth.llamarpc.com"})
	if err != nil {
		panic(err)
	}
	testPool = NewPool(GetDefaultConfPool())
	testCtx = context.TODO()
	//beforeTest(config.Conf)
	code := m.Run()
	//afterTest(config.Conf)
	os.Exit(code)
}

func Test_Unite_Client(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		testClient.Close()
	})
	t.Run("Web3ClientVersion", func(t *testing.T) {
		resp, err := testClient.ClientVersion(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("ChainID", func(t *testing.T) {
		resp, err := testClient.ChainID(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("NetworkID", func(t *testing.T) {
		resp, err := testClient.NetworkID(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("SuggestGasPrice", func(t *testing.T) {
		resp, err := testClient.SuggestGasPrice(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("BlockNumber", func(t *testing.T) {
		resp, err := testClient.BlockNumber(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("HexToAddress", func(t *testing.T) {
		resp, err := testClient.BalanceAt(testCtx, common.HexToAddress("0xB671e841a8e6DB528358Ed385983892552EF422f"), big.NewInt(17177806))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
