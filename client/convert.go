package client

import (
	"fmt"
	"github.com/6boris/web3-go/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

type GinGinMethodConvert struct {
	pool *Pool
}

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func NewGinMethodConvert(conf *ConfPool) *GinGinMethodConvert {
	g := &GinGinMethodConvert{}
	g.pool = NewPool(conf)
	return g
}

func (g *GinGinMethodConvert) ConvertGinHandler(ctx *gin.Context) {
	req := &EthCallProxyRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(400, ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: err.Error()})
		return
	}
	req.ChainID, err = strconv.ParseInt(ctx.Param("chain_id"), 10, 64)
	if err != nil {
		ctx.JSON(400, ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: err.Error()})
		return
	}
	client := g.pool.GetClient(req.ChainID)
	if client == nil {
		ctx.JSON(400, ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: "Not have enable client"})
		return
	}

	switch req.Method {
	case consts.AbiMethodEthChainID:
		g._ethChainID(ctx, req, client)
	case consts.AbiMethodClientVersion:
		g._web3ClientVersion(ctx, req, client)
	case consts.AbiMethodEthGasPrice:
		g._ethGasPrice(ctx, req, client)
	case consts.AbiMethodEthBlockNumber:
		g._ethBlockNumber(ctx, req, client)
	case consts.AbiMethodEthGetBalance:
		g._ethGetBalance(ctx, req, client)
	default:
		ctx.JSON(400, ErrReply{Code: 400, Reason: "PARAMS_ERR", Message: fmt.Sprintf("Not support this method:%s", req.Method)})
		return
	}

}

// eth_chainId https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_chainId
func (g *GinGinMethodConvert) _ethChainID(ctx *gin.Context, req *EthCallProxyRequest, client *Client) {
	resp := &EthCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.ChainID(ctx)
	if err != nil {
		ctx.JSON(500, ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", ethResp.Text(16))
	ctx.JSON(http.StatusOK, resp)
}

// web3_clientVersion https://ethereum.org/en/developers/docs/apis/json-rpc/#web3_clientversion
func (g *GinGinMethodConvert) _web3ClientVersion(ctx *gin.Context, req *EthCallProxyRequest, client *Client) {
	resp := &EthCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.ClientVersion(ctx)
	if err != nil {
		ctx.JSON(500, ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = ethResp
	ctx.JSON(http.StatusOK, resp)
}

// eth_gasPrice https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_gasprice
func (g *GinGinMethodConvert) _ethGasPrice(ctx *gin.Context, req *EthCallProxyRequest, client *Client) {
	resp := &EthCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.SuggestGasPrice(ctx)
	if err != nil {
		ctx.JSON(500, ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", ethResp.Text(16))
	ctx.JSON(http.StatusOK, resp)
}

// eth_blockNumber https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_blocknumber
func (g *GinGinMethodConvert) _ethBlockNumber(ctx *gin.Context, req *EthCallProxyRequest, client *Client) {
	resp := &EthCallProxyReply{
		ID:      req.ID,
		JsonRpc: req.JsonRpc,
	}
	ethResp, err := client.BlockNumber(ctx)
	if err != nil {
		ctx.JSON(500, ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", big.NewInt(int64(ethResp)).Text(16))
	ctx.JSON(http.StatusOK, resp)
}

// eth_getBalance https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_getbalance
func (g *GinGinMethodConvert) _ethGetBalance(ctx *gin.Context, req *EthCallProxyRequest, client *Client) {
	resp := &EthCallProxyReply{
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
	fmt.Println(ethResp, err)
	if err != nil {
		ctx.JSON(500, ErrReply{Code: 500, Reason: "ETH_ERR", Message: err.Error(), Metadata: map[string]string{"transport_url": client.TransportURL}})
		return
	}
	resp.Result = fmt.Sprintf("0x%s", ethResp.Text(16))
	ctx.JSON(http.StatusOK, resp)
}
