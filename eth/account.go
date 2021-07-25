package eth

import (
	"encoding/json"
	"errors"

	"web3.go/entity/resp"
)

func (eth *Eth) Accounts() ([]string, error) {

	bodyByte, err := eth.provider.SendRequest("eth_accounts", []string{})

	if err != nil {
		return nil, err
	}
	respAccounts := new(resp.JsonRpcStringArrayCommonResp)
	err = json.Unmarshal(bodyByte, respAccounts)
	if err != nil {
		return nil, err
	}
	if respAccounts.Error.Code != 0 {
		return nil, errors.New(respAccounts.Error.Message)
	}

	return respAccounts.Result, err
}
