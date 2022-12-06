package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
	"log"
)

/*
 * 验证签名需要 3 个数据：签名者的公钥、签名，原始数据的哈希。
 */

func main() {
	// 1.签名
	privateKey, err := crypto.HexToECDSA("6cdd2cd6a8f813c33ac0c59a92798da9c0e490e17da96fbb9dc86fd1e77d23e2")
	if err != nil {
		log.Fatal(err)
	}

	// 请求参数：
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Person": {
				{Name: "name", Type: "string"},
				{Name: "wallets", Type: "address"},
			},
			"Mail": {
				{Name: "from", Type: "Person"},
				{Name: "to", Type: "Person"},
				{Name: "contents", Type: "string"},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              "SL",
			Version:           "1.0.0",
			ChainId:           math.NewHexOrDecimal256(1),
			VerifyingContract: "0x0fC5025C764cE34df352757e82f7B5c4Df39A836",
			Salt:              "",
		},
		Message: map[string]interface{}{
			"contents": "Hello, tommy!",
			"from": map[string]interface{}{
				"name":    "tommy",
				"wallets": "0x111da67948Ef5Ed1f82D707B8cd7e3B1DFa87AEa",
			},
			"to": map[string]interface{}{
				"name":    "zorro",
				"wallets": "0x211da67948Ef5Ed1f72D707B8cd7e3B1DFa87AEb",
			},
		},
		PrimaryType: "Mail",
	}

	signature, err := SignWithEip721(privateKey, &typedData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----签名-----", signature)

	// 2、校验签名
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("-----地址------", address)

	ok, err = VerifyEip721Signature(address, hexutil.Encode(signature), typedData)
	if err != nil {
		log.Fatal("VerifyEip721Signature fail.", err)
	}
	fmt.Println("-----校验签名----", ok)
}

func VerifyEip721Signature(address, signature string, typedData apitypes.TypedData) (bool, error) {
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		common.HexToAddress(address),
		typedData,
		signature,
	)
	return valid, err
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
