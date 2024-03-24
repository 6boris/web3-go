package client

import (
	"context"
	"os"
	"testing"

	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/6boris/web3-go/pkg/pk"
	"github.com/shopspring/decimal"
)

var (
	testCtx              context.Context
	testPool             *Pool
	testEvmEthClient     *EvmClient
	testEvmPolygonClient *EvmClient
	testSolanaClient     *SolanaClient
)

func TestMain(m *testing.M) {
	var err error
	evmSigners := make([]*clientModel.ConfEvmChainSigner, 0)
	for _, v := range []string{"WEB3_GO_DEV_KEY_1", "WEB3_GO_DEV_KEY_1"} {
		signer, loopErr := pk.TransformPkToEvmSigner(os.Getenv(v))
		if loopErr != nil {
			continue
		}
		evmSigners = append(evmSigners, signer)
	}
	testEvmEthClient, err = NewEvmClient(&clientModel.ConfEvmChainClient{TransportURL: "https://1rpc.io/eth"})
	if err != nil {
		panic(err)
	}
	testEvmPolygonClient, err = NewEvmClient(&clientModel.ConfEvmChainClient{
		TransportURL: "https://1rpc.io/matic",
		GasFeeRate:   decimal.NewFromFloat(2),
		GasLimitMax:  decimal.NewFromInt(30000000),
		GasLimitRate: decimal.NewFromFloat(1.5),
		Signers:      evmSigners,
	})
	if err != nil {
		panic(err)
	}
	testSolanaClient, err = NewSolanaClient(&clientModel.ConfSolanaClient{TransportURL: "https://api.mainnet-beta.solana.com"})
	if err != nil {
		panic(err)
	}
	testPool = NewPool(clientModel.GetDefaultConfPool())
	testCtx = context.TODO()
	// Before Test
	code := m.Run()
	// After Test
	os.Exit(code)
}
