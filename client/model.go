package client

import "time"

type ConfPool struct {
	AppID   string               `yaml:"app_id" json:"app_id"`
	Zone    string               `yaml:"zone" json:"zone"`
	Cluster string               `yaml:"cluster" json:"cluster"`
	Chains  map[int64]*ConfChain `yaml:"chains" json:"chains"`
}
type ConfChain struct {
	ChainID         int64         `yaml:"chain_id" json:"chain_id"`
	ChainName       string        `yaml:"chain_name" json:"chain_name"`
	OfficialWebsite string        `yaml:"official_website_url" json:"official_website"`
	ExplorerURL     string        `yaml:"explorer_url" json:"explorer_url"`
	Faucets         []string      `yaml:"faucets" json:"faucets"`
	Clients         []*ConfClient `yaml:"clients" json:"clients"`
}

type ConfClient struct {
	Provider        string `yaml:"provider" json:"provider"`
	ProviderWebsite string `yaml:"provider" json:"provider_website"`
	TransportSchema string `yaml:"transport_schema" json:"transport_schema"`
	TransportURL    string `yaml:"transport_url" json:"transport_url"`
}

type Metadata struct {
	AbiMethod string    `yaml:"abi_method" json:"abi_method"`
	StartAt   time.Time `yaml:"transport_url" json:"transport_url"`
	Status    string    `yaml:"status" json:"status"`
}

type EthCallProxyRequest struct {
	ChainID int64         `json:"CHainID"`
	ID      int64         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type EthCallProxyReply struct {
	ID      int64       `json:"id"`
	JsonRpc string      `json:"json_rpc"`
	Result  interface{} `json:"result"`
}
type ErrReply struct {
	Code     int64             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
