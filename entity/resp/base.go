package resp

type JsonRpcCommonResp struct {
	ID      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
}
type JsonRpcHexCommonResp struct {
	JsonRpcCommonResp
	Result string `json:"result"`
}
type JsonRpcHashCommonResp struct {
	JsonRpcCommonResp
	Result string `json:"result"`
}

type JsonRpcStringArrayCommonResp struct {
	JsonRpcCommonResp
	Result []string `json:"result"`
}

// eth_getBalance
type EthGetBalanceResp struct {
	JsonRpcCommonResp
	Result string `json:"result"`
}

// web3_clientVersion
type Web3ClientVersionResp struct {
	JsonRpcCommonResp
	Result string `json:"result"`
}

// eth_protocolVersion
type EthProtocolVersionResp struct {
	JsonRpcCommonResp
	Result string `json:"result"`
}
