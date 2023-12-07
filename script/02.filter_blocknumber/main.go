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
	"os"
)

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://test-rpc.combonetwork.io")
	if err != nil {
		log.Fatal(err)
	}

	scanBlockFile, err := os.Create("need_scan_block.json")
	if err != nil {
		panic(err)
	}
	defer scanBlockFile.Close()

	var (
		fromBlock = int64(19750676)
		endBlock  = int64(19959476)
		bmap      = make(map[uint64]struct{}, 100000)
		s         = make([]uint64, 0, len(bmap))
		step      = int64(2000)
		toBlock   = fromBlock + step
	)

	for {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock), // 开始
			ToBlock:   big.NewInt(toBlock),   // 结束
			Addresses: []common.Address{
				common.HexToAddress("0xAC1f9Fadc33cC0799Cf7e3051E5f6b28C98966EE"), // medal_mint
			},
			Topics: [][]common.Hash{
				{
					common.HexToHash("0x2a0cd797412ef7ea19c365def73dc12d3ec05fbe2c714249219e6b2739cce14a"),
				},
			},
		}

		// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
			break
		}
		for _, eLog := range logs {
			if _, ok := bmap[eLog.BlockNumber]; !ok {
				bmap[eLog.BlockNumber] = struct{}{}
				s = append(s, eLog.BlockNumber)
			}
		}

		fromBlock = toBlock + 1
		toBlock = toBlock + step
		if toBlock > endBlock {
			toBlock = endBlock
		}
		if fromBlock > toBlock {
			break
		}
	}

	buf, _ := json.Marshal(s)
	n, err := scanBlockFile.Write(buf)
	fmt.Println("结束：", n, err)
}
