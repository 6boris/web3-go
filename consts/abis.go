package consts

const (
	AbiCallStatusSuccess = "SUCCESS"
	AbiCallStatusFail    = "FAILED"

	EvmMethodBlockByHash             = "EVM_BlockByHash"
	EvmMethodBlockByNumber           = "EVM_BlockByNumber"
	EvmMethodHeaderByHash            = "EVM_HeaderByHash"
	EvmMethodHeaderByNumber          = "EVM_HeaderByNumber"
	EvmMethodTransactionCount        = "EVM_TransactionCount"
	EvmMethodTransactionInBlock      = "EVM_TransactionInBlock"
	EvmMethodTransactionByHash       = "EVM_TransactionByHash"
	EvmMethodTransactionReceipt      = "EVM_TransactionReceipt"
	EvmMethodSendTransaction         = "EVM_SendTransaction"
	EvmMethodBalanceAt               = "EVM_BalanceAt"
	EvmMethodStorageAt               = "EVM_StorageAt"
	EvmMethodCodeAt                  = "EVM_CodeAt"
	EvmMethodNonceAt                 = "EVM_NonceAt"
	EvmMethodSuggestGasPrice         = "EVM_SuggestGasPrice"
	EvmMethodSuggestGasTipCap        = "EVM_SuggestGasTipCap"
	EvmMethodFeeHistory              = "EVM_FeeHistory"
	EvmMethodEstimateGas             = "EVM_EstimateGas"
	EvmMethodPendingBalanceAtp       = "EVM_PendingBalanceAt"
	EvmMethodPendingStorageAt        = "EVM_PendingStorageAt"
	EvmMethodPendingCodeAt           = "EVM_PendingCodeAt"
	EvmMethodPendingNonceAt          = "EVM_PendingNonceAt"
	EvmMethodPendingTransactionCount = "EVM_PendingTransactionCount"
	EvmMethodBlockNumber             = "EVM_BlockNumber"
	EvmMethodChainID                 = "EVM_ChainID"
	EvmMethodNetworkID               = "EVM_NetworkID"

	SolanaMethodGetBalance             = "SOLANA_GetBalance"
	SolanaMethodGetTokenAccountBalance = "SOLANA_GetTokenAccountBalance"
)
