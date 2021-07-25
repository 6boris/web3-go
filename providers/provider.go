package providers

type JsonRPCReq struct {
	ID      int         `json:"id"`
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type ProviderInterface interface {
	SendRequest(method string, params interface{}) ([]byte, error)
}
