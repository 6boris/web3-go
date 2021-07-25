package eth

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"web3.go/consts"
	"web3.go/entity/req"
	"web3.go/entity/resp"
	"web3.go/utils"
)

// GetTransactionCount 查询交易数量
func (eth *Eth) GetTransactionCount(address string, defaultBlockParameter string) (*big.Int, error) {
	if defaultBlockParameter == "" {
		defaultBlockParameter = "latest"
	}
	bodyByte, err := eth.provider.SendRequest(consts.METHOD_ETH_GET_TRANSTING_COUNT, []string{address, defaultBlockParameter})

	if err != nil {
		return nil, err
	}
	respBalance := new(resp.JsonRpcHexCommonResp)
	err = json.Unmarshal(bodyByte, respBalance)
	if err != nil {
		return nil, err
	}
	if respBalance.Error.Code != 0 {
		return nil, errors.New(respBalance.Error.Message)
	}
	balance, err := utils.HexToBigInt(respBalance.Result)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetBlockTransactionCountByHash 查询交易数量
func (eth *Eth) GetBlockTransactionCountByHash(hash string) (*big.Int, error) {
	bodyByte, err := eth.provider.SendRequest(consts.METHOD_ETH_GET_TRANSTING_COUNT_BY_HASH, []string{hash})

	if err != nil {
		return nil, err
	}
	respBalance := new(resp.JsonRpcHexCommonResp)
	err = json.Unmarshal(bodyByte, respBalance)
	if err != nil {
		return nil, err
	}
	if respBalance.Error.Code != 0 {
		return nil, errors.New(respBalance.Error.Message)
	}

	balance, isSuccess := big.NewInt(0).SetString(respBalance.Result[2:], 16)
	if !isSuccess {
		return nil, errors.New(fmt.Sprintf("Hexadecimal conversion failed: %s", respBalance.Result))
	}

	return balance, err
}

// SendTransaction 发送交易
func (eth *Eth) SendTransaction(params []req.SendTransactionReq) (string, error) {
	bodyByte, err := eth.provider.SendRequest(
		consts.METHOD_ETH_SEND_TRANSACTION,
		params,
	)
	if err != nil {
		return "", err
	}
	respBalance := new(resp.JsonRpcHashCommonResp)
	err = json.Unmarshal(bodyByte, respBalance)
	if err != nil {
		return "", err
	}

	return respBalance.Result, err
}
