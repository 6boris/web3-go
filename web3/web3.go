package web3

import (
	"encoding/json"
	"web3.go/consts"
	"web3.go/entity/resp"
	"web3.go/providers"
)

type Web3 struct {
	provider providers.ProviderInterface
}

func NewWeb3(provider providers.ProviderInterface) *Web3 {
	w3 := new(Web3)
	w3.provider = provider
	return w3
}

func (web3 *Web3) ClientVersion() (string, error) {
	var (
		err         error
		versionResp = new(resp.Web3ClientVersionResp)
		bodyByte    = make([]byte, 0)
	)

	bodyByte, err = web3.provider.SendRequest(consts.METHOD_WEB3_CLIENT_VERSION, []string{})
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bodyByte, versionResp)
	if err != nil {
		return "", err
	}
	return versionResp.Result, nil
}
