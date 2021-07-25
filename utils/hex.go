package utils

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

const ETHEREUM_HEX_PREFIX = "0x"

// HexToBigInt 16 进制转
func HexToBigInt(s string) (*big.Int, error) {
	if s[:2] != ETHEREUM_HEX_PREFIX {
		return nil, errors.New("hex must start with 0x")
	}
	n, isSuccess := big.NewInt(0).SetString(s[2:], 16)
	if !isSuccess {
		return nil, errors.New("hex format error")
	}
	return n, nil
}

func BigIntToHex(n big.Int) string {
	return fmt.Sprintf("%s%s", ETHEREUM_HEX_PREFIX, strconv.FormatInt(n.Int64(), 16))
}

// HexToBigInt 16 进制转
func HexToBigFloat(s string) (*big.Int, error) {
	if s[:2] != ETHEREUM_HEX_PREFIX {
		return nil, errors.New("hex must start with 0x")
	}
	n, isSuccess := big.NewInt(0).SetString(s[2:], 16)
	if !isSuccess {
		return nil, errors.New("hex format error")
	}
	return n, nil
}

func IntToHex(n int) string {
	return fmt.Sprintf("%s%s", ETHEREUM_HEX_PREFIX, strconv.FormatInt(int64(n), 16))
}

var ETHER_EXPONENTIAL = big.NewInt(1000000000000000000)

func ToWei(z *big.Float) (*big.Float, error) {
	return big.NewFloat(0).Set(z.Mul(z, big.NewFloat(0).SetInt(ETHER_EXPONENTIAL))), nil
	//return big.NewFloat(0).SetInt64(n.Mul(n, ETHER_EXPONENTIAL).Int64()), nil
}

func FromWei(z *big.Float) *big.Float {
	return big.NewFloat(0).Set(z.Quo(z, big.NewFloat(0).SetInt(ETHER_EXPONENTIAL)))
}
