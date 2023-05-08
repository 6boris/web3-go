package consts

const (
	AbiCallStatusSuccess = "success"
	AbiCallStatusFail    = "fail"

	AbiMethodNetworkID                           = "net_version"
	AbiMethodClientVersion                       = "web3_clientVersion"
	AbiMethodEthChainID                          = "eth_chainId"
	AbiMethodEthGetBalance                       = "eth_getBalance"
	AbiMethodEthBlockNumber                      = "eth_blockNumber"
	AbiMethodEthGasPrice                         = "eth_gasPrice"
	AbiMethodEthGetBlockByHash                   = "eth_getBlockByHash"
	AbiMethodEthGetBlockByNumber                 = "eth_getBlockByNumber"
	AbiMethodEthGetBlockTransactionCountByHash   = "eth_getBlockTransactionCountByHash"
	AbiMethodEthGetBlockTransactionCountByNumber = "eth_getBlockTransactionCountByNumber"
	AbiMethodEthGetUncleCountByBlockHash         = "eth_getUncleCountByBlockHash"
	AbiMethodEthGetUncleCountByBlockNumber       = "eth_getUncleCountByBlockNumber"
	AbiMethodEthFeeHistory                       = "eth_feeHistory"
	AbiMethodEthSyncing                          = "eth_syncing"
	AbiMethodEthGetLogs                          = "eth_getLogs"
	AbiMethodEthGetStorageAt                     = "eth_getStorageAt"
	AbiMethodEthGetCode                          = "eth_getCode"
)
