package web3

import (
	"net/http"

	"web3.go/eth"
	"web3.go/providers"
	"web3.go/web3"
)

type Web3 struct {
	Provider providers.ProviderInterface
	Eth      *eth.Eth
	Web3     *web3.Web3
	client   http.Client
}

func NewWeb3(provider providers.ProviderInterface) *Web3 {
	w3 := new(Web3)
	w3.Provider = provider
	w3.Eth = eth.NewEth(provider)
	return w3
}
