package client

import (
	"context"
	"encoding/json"
	"github.com/6boris/web3-go/consts"
	"github.com/6boris/web3-go/pkg/otel"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"math/big"
	"time"
)

type Client struct {
	ethClient    *ethclient.Client
	rpcClient    *rpc.Client
	ClientID     string
	AppID        string
	Zone         string
	Cluster      string
	EthChainID   int64
	EthChainName string
	EthChainEnv  string
	Provider     string
	TransportURL string
}

func NewClient(conf *ConfClient) (*Client, error) {
	var err error
	tmpBytes, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}
	client := &Client{
		ClientID:     common.HexToAddress(common.Bytes2Hex(tmpBytes)).String(),
		Provider:     conf.Provider,
		TransportURL: conf.TransportURL,
	}
	client.ethClient, err = ethclient.Dial(conf.TransportURL)
	if err != nil {
		return nil, err
	}
	client.rpcClient, err = rpc.Dial(conf.TransportURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (ec *Client) _beforeHooks(ctx context.Context, meta *Metadata) {
	meta.StartAt = time.Now()

}
func (ec *Client) _afterHooks(ctx context.Context, meta *Metadata) {
	otel.MetricsWeb3RequestCounter.Add(ctx, 1, metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
		attribute.Key("chain_name").String(ec.EthChainName),
		attribute.Key("chain_env").String(ec.EthChainEnv),
		attribute.Key("provider").String(ec.Provider),
		attribute.Key("abi_method").String(meta.AbiMethod),
		attribute.Key("status").String(meta.Status),
	))
	otel.MetricsWeb3RequestHistogram.Record(ctx, time.Now().Sub(meta.StartAt).Milliseconds(), metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
		attribute.Key("chain_env").String(ec.EthChainEnv),
		attribute.Key("chain_name").String(ec.EthChainName),
		attribute.Key("provider").String(ec.Provider),
		attribute.Key("abi_method").String(meta.AbiMethod),
	))
}

func (ec *Client) Close() {
	ec.ethClient.Close()
	return
}

// Account Information

/*
BalanceAt eth_getBalance

	Returns the balance of the account of a given address.
*/
func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	abiMethod := consts.AbiMethodEthGetBalance
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
StorageAt eth_getStorageAt

	Returns the value from a storage position at a given address, or in other words, returns the state of the contract's storage,
	which may not be exposed via the contract's methods.
*/
func (ec *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	abiMethod := consts.AbiMethodEthGetStorageAt
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result []byte
	var err error
	result, err = ec.ethClient.StorageAt(ctx, account, key, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
CodeAt eth_getCode

	Returns code at a given address.
*/
func (ec *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	abiMethod := consts.AbiMethodEthGetCode
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result []byte
	var err error
	result, err = ec.ethClient.CodeAt(ctx, account, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

// Chain Information

/*
ChainID eth_chainId

	The chain ID returned should always correspond to the information in the current known head block.
	This ensures that caller of this RPC method can always use the retrieved information to sign transactions built on top of the head.
*/
func (ec *Client) ChainID(ctx context.Context) (*big.Int, error) {
	meta := &Metadata{AbiMethod: consts.AbiMethodEthChainID}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()

	result, err := ec.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}

/*
ClientVersion web3_clientVersion

	Returns the current client version.
*/
func (ec *Client) ClientVersion(ctx context.Context) (string, error) {
	var version string
	var err error
	abiMethod := consts.AbiMethodClientVersion
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()

	err = ec.rpcClient.CallContext(ctx, &version, abiMethod)
	if err != nil {
		if err != nil {
			meta.Status = consts.AbiCallStatusFail
		}
		return "", err
	}
	return version, err
}

/*
NetworkID net_version

	Returns the current network id.
*/
func (ec *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.AbiMethodNetworkID
	meta := &Metadata{AbiMethod: abiMethod}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}

// Gas Information

/*
SyncProgress net_version
*/
func (ec *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	abiMethod := consts.AbiMethodEthSyncing
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result *ethereum.SyncProgress
	var err error
	result, err = ec.ethClient.SyncProgress(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
SuggestGasPrice eth_gasPrice

	Returns the current price per gas in wei.
*/
func (ec *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.AbiMethodEthGasPrice
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
EstimateGas eth_gasPrice

	Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.
	The transaction will not be added to the blockchain.
*/
func (ec *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	abiMethod := consts.AbiMethodEthEstimateGas
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.EstimateGas(ctx, msg)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

// Blocks

/*
BlockNumber eth_blockNumber

	Returns the number of the most recent block.
*/
func (ec *Client) BlockNumber(ctx context.Context) (uint64, error) {
	abiMethod := consts.AbiMethodEthBlockNumber
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BlockNumber(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
BlockByNumber eth_getBlockByNumber

	Returns information about a block by block number.
*/
func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	abiMethod := consts.AbiMethodEthGetBlockByNumber
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BlockByNumber(ctx, number)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
TransactionCount eth_getBlockTransactionCountByHash

	Returns the number of transactions in a block matching the given block hash.
*/
func (ec *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	abiMethod := consts.AbiMethodEthGetBlockTransactionCountByHash
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.TransactionCount(ctx, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
BlockByHash eth_getBlockByHash
*/
func (ec *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	abiMethod := consts.AbiMethodEthGetBlockByHash
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BlockByHash(ctx, hash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
PendingTransactionCount eth_getBlockTransactionCountByNumber
*/
func (ec *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	abiMethod := consts.AbiMethodEthGetBlockTransactionCountByNumber
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.PendingTransactionCount(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

// Event Logs

/*
FeeHistory eth_feeHistory
*/
func (ec *Client) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	abiMethod := consts.AbiMethodEthFeeHistory
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result *ethereum.FeeHistory
	var err error
	result, err = ec.ethClient.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
FilterLogs eth_getLogs
*/
func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	abiMethod := consts.AbiMethodEthGetLogs
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result []types.Log
	var err error
	result, err = ec.ethClient.FilterLogs(ctx, q)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

// Uncle Blocks

/*
GetUncleCountByBlockHash eth_getUncleCountByBlockHash

	Returns the number of uncles in a block matching the given block hash.
*/
func (ec *Client) GetUncleCountByBlockHash(ctx context.Context, blockHash common.Hash) (string, error) {
	abiMethod := consts.AbiMethodEthGetUncleCountByBlockHash
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result string
	var err error
	err = ec.rpcClient.CallContext(ctx, &result, abiMethod, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

/*
GetUncleCountByBlockNumber eth_getUncleCountByBlockNumber

	Returns the number of uncles in a block matching the give block number.
*/
func (ec *Client) GetUncleCountByBlockNumber(ctx context.Context, number *big.Int) (string, error) {
	abiMethod := consts.AbiMethodEthGetUncleCountByBlockNumber
	meta := &Metadata{AbiMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	var result string
	var err error
	err = ec.rpcClient.CallContext(ctx, &result, abiMethod, number)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
