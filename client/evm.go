package client

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/6boris/web3-go/consts"
	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/6boris/web3-go/pkg/otel"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type EvmClient struct {
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

func NewEvmClient(conf *clientModel.ConfEvmChainClient) (*EvmClient, error) {
	var err error
	ec := &EvmClient{
		ClientID:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		Provider:     conf.Provider,
		TransportURL: conf.TransportURL,
	}
	ec.ethClient, err = ethclient.Dial(conf.TransportURL)
	if err != nil {
		return nil, err
	}
	ec.rpcClient, err = rpc.Dial(conf.TransportURL)
	if err != nil {
		return nil, err
	}

	return ec, nil
}

func (ec *EvmClient) _beforeHooks(ctx context.Context, meta *clientModel.Metadata) {
	_ = ctx
	meta.StartAt = time.Now()
}
func (ec *EvmClient) _afterHooks(ctx context.Context, meta *clientModel.Metadata) {
	otel.MetricsWeb3RequestCounter.Add(ctx, 1, metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
		attribute.Key("chain_name").String(ec.EthChainName),
		attribute.Key("chain_env").String(ec.EthChainEnv),
		attribute.Key("provider").String(ec.Provider),
		attribute.Key("abi_method").String(meta.CallMethod),
		attribute.Key("status").String(meta.Status),
	))
	otel.MetricsWeb3RequestHistogram.Record(ctx, time.Since(meta.StartAt).Milliseconds(), metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
		attribute.Key("chain_env").String(ec.EthChainEnv),
		attribute.Key("chain_name").String(ec.EthChainName),
		attribute.Key("provider").String(ec.Provider),
		attribute.Key("abi_method").String(meta.CallMethod),
	))
}
func (ec *EvmClient) Close() {
	ec.ethClient.Close()
}

func (ec *EvmClient) BlockByHash(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	abiMethod := consts.EvmMethodBlockByHash
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BlockByHash(ctx, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) BlockByNumber(ctx context.Context, blockNumber *big.Int) (*types.Block, error) {
	abiMethod := consts.EvmMethodBlockByNumber
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.BlockByNumber(ctx, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error) {
	abiMethod := consts.EvmMethodHeaderByHash
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.HeaderByHash(ctx, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) HeaderByNumber(ctx context.Context, blockNumber *big.Int) (*types.Header, error) {
	abiMethod := consts.EvmMethodHeaderByNumber
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	abiMethod := consts.EvmMethodTransactionCount
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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
func (ec *EvmClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	abiMethod := consts.EvmMethodTransactionInBlock
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.TransactionInBlock(ctx, blockHash, index)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) TransactionByHash(ctx context.Context, blockHash common.Hash) (*types.Transaction, bool, error) {
	abiMethod := consts.EvmMethodTransactionByHash
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, isPending, err := ec.ethClient.TransactionByHash(ctx, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, isPending, err
}
func (ec *EvmClient) TransactionReceipt(ctx context.Context, blockHash common.Hash) (*types.Receipt, error) {
	abiMethod := consts.EvmMethodTransactionReceipt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.TransactionReceipt(ctx, blockHash)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

func (ec *EvmClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	abiMethod := consts.EvmMethodSendTransaction
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	err := ec.ethClient.SendTransaction(ctx, tx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return err
}

func (ec *EvmClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	abiMethod := consts.EvmMethodBalanceAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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
func (ec *EvmClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	abiMethod := consts.EvmMethodStorageAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.StorageAt(ctx, account, key, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	abiMethod := consts.EvmMethodCodeAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.CodeAt(ctx, account, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	abiMethod := consts.EvmMethodNonceAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.NonceAt(ctx, account, blockNumber)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}

func (ec *EvmClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.EvmMethodSuggestGasPrice
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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
func (ec *EvmClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.EvmMethodSuggestGasTipCap
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	abiMethod := consts.EvmMethodFeeHistory
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	abiMethod := consts.EvmMethodEstimateGas
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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

func (ec *EvmClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	abiMethod := consts.EvmMethodPendingBalanceAtp
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.PendingBalanceAt(ctx, account)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	abiMethod := consts.EvmMethodPendingStorageAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.PendingStorageAt(ctx, account, key)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	abiMethod := consts.EvmMethodPendingCodeAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.PendingCodeAt(ctx, account)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	abiMethod := consts.EvmMethodPendingNonceAt
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.PendingNonceAt(ctx, account)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	abiMethod := consts.EvmMethodPendingTransactionCount
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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

func (ec *EvmClient) BlockNumber(ctx context.Context) (uint64, error) {
	abiMethod := consts.EvmMethodBlockNumber
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
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
func (ec *EvmClient) ChainID(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.EvmMethodChainID
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.ChainID(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
func (ec *EvmClient) NetworkID(ctx context.Context) (*big.Int, error) {
	abiMethod := consts.EvmMethodNetworkID
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	result, err := ec.ethClient.NetworkID(ctx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return result, err
}
