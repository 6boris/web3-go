package solana

import "github.com/shopspring/decimal"

type Context struct {
	Slot decimal.Decimal `json:"slot"`
}

type GetBalanceRequest struct {
	Account string `json:"account"`
}
type GetBalanceReply struct {
	Context Context         `json:"context"`
	Value   decimal.Decimal `json:"value"`
}

type GetTokenAccountBalanceRequest struct {
	Account string `json:"account"`
}
type GetTokenAccountBalanceReply struct {
	Context Context         `json:"context"`
	Value   decimal.Decimal `json:"value"`
}
