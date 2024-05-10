package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/6boris/web3-go/model/solana"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

type SolanaClient struct {
	HttpClient   *req.Client
	ClientID     string
	AppID        string
	Zone         string
	Cluster      string
	ChainEnv     string
	Provider     string
	TransportURL string
}

func NewSolanaClient(conf *clientModel.ConfSolanaClient) (*SolanaClient, error) {
	client := &SolanaClient{
		ClientID:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		Provider:     conf.Provider,
		TransportURL: conf.TransportURL,
		ChainEnv:     conf.ChainEnv,
	}
	client.HttpClient = req.C().
		SetBaseURL(conf.TransportURL).
		WrapRoundTripFunc(func(rt req.RoundTripper) req.RoundTripFunc {
			return func(req *req.Request) (resp *req.Response, err error) {
				// before request
				// ...
				resp, err = rt.RoundTrip(req)
				// after response
				// ...
				return
			}
		}).
		SetTimeout(20 * time.Second)
	if conf.IsDev {
		client.HttpClient.DevMode()
	}

	return client, nil
}

func (sc *SolanaClient) GetAccountInfo(ctx context.Context, request *solana.GetAccountInfoRequest) (*solana.GetAccountInfoReply, error) {
	reply := &solana.GetAccountInfoReply{}
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getAccountInfo",
			"params": []interface{}{
				request.Account,
				map[string]string{
					"encoding": "base58",
				},
			},
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	err = json.Unmarshal([]byte(gjson.GetBytes(response.Bytes(), "result").String()), &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (sc *SolanaClient) GetVersion(ctx context.Context) (*solana.GetVersionReply, error) {
	reply := &solana.GetVersionReply{}
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getVersion",
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	reply.FeatureSet = gjson.GetBytes(response.Bytes(), "result.feature-set").String()
	reply.SolanaCore = gjson.GetBytes(response.Bytes(), "result.solana-core").String()
	return reply, nil
}
func (sc *SolanaClient) GetBalance(ctx context.Context, request *solana.GetBalanceRequest) (*solana.GetBalanceReply, error) {
	//callMethod := consts.SolanaMethodGetBalance
	reply := &solana.GetBalanceReply{}
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getBalance",
			"params": []string{request.Account},
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	err = json.Unmarshal([]byte(gjson.GetBytes(response.Bytes(), "result.context").String()), &reply.Context)
	if err != nil {
		return nil, err
	}
	reply.Value, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value").String())
	if err != nil {
		return nil, err
	}
	reply.Value = reply.Value.Div(decimal.New(1, 9))
	return reply, nil
}
func (sc *SolanaClient) GetTokenAccountBalance(ctx context.Context, request *solana.GetTokenAccountBalanceRequest) (*solana.GetTokenAccountBalanceReply, error) {
	reply := &solana.GetTokenAccountBalanceReply{}
	response, err := sc.HttpClient.
		SetBaseURL(sc.TransportURL).
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getTokenAccountBalance",
			"params": []string{request.Account},
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	err = json.Unmarshal([]byte(gjson.GetBytes(response.Bytes(), "result.context").String()), &reply.Context)
	if err != nil {
		return nil, err
	}
	reply.Amount, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value.amount").String())
	if err != nil {
		return nil, err
	}
	reply.Decimals, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value.decimals").String())
	if err != nil {
		return nil, err
	}
	reply.UIAmount, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value.uiAmount").String())
	if err != nil {
		return nil, err
	}
	reply.UIAmountString, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value.uiAmountString").String())
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (sc *SolanaClient) GetBlockHeight(ctx context.Context) (int64, error) {
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getBlockHeight",
		}).
		Post("")
	if err != nil {
		return 0, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return 0, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	return gjson.GetBytes(response.Bytes(), "result").Int(), nil
}
func (sc *SolanaClient) GetBlockTime(ctx context.Context, blockNumber int64) (int64, error) {
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getBlockTime",
			"params": []interface{}{blockNumber},
		}).
		Post("")
	if err != nil {
		return 0, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return 0, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	return gjson.GetBytes(response.Bytes(), "result").Int(), nil
}
func (sc *SolanaClient) GetBlock(ctx context.Context, request *solana.GetBlockRequest) (*solana.GetBlockReply, error) {
	//callMethod := consts.SolanaMethodGetBalance
	reply := &solana.GetBlockReply{}
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getBlock",
			"params": []interface{}{
				request.Slot,
				map[string]interface{}{
					"encoding":           request.Encoding,
					"transactionDetails": request.TransactionDetails,
					"rewards":            request.Rewards,
				},
			},
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	//err = json.Unmarshal([]byte(gjson.GetBytes(response.Bytes(), "result.context").String()), &reply.Context)
	//if err != nil {
	//	return nil, err
	//}
	return reply, nil
}
func (sc *SolanaClient) GetClusterNodes(ctx context.Context) ([]*solana.ClusterNodesItem, error) {
	//callMethod := consts.SolanaMethodGetBalance
	reply := make([]*solana.ClusterNodesItem, 0)
	response, err := sc.HttpClient.
		R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"jsonrpc": "2.0", "id": 1,
			"method": "getClusterNodes",
		}).
		Post("")
	if err != nil {
		return nil, err
	}
	if gjson.GetBytes(response.Bytes(), "error").String() != "" {
		return nil, errors.New(gjson.GetBytes(response.Bytes(), "error.message").String())
	}
	gjson.GetBytes(response.Bytes(), "result").ForEach(func(key, value gjson.Result) bool {
		fmt.Println(key, value.String())
		item := &solana.ClusterNodesItem{}
		if loopErr := json.Unmarshal([]byte(value.String()), item); loopErr == nil {
			reply = append(reply, item)
		}
		return true
	})
	return reply, nil
}
