package client

import (
	"github.com/6boris/web3-go/model/client"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
)

type Pool struct {
	conf *client.ConfPool
	// clients        map[int64]map[string]*Client
	breakerGroup   *circuitbreaker.CircuitBreaker
	_evmClients    map[int64]map[string]*EvmClient
	_solanaClients map[string]*SolanaClient
}

// func init() {
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

func NewPool(conf *client.ConfPool) *Pool {
	p := &Pool{
		conf:           conf,
		_solanaClients: map[string]*SolanaClient{},
	}
	b := sre.NewBreaker()
	p.breakerGroup = &b

	p._evmClients = make(map[int64]map[string]*EvmClient, 0)
	for _, chain := range conf.EvmChains {
		if _, ok := p._evmClients[chain.ChainID]; !ok {
			p._evmClients[chain.ChainID] = make(map[string]*EvmClient, 0)
		}
		for _, c := range chain.Clients {
			if c.TransportSchema == "https" {
				tmpC, err := NewEvmClient(c)
				tmpC.AppID = conf.AppID
				tmpC.Zone = conf.Zone
				tmpC.Cluster = conf.Cluster
				tmpC.EthChainID = chain.ChainID
				tmpC.EthChainName = chain.ChainName
				tmpC.EthChainEnv = chain.ChainEnv
				if err != nil {
					panic(err)
				} else {
					p._evmClients[chain.ChainID][tmpC.ClientID] = tmpC
				}
			}
		}
	}
	for _, c := range p.conf.SolanaChains {
		loopClient, loopErr := NewSolanaClient(c)
		if loopErr == nil {
			p._solanaClients[loopClient.ClientID] = loopClient
		}
	}
	return p
}

func (p *Pool) GetEvmClient(chainID int64) *EvmClient {
	if clientMap, ok := p._evmClients[chainID]; ok {
		if len(clientMap) > 0 {
			for _, val := range clientMap {
				return val
			}
		}
	}
	return nil
}
func (p *Pool) GetSolanaClient(chainEnv string) *SolanaClient {
	if len(p._solanaClients) == 0 {
		return nil
	}
	for _, v := range p._solanaClients {
		if v.ChainEnv == chainEnv {
			return v
		}
	}
	return nil
}
