package eth

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"web3.go/consts"
	"web3.go/entity/resp"
	"web3.go/providers"
)

type Eth struct {
	provider providers.ProviderInterface
}

// NewEth - Eth Module constructor to set the default provider
func NewEth(provider providers.ProviderInterface) *Eth {
	eth := new(Eth)
	eth.provider = provider
	return eth
}

func (eth *Eth) GasPrice() (*big.Int, error) {
	bodyByte, err := eth.provider.SendRequest(consts.METHOD_ETH_GAS_PRICE, []string{})
	if err != nil {
		return nil, err
	}
	Resp := new(resp.JsonRpcHexCommonResp)
	err = json.Unmarshal(bodyByte, Resp)
	if err != nil {
		return nil, err
	}
	v, isSuccess := big.NewInt(0).SetString(Resp.Result[2:], 16)
	if !isSuccess {
		return nil, errors.New(fmt.Sprintf("bit int transform err: %s", Resp.Result))
	}

	return v, err
}

func (eth *Eth) ProtocolVersion() (*big.Int, error) {
	bodyByte, err := eth.provider.SendRequest("eth_gasPrice", []string{})
	if err != nil {
		return nil, err
	}
	Resp := new(resp.EthProtocolVersionResp)
	err = json.Unmarshal(bodyByte, Resp)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyByte))
	v, isSuccess := big.NewInt(0).SetString(Resp.Result[2:], 16)
	if !isSuccess {
		return nil, errors.New(fmt.Sprintf("bit int transform err: %s", Resp.Result))
	}

	return v, err
}
