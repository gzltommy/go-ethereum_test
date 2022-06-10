package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum_test/04.smart_contract/02.deploy_contract/contract/store" // 导入包
	"log"
)

var (
	netWork   = "kovan"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("https://%s.infura.io/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	contractAddress := common.HexToAddress("0x3f6c4f9FDd3AAdfE545D45221477A1F553404Ae5")

	// 调用“NewXXX”方法，加载合约（需要合约地址）
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	// 使用合约实例提供的方法来阅读智能合约。

	// 合约中有一个名为 version 的全局变量。 因为它是公开的，这意味着它们将成为我们自动创建的 getter 函数。
	// 常量和 view 函数也接受 bind.CallOpts 作为第一个参数。一般情况下我们可以用 nil
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"

	// 下面的代码，在下节中将数据写入后便可读取到
	//key := [32]byte{}
	//copy(key[:], []byte("tommy"))
	//result, err := instance.Items(nil, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(result[:])) // "汤米"
}
