package client

import (
	"context"
	"errors"
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
	var err error
	client := &SolanaClient{
		ClientID:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		Provider:     conf.Provider,
		TransportURL: conf.TransportURL,
		ChainEnv:     conf.ChainEnv,
	}
	client.HttpClient = req.C().
		//DevMode().
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
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (sc *SolanaClient) GetBalance(ctx context.Context, request *solana.GetBalanceRequest) (*solana.GetBalanceReply, error) {
	//callMethod := consts.SolanaMethodGetBalance
	reply := &solana.GetBalanceReply{}
	response, err := sc.HttpClient.
		//DevMode().
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
	reply.Context.Slot, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.context.slot").String())
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
		DevMode().
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
	reply.Context.Slot, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.context.slot").String())
	if err != nil {
		return nil, err
	}
	reply.Context.Slot = reply.Context.Slot.Div(decimal.New(1, 9))

	reply.Value, err = decimal.NewFromString(gjson.GetBytes(response.Bytes(), "result.value.uiAmountString").String())
	if err != nil {
		return nil, err
	}
	return reply, nil
}
