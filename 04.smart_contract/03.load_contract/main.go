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
	_ = instance
}
