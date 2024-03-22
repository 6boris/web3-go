package client

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/6boris/web3-go/model/solana"

	clientModel "github.com/6boris/web3-go/model/client"

	"github.com/6boris/web3-go/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type GinMethodConvert struct {
	pool *Pool
}

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func NewGinMethodConvert(conf *clientModel.ConfPool) *GinMethodConvert {
	g := &GinMethodConvert{}
	g.pool = NewPool(conf)
	return g
}

func (g *GinMethodConvert) _convertGinHandler(ctx *gin.Context) {
	req := &clientModel.EvmCallProxyRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: err.Error()})
		return
	}
	evmClient := g.pool.GetEvmClient(req.ChainID)
	if req.ChainID > 0 && evmClient == nil {
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: "Not have enable client"})
		return
	}
	switch req.Method {
	case consts.EvmMethodChainID:
		g._convertEvmChainID(ctx, req, evmClient)
	case consts.EvmMethodSuggestGasTipCap:
		g._convertEvmGasPrice(ctx, req, evmClient)
	case consts.EvmMethodBlockNumber:
		g._convertEvmBlockNumber(ctx, req, evmClient)
	case consts.EvmMethodBalanceAt:
		g._convertEvmGetBalance(ctx, req, evmClient)
	default:
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: fmt.Sprintf("Not support this method:%s", req.Method)})
		return
	}
}

// eth_chainId https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_chainId
func (g *GinMethodConvert) _convertEvmChainID(ctx *gin.Context, req *clientModel.EvmCallProxyRequest, client *EvmClient) {
	resp := &clientModel.EvmCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.ChainID(ctx)
	if err != nil {
		ctx.JSON(500, &clientModel.ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", ethResp.Text(16))
	ctx.JSON(http.StatusOK, resp)
}

// eth_gasPrice https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_gasprice
func (g *GinMethodConvert) _convertEvmGasPrice(ctx *gin.Context, req *clientModel.EvmCallProxyRequest, client *EvmClient) {
	resp := &clientModel.EvmCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.SuggestGasPrice(ctx)
	if err != nil {
		ctx.JSON(500, &clientModel.ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}

	resp.Result = ethResp
	ctx.JSON(http.StatusOK, resp)
}

// eth_blockNumber https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_blocknumber
func (g *GinMethodConvert) _convertEvmBlockNumber(ctx *gin.Context, req *clientModel.EvmCallProxyRequest, client *EvmClient) {
	resp := &clientModel.EvmCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.BlockNumber(ctx)
	if err != nil {
		ctx.JSON(500, &clientModel.ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", big.NewInt(int64(ethResp)).Text(16))
	ctx.JSON(http.StatusOK, resp)
}

// eth_getBalance https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_getbalance
func (g *GinMethodConvert) _convertEvmGetBalance(ctx *gin.Context, req *clientModel.EvmCallProxyRequest, client *EvmClient) {
	resp := &clientModel.EvmCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}

	account := req.Params[0].(string)
	var blockNumber *big.Int
	if len(req.Params) > 2 {
		blockStr, ok := req.Params[1].(string)
		if ok {
			block, err := strconv.ParseInt(blockStr, 10, 64)
			if err == nil {
				blockNumber = big.NewInt(block)
			}
		}
	}
	ethResp, err := client.BalanceAt(ctx, common.HexToAddress(account), blockNumber)
	if err != nil {
		ctx.JSON(500, &clientModel.ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", ethResp.Text(16))
	ctx.JSON(http.StatusOK, resp)
}

func (g *GinMethodConvert) _convertSolanaHandler(ctx *gin.Context) {
	req := &clientModel.SolanaCallProxyRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: err.Error()})
		return
	}
	solanaClient := g.pool.GetSolanaClient(req.ChainEnv)
	if solanaClient == nil {
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: "Not have enable client"})
		return
	}
	switch req.Method {
	case consts.SolanaMethodGetBalance:
		g._convertSolanaGetBalance(ctx, req, solanaClient)
	default:
		ctx.JSON(400, clientModel.ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: fmt.Sprintf("Not support this method:%s", req.Method)})
		return
	}
}

func (g *GinMethodConvert) _convertSolanaGetBalance(ctx *gin.Context, req *clientModel.SolanaCallProxyRequest, client *SolanaClient) {
	resp := &clientModel.EvmCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}

	account := req.Params[0].(string)
	callResp, err := client.GetBalance(ctx, &solana.GetBalanceRequest{
		Account: account,
	})
	if err != nil {
		ctx.JSON(500, &clientModel.ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = callResp.Value
	ctx.JSON(http.StatusOK, resp)
}
