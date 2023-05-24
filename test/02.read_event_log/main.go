package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://scroll-alphanet.blastapi.io/b9810890-43b0-41f4-b9ce-72f9d0af8699")
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个过滤查询
	// 指定我们想过滤的区块范围并指定从中读取此日志的合约地址
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2043125),
		ToBlock:   big.NewInt(2043129),
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		// 读取日志结构体中的附件信息（区块摘要，区块号和交易摘要）
		fmt.Println(vLog.BlockHash.Hex()) // 0x98576e84c24ba628a41f7f7eb6f24c1fbb74a527995de981a9f54b21cb5ca04e
		fmt.Println(vLog.BlockNumber)     // 32085598
		fmt.Println(vLog.TxHash.Hex())    // 0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b

		fmt.Println(vLog.Topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}
}
