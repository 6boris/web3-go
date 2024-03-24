package pk

import (
	"crypto/ecdsa"
	"errors"

	clientModel "github.com/6boris/web3-go/model/client"
	"github.com/ethereum/go-ethereum/crypto"
)

func TransformPkToEvmSigner(privateKey string) (*clientModel.ConfEvmChainSigner, error) {
	signer := &clientModel.ConfEvmChainSigner{}
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("TransformPkToEvmSigner Failed")
	}
	signer.PrivateKey = pk
	signer.PublicAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return signer, err
}
