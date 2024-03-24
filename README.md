<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/6boris/web3-go" target="_blank">
    <img src="./assets/img/Ethereum-icon-purple.svg" alt="Logo" width="680" height="256">
  </a>

  <h3 align="center">Web3 Go</h3>

  <p align="center">
    Ethereum Dapp Go API, inspired by 
    <a href="https://github.com/ChainSafe/web3.js" target="_blank">web3.js</a>.
    <br />
    <a href="https://github.com/6boris/web3-go/issues" target="_blank">Report Bug</a>
    ·
    <a href="https://github.com/6boris/web3-go/pulls" target="_blank">Pull Request</a>
  </p>
</p>

[![WEBSITE](https://img.shields.io/badge/Web3-Go-brightgreen)](https://github.com/kylesliu/web3-go)
[![LISTENSE](https://img.shields.io/github/license/6boris/web3-go)](https://github.com/kylesliu/web3-go/blob/main/LICENSE)

## Introduction

This is the Ethereum [Golang API](https://github.com/kylesliu/web3-go) which connects to the Generic [JSON-RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) spec.

You need to run a local or remote Ethereum node to use this library.

Here is an open source case [Web3 Studio](https://web3-studio.leek.dev/d/demo/web3-studio) reference under.

<a href="https://web3-studio.leek.dev/d/demo/web3-studio" target="_blank">
  <img src="https://s.gin.sh/develop/web3/web3-studio-demo.png" alt="Logo">
</a>


### Client


```bash
export WEB3_GO_DEV_KEY_1="YOU_EVM_PRIVATE_KEY 1"
export WEB3_GO_DEV_KEY_2="YOU_EVM_PRIVATE_KEY 1"
```

```bash
go get github.com/6boris/web3-go
```


```go
package main

import (
  "context"
  "fmt"
  "github.com/6boris/web3-go/client"
  clientModel "github.com/6boris/web3-go/model/client"
  "github.com/6boris/web3-go/pkg/pk"
  "github.com/ethereum/go-ethereum/common"
  "github.com/shopspring/decimal"
  "math/big"
  "os"
)

func main() {
  ctx := context.Background()
  evmSigners := make([]*clientModel.ConfEvmChainSigner, 0)
  for _, v := range []string{"WEB3_GO_DEV_KEY_1", "WEB3_GO_DEV_KEY_1"} {
    signer, loopErr := pk.TransformPkToEvmSigner(os.Getenv(v))
    if loopErr != nil {
      continue
    }
    evmSigners = append(evmSigners, signer)
  }
  ec, err := client.NewEvmClient(&clientModel.ConfEvmChainClient{
    TransportURL: "https://1rpc.io/matic",
    GasFeeRate:   decimal.NewFromFloat(2),
    GasLimitRate: decimal.NewFromFloat(1.5),
    Signers:      evmSigners,
  })

  nativeTx, err := ec.SendTransactionSimple(
    ctx,
    ec.GetAllSinners()[0], ec.GetAllSinners()[1],
    big.NewInt(1),
  )
  if err != nil {
    panic(err)
  }
  fmt.Println(fmt.Sprintf("Native Token Tx: %s", nativeTx.Hash()))

  erc20Tx, err := ec.ERC20Transfer(
    ctx,
    common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
    ec.GetAllSinners()[0], ec.GetAllSinners()[1],
    decimal.NewFromFloat(0.000001).Mul(decimal.New(1, int32(6))).BigInt(),
  )
  if err != nil {
    panic(err)
  }
  fmt.Println(fmt.Sprintf("USDT Token Tx: %s", erc20Tx.Hash()))
}
/*
Output:
    Native Token Tx: 0xf3aa0e634357a222c39e8.......8f1a7e7d313db71827c3
    USDT Token Tx: 0xedd54d9e6bd3738880cd55........21aa69f06dba1e011625
*/
```

## Development Trips
- [X] Client
  - [ ] Base Method
    - [X] eth_chainId
    - [X] web3_clientVersion
    - [X] eth_gasPrice
    - [X] eth_blockNumber
    - [X] eth_getBalance
    - [ ] ...
  - [ ] Middleware
    - [X] LoadBalance
    - [X] Metrics
    - [ ] Grafana
    - [ ] CircuitBreaker
  - [ ] Business Cases
    - [ ] Web3 Studio
- [ ] Other ...



## Community

- [web3.js](https://github.com/ChainSafe/web3.js) Ethereum JavaScript API.
- [Web3j](https://github.com/web3j/web3j) Web3 Java Ethereum Ðapp API.
- [Web3.py](https://github.com/ethereum/web3.py) A Python library for interacting with Ethereum.

## Provider
- https://public.blockpi.io/


## Dev tool

- [JSON RPC](https://www.jsonrpc.org/specification) Defining the JSON RPC specification.
- [Go Ethereum](https://github.com/ethereum/go-ethereum) Official Golang implementation of the Ethereum protocol.
- [Ethereum 1.0 API](https://github.com/ethereum/eth1.0-apis) Ethereum JSON-RPC Specification.
