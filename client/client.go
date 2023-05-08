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
