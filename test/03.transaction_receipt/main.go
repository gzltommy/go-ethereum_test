package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0xe68204598c692c19e13a60eaf654643a03d4689493797a19a6ecfbedba1eeb4f")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status) // 1

	// 才能够 receipt 中获取 log
	logs, _ := json.Marshal(receipt.Logs)
	fmt.Println("--------", string(logs))
}
