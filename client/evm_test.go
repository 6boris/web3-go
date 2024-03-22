package client

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewEvmClient(t *testing.T) {
	testAccount := common.HexToAddress("0xf15689636571dba322b48E9EC9bA6cFB3DF818e1")
	testBlockHash := common.HexToHash("0x21e0118bd8618a14632f82e1445c1f178cfab5bbd2debf1eb3338886ced75e15")
	testBlockNumber := big.NewInt(19454574)
	testTxHash := common.HexToHash("0xb382be540032c6751b97cb623ff2bbcd0f1da2a386c4a75f7e561ce4ebefb787")

	t.Run("BlockByHash", func(t *testing.T) {
		blockInfo, err := testEvmClient.BlockByHash(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(blockInfo)
	})
	t.Run("BlockByNumber", func(t *testing.T) {
		blockInfo, err := testEvmClient.BlockByNumber(testCtx, testBlockNumber)
		assert.Nil(t, err)
		spew.Dump(blockInfo)
	})
	t.Run("HeaderByHash", func(t *testing.T) {
		headerInfo, err := testEvmClient.HeaderByHash(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(headerInfo)
	})
	t.Run("HeaderByNumber", func(t *testing.T) {
		headerInfo, err := testEvmClient.HeaderByNumber(testCtx, testBlockNumber)
		assert.Nil(t, err)
		spew.Dump(headerInfo)
	})
	t.Run("TransactionCount", func(t *testing.T) {
		txCount, err := testEvmClient.TransactionCount(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(txCount)
	})
	t.Run("TransactionInBlock", func(t *testing.T) {
		txCount, err := testEvmClient.TransactionInBlock(testCtx, testBlockHash, 1)
		assert.Nil(t, err)
		spew.Dump(txCount)
	})
	t.Run("TransactionByHash", func(t *testing.T) {
		txInfo, isPending, err := testEvmClient.TransactionByHash(testCtx, testTxHash)
		assert.Nil(t, err)
		spew.Dump(isPending)
		spew.Dump(txInfo)
	})
	t.Run("TransactionReceipt", func(t *testing.T) {
		txInfo, err := testEvmClient.TransactionReceipt(testCtx, testTxHash)
		assert.Nil(t, err)
		spew.Dump(txInfo)
	})
	t.Run("SendTransaction", func(t *testing.T) {
		tx := types.NewContractCreation(0, big.NewInt(0), 2100, big.NewInt(12000), common.FromHex("0x"))
		err := testEvmClient.SendTransaction(testCtx, tx)
		assert.NotNil(t, err)
	})
	t.Run("BalanceAt", func(t *testing.T) {
		accountBalance, err := testEvmClient.BalanceAt(testCtx, testAccount, nil)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(accountBalance, -18))
	})
	t.Run("StorageAt", func(t *testing.T) {
		storageBytes, err := testEvmClient.StorageAt(
			testCtx, testAccount,
			common.HexToHash(""), big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(storageBytes)
	})
	t.Run("CodeAt", func(t *testing.T) {
		codeBytes, err := testEvmClient.CodeAt(
			testCtx, common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
			big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(codeBytes)
	})
	t.Run("NonceAt", func(t *testing.T) {
		nonceAt, err := testEvmClient.NonceAt(
			testCtx, testAccount,
			big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(nonceAt)
	})
	t.Run("SuggestGasPrice", func(t *testing.T) {
		gasPrice, err := testEvmClient.SuggestGasPrice(testCtx)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(gasPrice, -18))
	})
	t.Run("SuggestGasTipCap", func(t *testing.T) {
		gasPrice, err := testEvmClient.SuggestGasTipCap(testCtx)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(gasPrice, -18))
	})
	t.Run("FeeHistory", func(t *testing.T) {
		feeHistory, err := testEvmClient.FeeHistory(testCtx,
			2,
			testBlockNumber,
			[]float64{0.008912678667376286},
		)
		assert.Nil(t, err)
		spew.Dump(feeHistory)
	})
	t.Run("EstimateGas", func(t *testing.T) {
		msg := ethereum.CallMsg{
			From:  testAccount,
			To:    &common.Address{},
			Gas:   21000,
			Value: big.NewInt(1),
		}
		gasInfo, err := testEvmClient.EstimateGas(testCtx, msg)
		assert.Nil(t, err)
		spew.Dump(gasInfo)
	})
	t.Run("PendingBalanceAt", func(t *testing.T) {
		accountInfo, err := testEvmClient.PendingBalanceAt(testCtx, testAccount)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(accountInfo, -18))
	})
	t.Run("PendingStorageAt", func(t *testing.T) {
		storageBytes, err := testEvmClient.PendingStorageAt(
			testCtx, testAccount,
			common.HexToHash(""),
		)
		assert.Nil(t, err)
		spew.Dump(storageBytes)
	})
	t.Run("PendingCodeAt", func(t *testing.T) {
		codeBytes, err := testEvmClient.PendingCodeAt(
			testCtx, common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		)
		assert.Nil(t, err)
		spew.Dump(codeBytes)
	})
	t.Run("PendingNonceAt", func(t *testing.T) {
		nonceAt, err := testEvmClient.PendingNonceAt(
			testCtx, testAccount,
		)
		assert.Nil(t, err)
		spew.Dump(nonceAt)
	})
	t.Run("PendingTransactionCount", func(t *testing.T) {
		pendingTxCount, err := testEvmClient.PendingTransactionCount(testCtx)
		assert.Nil(t, err)
		spew.Dump(pendingTxCount)
	})
	t.Run("BlockNumber", func(t *testing.T) {
		blockNumber, err := testEvmClient.BlockNumber(testCtx)
		assert.Nil(t, err)
		spew.Dump(blockNumber)
	})
	t.Run("ChainID", func(t *testing.T) {
		chanID, err := testEvmClient.ChainID(testCtx)
		assert.Nil(t, err)
		spew.Dump(chanID)
	})
}
