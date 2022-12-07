package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

func Author(privateKey *ecdsa.PrivateKey, client *ethclient.Client, gasLimit int64, fixGasPrice int64) (auth *bind.TransactOpts, err error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	if err != nil {
		return nil, err
	}

	var gasPrice *big.Int
	if fixGasPrice > 0 {
		gasPrice = big.NewInt(fixGasPrice * 1000000000)
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	}

	if err != nil {
		return nil, nil
	}
	chainId, err := client.ChainID(context.Background())
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func SignWithEip721(privateKey *ecdsa.PrivateKey, typedData *apitypes.TypedData) ([]byte, error) {
	if privateKey == nil || typedData == nil {
		return nil, errors.New("invalid parameter")
	}

	// 1、获取需要签名的数据的 Keccak-256 的哈希
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sigHash := crypto.Keccak256(rawData)
	// 2、使用私钥签名哈希，得到签名
	signature, err := crypto.Sign(sigHash, privateKey)
	if err != nil {
		return nil, err
	}
	if signature[64] < 27 {
		signature[64] += 27
	}
	return signature, nil
}
