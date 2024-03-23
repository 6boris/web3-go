package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/shopspring/decimal"

	"github.com/6boris/web3-go/consts"

	"github.com/6boris/web3-go/model/solana"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_Unite_Pool(t *testing.T) {
	chainID := int64(1)
	t.Run("ChainID", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			resp, err := testPool.GetEvmClient(chainID).ChainID(testCtx)
			assert.Nil(t, err)
			spew.Dump(resp)
		}
	})
	t.Run("BlockNumber", func(t *testing.T) {
		chainCase := []int64{
			1, 5, 11155111, // Ethereum
			56, 97, // Bsc
			137, 80001, // Polygon
			250, 4002, // Fantom
			10, 420, // Optimistic
		}
		for _, c := range chainCase {
			resp, err := testPool.GetEvmClient(c).BlockNumber(testCtx)
			assert.Nil(t, err)
			fmt.Println(fmt.Printf("ChainID:%d BlockNumber:%d", c, resp))
		}
	})
	t.Run("Evm_SendTransaction", func(t *testing.T) {

		privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}
		testAccount := crypto.PubkeyToAddress(*publicKeyECDSA)
		fmt.Println("testAccount", testAccount)
		testChainID := int64(137)
		nonce, err := testPool.GetEvmClient(testChainID).PendingNonceAt(testCtx, testAccount)
		assert.Nil(t, err)
		fmt.Println("Nonce:", nonce)
		//value := big.NewInt(1000000000000000000)
		//gasLimit := uint64(21000)
		gasPrice, err := testPool.GetEvmClient(testChainID).SuggestGasPrice(context.Background())
		assert.Nil(t, err)
		fmt.Println("gasPrice:", gasPrice)

		//err := testPool.GetEvmClient(137).SendTransaction(testCtx, nil)
		//assert.Nil(t, err)
	})
	t.Run("Evm_ChainID", func(t *testing.T) {
		resp, err := testPool.GetEvmClient(1).ChainID(testCtx)
		assert.Nil(t, err)
		spew.Dump(resp)
	})
	t.Run("Evm_SuggestGasTipCap", func(t *testing.T) {
		gasPrice, err := testPool.GetEvmClient(1).SuggestGasTipCap(testCtx)
		assert.Nil(t, err)
		spew.Dump(decimal.NewFromBigInt(gasPrice, -18))
	})

	t.Run("Solana_GetBalance", func(t *testing.T) {
		resp, err := testPool.GetSolanaClient(consts.ChainEnvMainnet).GetBalance(testCtx, &solana.GetBalanceRequest{Account: "5EhGYUyQNrxgUbuYF4vbL2SZDT6RMfhq3yjeyevvULeC"})
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
