package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	// 2.获取需要签名的数据的 Keccak-256 的哈希
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	// 使用私钥签名哈希，得到签名
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x50075cc091032e1ea83a745b7aa7fe5e84ebc95d40f57b5b61556f193263809842b745d6d267a127505a127e7dd98ae93a21613179e54df7c4c15bd29de87aee00
}
