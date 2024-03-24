package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/6boris/web3-go/pkg/wjson"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewEvmClient(t *testing.T) {
	testAccount := common.HexToAddress("0xf15689636571dba322b48E9EC9bA6cFB3DF818e1")
	testBlockHash := common.HexToHash("0x21e0118bd8618a14632f82e1445c1f178cfab5bbd2debf1eb3338886ced75e15")
	testBlockNumber := big.NewInt(19454574)
	testTxHash := common.HexToHash("0xb382be540032c6751b97cb623ff2bbcd0f1da2a386c4a75f7e561ce4ebefb787")
	testEvmPolygonUSDTAddress := common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	t.Run("BlockByHash", func(t *testing.T) {
		blockInfo, err := testEvmEthClient.BlockByHash(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(blockInfo)
	})
	t.Run("BlockByNumber", func(t *testing.T) {
		blockInfo, err := testEvmEthClient.BlockByNumber(testCtx, testBlockNumber)
		assert.Nil(t, err)
		spew.Dump(blockInfo)
	})
	t.Run("HeaderByHash", func(t *testing.T) {
		headerInfo, err := testEvmEthClient.HeaderByHash(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(headerInfo)
	})
	t.Run("HeaderByNumber", func(t *testing.T) {
		headerInfo, err := testEvmEthClient.HeaderByNumber(testCtx, testBlockNumber)
		assert.Nil(t, err)
		spew.Dump(headerInfo)
	})
	t.Run("TransactionCount", func(t *testing.T) {
		txCount, err := testEvmEthClient.TransactionCount(testCtx, testBlockHash)
		assert.Nil(t, err)
		spew.Dump(txCount)
	})
	t.Run("TransactionInBlock", func(t *testing.T) {
		txCount, err := testEvmEthClient.TransactionInBlock(testCtx, testBlockHash, 1)
		assert.Nil(t, err)
		spew.Dump(txCount)
	})
	t.Run("TransactionByHash", func(t *testing.T) {
		txInfo, isPending, err := testEvmEthClient.TransactionByHash(testCtx, testTxHash)
		assert.Nil(t, err)
		spew.Dump(isPending)
		spew.Dump(txInfo)
	})
	t.Run("TransactionReceipt", func(t *testing.T) {
		txInfo, err := testEvmEthClient.TransactionReceipt(testCtx, testTxHash)
		assert.Nil(t, err)
		spew.Dump(txInfo)
	})
	t.Run("SendTransaction_Native_Token", func(t *testing.T) {
		privateKey, err := crypto.HexToECDSA(os.Getenv("WEB3_GO_DEV_KEY_2"))
		assert.Nil(t, err)
		publicKey := privateKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		assert.Nil(t, err)
		sendAccount := crypto.PubkeyToAddress(*publicKeyECDSA)

		chainID, err := testEvmPolygonClient.ChainID(context.Background())
		assert.Nil(t, err)

		nonce, err := testEvmPolygonClient.PendingNonceAt(context.Background(), sendAccount)
		assert.Nil(t, err)

		gasLimit := uint64(21000)
		value := decimal.New(1, 0).BigInt()
		gasPrice, err := testEvmPolygonClient.SuggestGasPrice(testCtx)
		gasPrice = decimal.NewFromBigInt(gasPrice, 0).
			Mul(decimal.NewFromFloat(2)).BigInt()
		assert.Nil(t, err)
		signedTx, err := types.SignNewTx(privateKey, types.NewEIP155Signer(chainID), &types.LegacyTx{
			To:       &sendAccount,
			Nonce:    nonce,
			Value:    value,
			Gas:      gasLimit,
			GasPrice: gasPrice,
		})
		assert.Nil(t, err)

		err = testEvmPolygonClient.SendTransaction(context.Background(), signedTx)
		assert.Nil(t, err)
		fmt.Println(wjson.StructToJsonStringWithIndent(map[string]interface{}{
			"chain_id":    chainID,
			"sendAccount": sendAccount,
			"toAccount":   sendAccount,
			"nonce":       nonce,
			"gas_limit":   gasLimit,
			"value":       value,
			"gas_price":   decimal.NewFromBigInt(gasPrice, -9),
			"tx_hash":     signedTx.Hash().String(),
		}, "", "  "))
	})
	t.Run("SendTransactionSimple", func(t *testing.T) {
		tx, err := testEvmPolygonClient.SendTransactionSimple(
			testCtx,
			testEvmPolygonClient.GetAllSinners()[0],
			testEvmPolygonClient.GetAllSinners()[1],
			decimal.NewFromFloat(0.00001).Mul(decimal.New(1, 18)).BigInt(),
		)
		assert.Nil(t, err)
		spew.Dump(tx.Hash())
	})
	t.Run("BalanceAt", func(t *testing.T) {
		accountBalance, err := testEvmEthClient.BalanceAt(testCtx, testAccount, nil)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(accountBalance, -18))
	})
	t.Run("StorageAt", func(t *testing.T) {
		storageBytes, err := testEvmEthClient.StorageAt(
			testCtx, testAccount,
			common.HexToHash(""), big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(storageBytes)
	})
	t.Run("CodeAt", func(t *testing.T) {
		codeBytes, err := testEvmEthClient.CodeAt(
			testCtx, common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
			big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(codeBytes)
	})
	t.Run("NonceAt", func(t *testing.T) {
		nonceAt, err := testEvmEthClient.NonceAt(
			testCtx, testAccount,
			big.NewInt(19454700),
		)
		assert.Nil(t, err)
		spew.Dump(nonceAt)
	})
	t.Run("SuggestGasPrice", func(t *testing.T) {
		gasPrice, err := testEvmEthClient.SuggestGasPrice(testCtx)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(gasPrice, -18))
	})
	t.Run("SuggestGasTipCap", func(t *testing.T) {
		gasPrice, err := testEvmEthClient.SuggestGasTipCap(testCtx)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(gasPrice, -18))
	})
	t.Run("FeeHistory", func(t *testing.T) {
		feeHistory, err := testEvmEthClient.FeeHistory(testCtx,
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
		gasInfo, err := testEvmEthClient.EstimateGas(testCtx, msg)
		assert.Nil(t, err)
		spew.Dump(gasInfo)
	})
	t.Run("PendingBalanceAt", func(t *testing.T) {
		accountInfo, err := testEvmEthClient.PendingBalanceAt(testCtx, testAccount)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(accountInfo, -18))
	})
	t.Run("PendingStorageAt", func(t *testing.T) {
		storageBytes, err := testEvmEthClient.PendingStorageAt(
			testCtx, testAccount,
			common.HexToHash(""),
		)
		assert.Nil(t, err)
		spew.Dump(storageBytes)
	})
	t.Run("PendingCodeAt", func(t *testing.T) {
		codeBytes, err := testEvmEthClient.PendingCodeAt(
			testCtx, common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		)
		assert.Nil(t, err)
		spew.Dump(codeBytes)
	})
	t.Run("PendingNonceAt", func(t *testing.T) {
		nonceAt, err := testEvmEthClient.PendingNonceAt(
			testCtx, testAccount,
		)
		assert.Nil(t, err)
		spew.Dump(nonceAt)
	})
	t.Run("PendingTransactionCount", func(t *testing.T) {
		pendingTxCount, err := testEvmEthClient.PendingTransactionCount(testCtx)
		assert.Nil(t, err)
		spew.Dump(pendingTxCount)
	})
	t.Run("BlockNumber", func(t *testing.T) {
		blockNumber, err := testEvmEthClient.BlockNumber(testCtx)
		assert.Nil(t, err)
		spew.Dump(blockNumber)
	})
	t.Run("ChainID", func(t *testing.T) {
		chanID, err := testEvmEthClient.ChainID(testCtx)
		assert.Nil(t, err)
		spew.Dump(chanID)
	})
	t.Run("ERC20Name", func(t *testing.T) {
		tokenName, err := testEvmPolygonClient.ERC20Name(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		spew.Dump(tokenName)
	})
	t.Run("ERC20Symbol", func(t *testing.T) {
		tokenSymbol, err := testEvmPolygonClient.ERC20Symbol(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		spew.Dump(tokenSymbol)
	})
	t.Run("ERC20Decimals", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		spew.Dump(tokenDecimals)
	})
	t.Run("ERC20BalanceOf", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		tokenBalance, err := testEvmPolygonClient.ERC20BalanceOf(testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(tokenBalance, -int32(tokenDecimals)))
	})
	t.Run("ERC20TotalSupply", func(t *testing.T) {
		tokenSymbol, err := testEvmPolygonClient.ERC20Symbol(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		tokenBalance, err := testEvmPolygonClient.ERC20TotalSupply(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		fmt.Printf("TotalSupply:%s %s\n",
			decimal.NewFromBigInt(tokenBalance, -int32(tokenDecimals)).String(), tokenSymbol)
	})
	t.Run("ERC20Transfer", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)
		txInfo, err := testEvmPolygonClient.ERC20Transfer(
			testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonClient.GetAllSinners()[0],
			testEvmPolygonClient.GetAllSinners()[1],
			decimal.NewFromFloat(0.00001).Mul(decimal.New(1, int32(tokenDecimals))).BigInt(),
		)
		assert.Nil(t, err)
		fmt.Println(txInfo.Hash().String())
	})
	t.Run("ERC20Approve", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		txInfo, err := testEvmPolygonClient.ERC20Approve(
			testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonClient.GetAllSinners()[0], testEvmPolygonClient.GetAllSinners()[1],
			decimal.NewFromFloat(0.00001).Mul(decimal.New(1, int32(tokenDecimals))).BigInt(),
		)
		assert.Nil(t, err)
		fmt.Println("ERC20Approve", txInfo.Hash().String())
	})
	t.Run("ERC20IncreaseAllowance", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		txInfo, err := testEvmPolygonClient.ERC20IncreaseAllowance(
			testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonClient.GetAllSinners()[0], testEvmPolygonClient.GetAllSinners()[1],
			decimal.NewFromFloat(0.00001).Mul(decimal.New(1, int32(tokenDecimals))).BigInt(),
		)
		assert.Nil(t, err)
		fmt.Println("ERC20Approve", txInfo.Hash().String())
	})
	t.Run("ERC20DecreaseAllowance", func(t *testing.T) {
		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		txInfo, err := testEvmPolygonClient.ERC20DecreaseAllowance(
			testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonClient.GetAllSinners()[0], testEvmPolygonClient.GetAllSinners()[1],
			decimal.NewFromFloat(0.00001).Mul(decimal.New(1, int32(tokenDecimals))).BigInt(),
		)
		assert.Nil(t, err)
		fmt.Println("ERC20Approve", txInfo.Hash().String())
	})
	t.Run("ERC20Allowance", func(t *testing.T) {
		tokenSymbol, err := testEvmPolygonClient.ERC20Symbol(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		tokenDecimals, err := testEvmPolygonClient.ERC20Decimals(testCtx, testEvmPolygonUSDTAddress)
		assert.Nil(t, err)

		tokenAllowance, err := testEvmPolygonClient.ERC20Allowance(
			testCtx,
			testEvmPolygonUSDTAddress,
			testEvmPolygonClient.GetAllSinners()[0], testEvmPolygonClient.GetAllSinners()[1],
		)
		assert.Nil(t, err)
		fmt.Printf("ERC20Allowance:%s %s\n",
			decimal.NewFromBigInt(tokenAllowance, -int32(tokenDecimals)).String(), tokenSymbol)
	})
}
