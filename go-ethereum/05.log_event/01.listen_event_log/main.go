package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 测试使用 BNB 的测试网络，节点使用 nodereal 免费的服务
var (
	//baseUrl = "https://bsc-testnet.nodereal.io/v1"
	baseUrl = "wss://bsc-testnet.nodereal.io/ws/v1"
	apiKey  = "307ba8bef61741ea996267d48bdfcf62"
)

func main() {
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial(fmt.Sprintf("%s/%s", baseUrl, apiKey))
	if err != nil {
		log.Fatal(err)
	}

	// 创建需要监听的筛选查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			// 合约地址(投票的合约，用于测试，可投票到 2099 年)
			common.HexToAddress("0x5102150D38D1c3443DD3522B344A39989282cad7"),
			// 测试网狗的合约
			common.HexToAddress("0xcc0ec652a675fd8d41e4bf28953ae3f6490f1bd8"),
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs) // 订阅
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
