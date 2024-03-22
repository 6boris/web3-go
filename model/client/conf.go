package client

import (
	"time"
)

type ConfPool struct {
	AppID        string                      `yaml:"app_id" json:"app_id"`
	Zone         string                      `yaml:"zone" json:"zone"`
	Cluster      string                      `yaml:"cluster" json:"cluster"`
	EvmChains    map[int64]*ConfEvmChainInfo `yaml:"evm_chains" json:"evm_chains"`
	SolanaChains []*ConfSolanaClient         `yaml:"solana_chains" json:"solana_chains"`
}
type ConfEvmChainInfo struct {
	ChainID         int64                 `yaml:"chain_id" json:"chain_id"`
	ChainName       string                `yaml:"chain_name" json:"chain_name"`
	ChainEnv        string                `yaml:"chain_env" json:"chain_env"`
	OfficialWebsite string                `yaml:"official_website_url" json:"official_website"`
	ExplorerURL     string                `yaml:"explorer_url" json:"explorer_url"`
	Faucets         []string              `yaml:"faucets" json:"faucets"`
	Clients         []*ConfEvmChainClient `yaml:"clients" json:"clients"`
}
type ConfEvmChainClient struct {
	ClientID        string `yaml:"client_id" json:"client_id"`
	Provider        string `yaml:"provider" json:"provider"`
	ProviderWebsite string `yaml:"provider_website" json:"provider_website"`
	TransportSchema string `yaml:"transport_schema" json:"transport_schema"`
	TransportURL    string `yaml:"transport_url" json:"transport_url"`
}
type ConfEvmChain struct {
	ChainID         int64         `yaml:"chain_id" json:"chain_id"`
	ChainName       string        `yaml:"chain_name" json:"chain_name"`
	ChainEnv        string        `yaml:"chain_env" json:"chain_env"`
	OfficialWebsite string        `yaml:"official_website_url" json:"official_website"`
	ExplorerURL     string        `yaml:"explorer_url" json:"explorer_url"`
	Faucets         []string      `yaml:"faucets" json:"faucets"`
	Clients         []*ConfClient `yaml:"clients" json:"clients"`
}

type ConfClient struct {
	ClientID        string `yaml:"client_id" json:"client_id"`
	Provider        string `yaml:"provider" json:"provider"`
	ProviderWebsite string `yaml:"provider_website" json:"provider_website"`
	TransportSchema string `yaml:"transport_schema" json:"transport_schema"`
	TransportURL    string `yaml:"transport_url" json:"transport_url"`
}
type ConfEvmClient struct {
	ClientID        string `yaml:"client_id" json:"client_id"`
	Provider        string `yaml:"provider" json:"provider"`
	ProviderWebsite string `yaml:"provider_website" json:"provider_website"`
	TransportSchema string `yaml:"transport_schema" json:"transport_schema"`
	TransportURL    string `yaml:"transport_url" json:"transport_url"`
}
type ConfSolanaClient struct {
	ClientID        string `yaml:"client_id" json:"client_id"`
	Provider        string `yaml:"provider" json:"provider"`
	TransportSchema string `yaml:"transport_schema" json:"transport_schema"`
	ChainEnv        string `yaml:"chain_env" json:"chain_env"`
	TransportURL    string `yaml:"transport_url" json:"transport_url"`
}

type Metadata struct {
	CallMethod string    `yaml:"call_method" json:"call_method"`
	StartAt    time.Time `yaml:"start_at" json:"start_at"`
	Status     string    `yaml:"status" json:"status"`
}

type EvmCallProxyRequest struct {
	ChainID int64         `json:"chain_id"`
	ID      int64         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
type SolanaCallProxyRequest struct {
	ChainEnv string        `json:"chain_env"`
	ID       int64         `json:"id"`
	JsonRpc  string        `json:"jsonrpc"`
	Method   string        `json:"method"`
	Params   []interface{} `json:"params"`
}

type EvmCallProxyReply struct {
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
