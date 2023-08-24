package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://opbnb-testnet-rpc.bnbchain.org/")
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个过滤查询
	// 指定我们想过滤的区块范围并指定从中读取此日志的合约地址
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(3551791),
		ToBlock:   big.NewInt(5913531),
		Addresses: []common.Address{
			common.HexToAddress("0xcf8d609948d9df91f8719018887a67b4e6360dc8"), // medal_mint
		},
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x4acbe73b979f0a6452384be471194eef86ad06759695d9d54c94191fb5f29e8c"),
			},
		},
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	bmap := make(map[uint64]struct{}, 100000)
	for _, eLog := range logs {
		if _, ok := bmap[eLog.BlockNumber]; !ok {
			bmap[eLog.BlockNumber] = struct{}{}
		}
	}
	s := make([]uint64, 0, len(bmap))
	for k, _ := range bmap {
		s = append(s, k)
	}
	buf, _ := json.Marshal(s)
	fmt.Println(string(buf))
}
