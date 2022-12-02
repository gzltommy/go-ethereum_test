package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"log"
)

/*
 * 用于生成签名的组件是：签名者私钥，以及将要签名的数据的哈希。
 * 只要输出为 32 字节，就可以使用任何哈希算法。
 * 我们将使用 Keccak-256 作为哈希算法，这是以太坊常常使用的算法。
 */

func main() {
	// 1.加载私钥
	privateKey, err := crypto.HexToECDSA("debe1d34d97ad7ad00d4eaff4c291c9261a11c7f9432b867d7a987f698bbc33b")
	if err != nil {
		log.Fatal(err)
	}
	//StructSignExample1(privateKey)
	StructSignExample2(privateKey)
	//ArraySignExample(privateKey)

}

func StructSignExample1(privateKey *ecdsa.PrivateKey) {
	// 请求参数：
	typedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Claim": {
				{Name: "pid", Type: "uint256"},
				{Name: "claimNFT", Type: "address"},
				{Name: "dummyId", Type: "uint256"},
				{Name: "level", Type: "uint256"},
				{Name: "mintTo", Type: "address"},
				{Name: "royaltyReceiver", Type: "address"},
				{Name: "royaltyFeeNumerator", Type: "uint96"},
				{Name: "canTranster", Type: "bool"},
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
			"pid":                 math.NewHexOrDecimal256(7904),
			"claimNFT":            "0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826",
			"dummyId":             math.NewHexOrDecimal256(483759),
			"level":               math.NewHexOrDecimal256(1),
			"mintTo":              "0x9305efA882316e104BF74CfF9685b7e593b316ae",
			"royaltyReceiver":     "0x9305efA882316e104BF74CfF9685b7e593b316ae",
			"royaltyFeeNumerator": math.NewHexOrDecimal256(200),
			"canTranster":         true,
		},
		PrimaryType: "Claim",
	}

	signature, err := SignWithEip721(privateKey, typedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
}
func StructSignExample2(privateKey *ecdsa.PrivateKey) {
	// 请求参数：
	typedData := &apitypes.TypedData{
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

	signature, err := SignWithEip721(privateKey, typedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
}

// TODO:无法运行，截止到 2022-12-2 日网络上并没有解决方案
func ArraySignExample(privateKey *ecdsa.PrivateKey) {
	// 请求参数：
	typedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Claim": {
				{Name: "users", Type: "address[]"},
				{Name: "amounts", Type: "uint256[]"},
				{Name: "user", Type: "address"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
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
			"users":    []string{"0x111da67948Ef5Ed1f82D707B8cd7e3B1DFa87AEa"},
			"amounts":  []string{"1"},
			"user":     "0x9305efA882316e104BF74CfF9685b7e593b316ae",
			"nonce":    math.NewHexOrDecimal256(1),
			"deadline": math.NewHexOrDecimal256(100),
		},
		PrimaryType: "Claim",
	}

	signature, err := SignWithEip721(privateKey, typedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
}

func SignWithEip721(privateKey *ecdsa.PrivateKey, typedData *apitypes.TypedData) (string, error) {
	if privateKey == nil || typedData == nil {
		return "", errors.New("invalid parameter")
	}

	// 1、获取需要签名的数据的 Keccak-256 的哈希
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return "", err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return "", err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sigHash := crypto.Keccak256(rawData)

	// 2、使用私钥签名哈希，得到签名
	signature, err := crypto.Sign(sigHash, privateKey)
	if err != nil {
		return "", err
	}
	if signature[64] < 27 {
		signature[64] += 27
	}

	// 3、转换为 16 进制的字符串
	return hexutil.Encode(signature), nil
}
