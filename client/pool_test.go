package client

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Unite_Pool(t *testing.T) {
	chainID := int64(1)
	t.Run("Web3ClientVersion", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			resp, err := testPool.GetClient(chainID).ClientVersion(testCtx)
			assert.Nil(t, err)
			spew.Dump(resp)
		}
	})
	t.Run("ChainID", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			resp, err := testPool.GetClient(chainID).ChainID(testCtx)
			assert.Nil(t, err)
			spew.Dump(resp)
		}
	})
	t.Run("NetworkID", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			resp, err := testPool.GetClient(chainID).NetworkID(testCtx)
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
			resp, err := testPool.GetClient(c).BlockNumber(testCtx)
			assert.Nil(t, err)
			fmt.Println(fmt.Sprintf("ChainID:%d BlockNumber:%d", c, resp))
		}
	})
}
