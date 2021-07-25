package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"web3.go/entity/resp"
)

type HTTPProvider struct {
	address string
	timeout int32
	client  *http.Client
}

func NewHTTPProvider(address string, timeout int32) *HTTPProvider {
	return NewHTTPProviderWithClient(address, timeout, &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	})
}

func NewHTTPProviderWithClient(address string, timeout int32, client *http.Client) *HTTPProvider {
	provider := new(HTTPProvider)
	provider.address = address
	provider.timeout = timeout
	provider.client = client

	return provider
}

func (provider HTTPProvider) SendRequest(method string, params interface{}) ([]byte, error) {
	var (
		reqBodyByte = make([]byte, 0)
		err         error
	)

	reqBodyByte, err = json.Marshal(JsonRPCReq{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      rand.Int(),
	})
	fmt.Println("send body: ", string(reqBodyByte))

	req, err := http.NewRequest("POST", provider.address, strings.NewReader(string(reqBodyByte)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpResp, err := provider.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var bodyBytes []byte

	if httpResp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}
	}
	err = provider.callbackCheck(bodyBytes)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil

}

func (provider HTTPProvider) callbackCheck(bodyBytes []byte) error {
	var commonResp = new(resp.JsonRpcCommonResp)
	err := json.Unmarshal(bodyBytes, commonResp)
	if err != nil {
		return err
	}
	if commonResp.Error.Code != 0 {
		return errors.New(commonResp.Error.Message)
	}
	return nil
}
