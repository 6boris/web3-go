package eth

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"web3.go/entity/resp"
)

func (eth *Eth) GetBalance(address string, defaultBlockParameter string) (*big.Float, error) {
	if defaultBlockParameter == "" {
		defaultBlockParameter = "latest"
	}
	bodyByte, err := eth.provider.SendRequest("eth_getBalance", []string{address, defaultBlockParameter})

	if err != nil {
		return nil, err
	}
	respBalance := new(resp.EthGetBalanceResp)
	err = json.Unmarshal(bodyByte, respBalance)
	if err != nil {
		return nil, err
	}
	if respBalance.Error.Code != 0 {
		return nil, errors.New(respBalance.Error.Message)
	}

	balance, isSuccess := big.NewInt(0).SetString(respBalance.Result[2:], 16)
	if !isSuccess {
		return nil, errors.New(fmt.Sprintf("format balance err: %s", respBalance.Result))
	}
	return big.NewFloat(0).SetInt(balance), err
}
