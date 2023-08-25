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
	client, err := ethclient.Dial("https://arb1.arbitrum.io/rpc")
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个过滤查询
	// 指定我们想过滤的区块范围并指定从中读取此日志的合约地址
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(105033747),
		ToBlock:   big.NewInt(124440584),
		Addresses: []common.Address{
			common.HexToAddress("0x481a22a95acb664a574dbc959a1d6aec7e245cdd"), // medal_mint
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
	s := make([]uint64, 0, len(bmap))
	for _, eLog := range logs {
		if _, ok := bmap[eLog.BlockNumber]; !ok {
			bmap[eLog.BlockNumber] = struct{}{}
			s = append(s, eLog.BlockNumber)
		}
	}

	buf, _ := json.Marshal(s)
	fmt.Println(string(buf))
}
