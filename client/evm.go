package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/6boris/web3-go/consts"
	"github.com/6boris/web3-go/erc/erc20"
	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/6boris/web3-go/pkg/otel"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type EvmClient struct {
	ethClient     *ethclient.Client
	rpcClient     *rpc.Client
	_signers      []*clientModel.ConfEvmChainSigner
	_gasLimitMax  decimal.Decimal
	_gasFeeRate   decimal.Decimal
	_gasLimitRate decimal.Decimal
	_clientID     string
	_appID        string
	_zone         string
	_cluster      string
	_ethChainID   int64
	_ethChainName string
	_ethChainEnv  string
	_provider     string
	_transportURL string
}

func NewEvmClient(conf *clientModel.ConfEvmChainClient) (*EvmClient, error) {
	var err error
	ec := &EvmClient{
		_clientID:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		_provider:     conf.Provider,
		_transportURL: conf.TransportURL,
		_gasFeeRate:   conf.GasFeeRate,
		_gasLimitRate: conf.GasLimitRate,
		_gasLimitMax:  conf.GasLimitMax,
	}

	if ec._gasFeeRate == decimal.Zero {
		ec._gasFeeRate = decimal.NewFromFloat(1.1)
	}
	if ec._gasLimitRate == decimal.Zero {
		ec._gasLimitRate = decimal.NewFromFloat(2)
	}
	if ec._gasLimitMax == decimal.Zero {
		ec._gasLimitMax = decimal.NewFromFloat(30000000)
	}
	ec._signers = conf.Signers

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

func (ec *EvmClient) _getSinnerPrivateKey(account common.Address) (*ecdsa.PrivateKey, error) {
	for _, v := range ec._signers {
		if v.PublicAddress.String() == account.String() {
			return v.PrivateKey, nil
		}
	}
	return nil, errors.New("signer not config")
}
func (ec *EvmClient) _beforeHooks(ctx context.Context, meta *clientModel.Metadata) {
	_ = ctx
	meta.StartAt = time.Now()
}
func (ec *EvmClient) _afterHooks(ctx context.Context, meta *clientModel.Metadata) {
	otel.MetricsWeb3RequestCounter.Add(ctx, 1, metric.WithAttributes(
		attribute.Key("client_id").String(ec._clientID),
		attribute.Key("app_id").String(ec._appID),
		attribute.Key("zone").String(ec._appID),
		attribute.Key("cluster").String(ec._cluster),
		attribute.Key("chain_id").Int64(ec._ethChainID),
		attribute.Key("chain_name").String(ec._ethChainName),
		attribute.Key("chain_env").String(ec._ethChainEnv),
		attribute.Key("provider").String(ec._provider),
		attribute.Key("abi_method").String(meta.CallMethod),
		attribute.Key("status").String(meta.Status),
	))
	otel.MetricsWeb3RequestHistogram.Record(ctx, time.Since(meta.StartAt).Milliseconds(), metric.WithAttributes(
		attribute.Key("client_id").String(ec._clientID),
		attribute.Key("app_id").String(ec._appID),
		attribute.Key("zone").String(ec._appID),
		attribute.Key("cluster").String(ec._cluster),
		attribute.Key("chain_id").Int64(ec._ethChainID),
		attribute.Key("chain_env").String(ec._ethChainEnv),
		attribute.Key("chain_name").String(ec._ethChainName),
		attribute.Key("provider").String(ec._provider),
		attribute.Key("abi_method").String(meta.CallMethod),
	))
}
func (ec *EvmClient) _getTransactOpts(ctx context.Context, signer common.Address, to common.Address, dataHex string) (*bind.TransactOpts, error) {
	msgSignerPk, err := ec._getSinnerPrivateKey(signer)
	if err != nil {
		return nil, err
	}
	nonce, err := ec.PendingNonceAt(context.Background(), signer)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	estimateGas, err := ec.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:  signer,
		To:    &to,
		Gas:   uint64(ec._gasLimitMax.BigInt().Int64()),
		Value: big.NewInt(0),
		Data:  common.Hex2Bytes(dataHex),
	})
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(msgSignerPk, chainID)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.GasLimit = 0
	opts.GasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(ec._gasFeeRate).BigInt()
	opts.Context = ctx
	if estimateGas > 0 && !ec._gasLimitRate.IsZero() {
		opts.GasLimit = decimal.NewFromInt(int64(estimateGas)).Mul(ec._gasLimitRate).BigInt().Uint64()
	}
	if !ec._gasLimitMax.IsZero() && opts.GasLimit > uint64(ec._gasLimitMax.BigInt().Int64()) {
		opts.GasLimit = uint64(ec._gasLimitMax.BigInt().Int64())
	}
	return opts, nil
}
func (ec *EvmClient) _getCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	return opts, nil
}

func (ec *EvmClient) Close() {
	ec.ethClient.Close()
	ec.rpcClient.Close()
}
func (ec *EvmClient) GetAllSinners() []common.Address {
	data := make([]common.Address, 0)
	for _, v := range ec._signers {
		data = append(data, v.PublicAddress)
	}
	return data
}
func (ec *EvmClient) GetTransportURL() string {
	return ec._transportURL
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
func (ec *EvmClient) SendTransactionSimple(ctx context.Context, signer common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	abiMethod := consts.EvmMethodSendTransaction
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	msgSignerPk, err := ec._getSinnerPrivateKey(signer)
	if err != nil {
		return nil, err
	}
	opts, err := ec._getTransactOpts(ctx, signer, to, "0x")
	if err != nil {
		return nil, err
	}
	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignNewTx(msgSignerPk, types.NewEIP155Signer(chainID), &types.LegacyTx{
		To:       &to,
		Nonce:    opts.Nonce.Uint64(),
		Value:    value,
		Gas:      opts.GasLimit,
		GasPrice: opts.GasPrice,
	})
	if err != nil {
		return nil, err
	}
	err = ec.ethClient.SendTransaction(ctx, signedTx)
	if err != nil {
		meta.Status = consts.AbiCallStatusFail
	}
	return signedTx, err
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

func (ec *EvmClient) ERC20Name(ctx context.Context, token common.Address) (string, error) {
	abiMethod := consts.EvmErc20MethodName
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return "", err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return "", err
	}
	callResp, err := inst.Name(opts)
	if err != nil {
		return "", err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20Symbol(ctx context.Context, token common.Address) (string, error) {
	abiMethod := consts.EvmErc20MethodSymbol
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return "", err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return "", err
	}
	callResp, err := inst.Symbol(opts)
	if err != nil {
		return "", err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20Decimals(ctx context.Context, token common.Address) (uint8, error) {
	abiMethod := consts.EvmErc20MethodDecimals
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return 0, err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return 0, err
	}
	callResp, err := inst.Decimals(opts)
	if err != nil {
		return 0, err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20BalanceOf(ctx context.Context, token common.Address, account common.Address) (*big.Int, error) {
	abiMethod := consts.EvmErc20MethodBalanceOf
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return big.NewInt(0), err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return big.NewInt(0), err
	}
	callResp, err := inst.BalanceOf(opts, account)
	if err != nil {
		return big.NewInt(0), err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20TotalSupply(ctx context.Context, token common.Address) (*big.Int, error) {
	abiMethod := consts.EvmErc20MethodTotalSupply
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return big.NewInt(0), err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return big.NewInt(0), err
	}
	callResp, err := inst.TotalSupply(opts)
	if err != nil {
		return big.NewInt(0), err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20Transfer(ctx context.Context, token common.Address, signer common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	abiMethod := consts.EvmErc20MethodTransfer
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return nil, err
	}
	abi, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	abiData, err := abi.Pack("transfer", to, value)
	if err != nil {
		return nil, err
	}

	opts, err := ec._getTransactOpts(ctx, signer, token, common.Bytes2Hex(abiData))
	if err != nil {
		return nil, err
	}
	callResp, err := inst.Transfer(opts, to, value)
	if err != nil {
		return nil, err
	}

	return callResp, nil
}
func (ec *EvmClient) ERC20Approve(ctx context.Context, token common.Address, signer common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	abiMethod := consts.EvmErc20MethodApprove
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return nil, err
	}
	abi, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	abiData, err := abi.Pack("transfer", to, value)
	if err != nil {
		return nil, err
	}
	opts, err := ec._getTransactOpts(ctx, signer, token, common.Bytes2Hex(abiData))
	if err != nil {
		return nil, err
	}
	callResp, err := inst.Approve(opts, to, value)
	if err != nil {
		return nil, err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20IncreaseAllowance(ctx context.Context, token common.Address, signer common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	abiMethod := consts.EvmErc20MethodIncreaseAllowance
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return nil, err
	}
	abi, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	abiData, err := abi.Pack("transfer", to, value)
	if err != nil {
		return nil, err
	}
	opts, err := ec._getTransactOpts(ctx, signer, token, common.Bytes2Hex(abiData))
	if err != nil {
		return nil, err
	}
	callResp, err := inst.IncreaseAllowance(opts, to, value)
	if err != nil {
		return nil, err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20DecreaseAllowance(ctx context.Context, token common.Address, signer common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	abiMethod := consts.EvmErc20MethodDecreaseAllowance
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return nil, err
	}
	abi, err := erc20.ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	abiData, err := abi.Pack("transfer", to, value)
	if err != nil {
		return nil, err
	}
	opts, err := ec._getTransactOpts(ctx, signer, token, common.Bytes2Hex(abiData))
	if err != nil {
		return nil, err
	}
	callResp, err := inst.DecreaseAllowance(opts, to, value)
	if err != nil {
		return nil, err
	}
	return callResp, nil
}
func (ec *EvmClient) ERC20Allowance(ctx context.Context, token common.Address, owner common.Address, spender common.Address) (*big.Int, error) {
	abiMethod := consts.EvmErc20MethodAllowance
	meta := &clientModel.Metadata{CallMethod: abiMethod, Status: consts.AbiCallStatusSuccess}
	ec._beforeHooks(ctx, meta)
	defer func() {
		ec._afterHooks(ctx, meta)
	}()
	inst, err := erc20.NewERC20(token, ec.ethClient)
	if err != nil {
		return big.NewInt(0), err
	}
	opts, err := ec._getCallOpts(ctx)
	if err != nil {
		return big.NewInt(0), err
	}
	callResp, err := inst.Allowance(opts, owner, spender)
	if err != nil {
		return big.NewInt(0), err
	}
	return callResp, nil
}
