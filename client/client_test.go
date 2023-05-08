package client

import (
	"context"
	"github.com/6boris/web3-go/pkg/otel"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
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
	otel.InitProvider()
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
	t.Run("BlockByHash", func(t *testing.T) {
		resp, err := testClient.BlockByHash(testCtx, common.HexToHash("0xb07928fe01fe07c2fc5b743aea2b9b7262bf79854214ff2e828760b55d191c1e"))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("BlockByNumber", func(t *testing.T) {
		resp, err := testClient.BlockByNumber(testCtx, big.NewInt(17180303))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("TransactionCount", func(t *testing.T) {
		resp, err := testClient.TransactionCount(testCtx, common.HexToHash("0xb07928fe01fe07c2fc5b743aea2b9b7262bf79854214ff2e828760b55d191c1e"))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("PendingTransactionCount", func(t *testing.T) {
		resp, err := testClient.PendingTransactionCount(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("GetUncleCountByBlockHash", func(t *testing.T) {
		resp, err := testClient.GetUncleCountByBlockHash(testCtx, common.HexToHash("0xb07928fe01fe07c2fc5b743aea2b9b7262bf79854214ff2e828760b55d191c1e"))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("GetUncleCountByBlockNumber", func(t *testing.T) {
		resp, err := testClient.GetUncleCountByBlockNumber(testCtx, big.NewInt(17180303))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("SyncProgress", func(t *testing.T) {
		resp, err := testClient.SyncProgress(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("FeeHistory", func(t *testing.T) {
		resp, err := testClient.FeeHistory(testCtx, 1, big.NewInt(17180303), []float64{})
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("FilterLogs", func(t *testing.T) {
		resp, err := testClient.FilterLogs(testCtx, ethereum.FilterQuery{
			FromBlock: big.NewInt(17180303),
			ToBlock:   big.NewInt(17180303),
		})
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("StorageAt", func(t *testing.T) {
		resp, err := testClient.StorageAt(testCtx, common.HexToAddress("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5"), common.HexToHash("0x7c6747490ff8726ef0a5348dfd0323339a7de181afdd6a4b6a8053789ebffed2"), big.NewInt(17180303))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("CodeAt", func(t *testing.T) {
		resp, err := testClient.CodeAt(testCtx, common.HexToAddress("0x08d740B96Cd673cD00626a2286D1f488597C15fd"), big.NewInt(17180536))
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
