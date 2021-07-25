package client

import (
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
}
