package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 正如您所见，首个主题只是被哈希过的事件签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
