package solana

import "github.com/shopspring/decimal"

type ContextItem struct {
	ApiVersion string `json:"apiVersion"`
	Slot       int64  `json:"slot"`
}
type AccountInfoItem struct {
	Lamports   decimal.Decimal `json:"lamports"`
	Owner      string          `json:"owner"`
	Executable bool            `json:"executable"`
	RentEpoch  decimal.Decimal `json:"rentEpoch"`
	Space      decimal.Decimal `json:"space"`
	Data       []string        `json:"data"`
}

type GetAccountInfoRequest struct {
	Account string `json:"account"`
}
type GetAccountInfoReply struct {
	Context *ContextItem     `json:"context"`
	Value   *AccountInfoItem `json:"value"`
}

type GetVersionRequest struct{}
type GetVersionReply struct {
	FeatureSet string `json:"feature_set"`
	SolanaCore string `json:"solana_core"`
}

type GetBalanceRequest struct {
	Account string `json:"account"`
}
type GetBalanceReply struct {
	Context *ContextItem    `json:"context"`
	Value   decimal.Decimal `json:"value"`
}

type GetTokenAccountBalanceRequest struct {
	Account string `json:"account"`
}
type GetTokenAccountBalanceReply struct {
	Context        *ContextItem    `json:"context"`
	Amount         decimal.Decimal `json:"amount"`
	Decimals       decimal.Decimal `json:"decimals"`
	UIAmount       decimal.Decimal `json:"ui_amount"`
	UIAmountString decimal.Decimal `json:"ui_amount_string"`
}

type GetBlockHeightRequest struct {
}
type GetBlockHeightReply struct {
	BlockHeight int64 `json:"block_height"`
}

type GetBlockRequest struct {
	Slot               int64  `json:"slot"`
	Encoding           string `json:"encoding"`
	TransactionDetails string `json:"transaction_details"`
	Rewards            bool   `json:"rewards"`
}
type GetBlockReply struct {
	Context *ContextItem    `json:"context"`
	Value   decimal.Decimal `json:"value"`
}

type ClusterNodesItem struct {
	Gossip  string `json:"gossip"`
	PubKey  string `json:"pubkey"`
	RPC     string `json:"rpc"`
	TPU     string `json:"tpu"`
	Version string `json:"version"`
}
