package client

import (
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
)

type Pool struct {
	conf         *ConfPool
	clients      map[int64]map[string]*Client
	breakerGroup *circuitbreaker.CircuitBreaker
}

func NewPool(conf *ConfPool) *Pool {
	p := &Pool{
		conf: conf,
	}
	b := sre.NewBreaker()
	p.breakerGroup = &b

	p.clients = make(map[int64]map[string]*Client, 0)
	for _, chain := range conf.Chains {
		if _, ok := p.clients[chain.ChainID]; !ok {
			p.clients[chain.ChainID] = make(map[string]*Client, 0)
		}
		for _, client := range chain.Clients {
			if client.TransportSchema == "https" {
				tmpC, err := NewClient(client)
				tmpC.AppID = conf.AppID
				tmpC.Zone = conf.Zone
				tmpC.Cluster = conf.Cluster
				tmpC.EthChainID = chain.ChainID
				tmpC.EthChainName = chain.ChainName
				tmpC.EthChainName = chain.ChainName
				if err != nil {
				} else {
					p.clients[chain.ChainID][tmpC.ClientID] = tmpC
				}
			}
		}
	}
	return p
}

func (p *Pool) GetClient(chainID int64) *Client {
	if clientMap, ok := p.clients[chainID]; ok {
		if len(clientMap) > 0 {
			for _, val := range clientMap {
				return val
			}
		}
	}
	return &Client{}
}
