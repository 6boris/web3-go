package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
)

func TestToBigInt(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		decN := 111111
		hexN := IntToHex(decN)
		actualN, err := HexToBigInt(hexN)
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("%d", decN), actualN.String())
	})
	t.Run("big", func(t *testing.T) {
		decN := 111111000000000000
		hexN := IntToHex(decN)
		actualN, err := HexToBigInt(hexN)
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("%d", decN), actualN.String())
	})

}

func TestBigIntToHex(t *testing.T) {
	t.Run("hex", func(t *testing.T) {
		bInt, succes := big.NewInt(0).SetString("99999999999999999999", 10)
		assert.Equal(t, succes, true)
		fmt.Println(BigIntToHex(*bInt))
	})
}

func TestWei(t *testing.T) {
	t.Run("to wei", func(t *testing.T) {
		eBalance := big.NewFloat(float64(233))
		wBalance, err := ToWei(eBalance)
		assert.Nil(t, err)
		fmt.Println(wBalance)
	})
	t.Run("from wei", func(t *testing.T) {
		wBalance, err := FromWei(big.NewFloat(float64(12345678912345678912344567)))
		assert.Nil(t, err)
		fmt.Println(wBalance)
	})
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}
