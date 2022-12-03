package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//netWork   = "kovan"
	netWork   = "mainnet"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	// 1. 设置订阅以便在新区块被开采时获取事件。首先，我们需要一个支持 websocket RPC 的以太坊服务提供者。在示例中，我们将使用 infura 的 websocket 端点。
	client, err := ethclient.Dial(fmt.Sprintf("wss://%s.infura.io/ws/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal("-1--", err)
	}

	// 2.接下来，我们将创建一个新的通道，用于接收最新的区块头。
	headers := make(chan *types.Header)

	// 3.现在我们调用客户端的 SubscribeNewHead 方法，它接收我们刚创建的区块头通道，该方法将返回一个订阅对象。
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("-2--", err)
	}

	// 订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal("-3--", err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			// 要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的 BlockByHash 函数。
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal("-4--", err)
			}

			// 现在，您可以读取整个区块的元数据字段，交易列表等等。
			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
