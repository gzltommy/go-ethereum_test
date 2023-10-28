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
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	if err != nil {
		log.Fatal(err)
	}

	var (
		fromBlock = int64(32871288)
		endBlock  = int64(32961288)
		bmap      = make(map[uint64]struct{}, 100000)
		s         = make([]uint64, 0, len(bmap))
	)

	for {
		toBlock := fromBlock + 50000
		if toBlock > endBlock {
			toBlock = endBlock
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock), // 开始
			ToBlock:   big.NewInt(toBlock),   // 结束
			Addresses: []common.Address{
				common.HexToAddress("0xCE0e4e4D2Dc0033cE2dbc35855251F4F3D086D0A"), // medal_mint
			},
			Topics: [][]common.Hash{
				{
					common.HexToHash("0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8"),
				},
			},
		}

		// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		for _, eLog := range logs {
			if _, ok := bmap[eLog.BlockNumber]; !ok {
				bmap[eLog.BlockNumber] = struct{}{}
				s = append(s, eLog.BlockNumber)
			}
		}
		if toBlock >= endBlock {
			break
		}
	}

	buf, _ := json.Marshal(s)
	fmt.Println(string(buf))
}
