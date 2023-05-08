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

//func init() {
//	//exporter, err := prometheus.New()
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//provider := metricSdk.NewMeterProvider(metricSdk.WithReader(exporter))
//	//meter := provider.Meter("metrics", metric.WithInstrumentationVersion(runtime.Version()))
//	m1, err := global.Meter("web3.go").Int64Counter("web3_abi_call", metric.WithDescription("Web3 Gateway abi call counter"))
//	if err != nil {
//		panic(err)
//	}
//	m2, err := global.Meter("web3.go").Int64Histogram("web3_abi_call", metric.WithDescription("Web3 Gateway abi call hist"))
//	if err != nil {
//		panic(err)
//	}
//	MetricsWeb3RequestCounter = m1
//	MetricsWeb3RequestHistogram = m2
//
//}

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
				tmpC.EthChainEnv = chain.ChainEnv
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
