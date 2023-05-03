package client

import (
	"context"
	"encoding/json"
	"github.com/6boris/web3-go/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	"log"
	"math/big"
	"runtime"
	"time"
)

var MetricsWeb3RequestCounter metric.Int64Counter
var MetricsWeb3RequestHistogram metric.Int64Histogram

type Client struct {
	ethClient    *ethclient.Client
	rpcClient    *rpc.Client
	ClientID     string
	AppID        string
	Zone         string
	Cluster      string
	EthChainID   int64
	EthChainName string
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

func init() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := metricSdk.NewMeterProvider(metricSdk.WithReader(exporter))
	meter := provider.Meter("metrics", metric.WithInstrumentationVersion(runtime.Version()))
	m1, err := meter.Int64Counter("web3_abi_call", metric.WithDescription("Web3 Gateway abi call counter"))
	if err != nil {
		panic(err)
	}
	m2, err := meter.Int64Histogram("web3_abi_call", metric.WithDescription("Web3 Gateway abi call hist"))
	if err != nil {
		panic(err)
	}
	MetricsWeb3RequestCounter = m1
	MetricsWeb3RequestHistogram = m2

}
func (ec *Client) _beforeHooks(ctx context.Context, meta *Metadata) {

	meta.StartAt = time.Now()

}
func (ec *Client) _afterHooks(ctx context.Context, meta *Metadata) {
	MetricsWeb3RequestCounter.Add(ctx, 1, metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
		attribute.Key("chain_name").String(ec.EthChainName),
		attribute.Key("provider").String(ec.Provider),
		attribute.Key("abi_method").String(meta.AbiMethod),
		attribute.Key("status").String(meta.Status),
	))
	MetricsWeb3RequestHistogram.Record(ctx, time.Now().Sub(meta.StartAt).Milliseconds(), metric.WithAttributes(
		attribute.Key("client_id").String(ec.ClientID),
		attribute.Key("app_id").String(ec.AppID),
		attribute.Key("zone").String(ec.AppID),
		attribute.Key("cluster").String(ec.Cluster),
		attribute.Key("chain_id").Int64(ec.EthChainID),
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
