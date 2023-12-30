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
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	//client, err := ethclient.Dial("https://rpc.combonetwork.io")
	if err != nil {
		log.Fatal(err)
	}

	scanBlockFile, err := os.Create("need_scan_block.json")
	if err != nil {
		panic(err)
	}
	defer scanBlockFile.Close()

	var (
		//fromBlock = int64(2658317)
		//endBlock  = int64(2913789)

		fromBlock = int64(33074075)
		endBlock  = int64(34774075)
		bmap      = make(map[uint64]struct{}, 100000)
		s         = make([]uint64, 0, len(bmap))
		step      = int64(20000)
		toBlock   = fromBlock + step
	)

	for {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock), // 开始
			ToBlock:   big.NewInt(toBlock),   // 结束
			Addresses: []common.Address{
				common.HexToAddress("0x04d155b0ce5eca2edadff3f49cd9f5f0f37432e8"),
				common.HexToAddress("0x17316bb0b16a57177272dd88617841037bcec3d3"),
				common.HexToAddress("0x2add449582863e75b81338244546b886724ab8b9"),
				common.HexToAddress("0x33be2fb2f6d0476105ec155d0a2319e8074b54e9"),
				common.HexToAddress("0x423a88e34adf421b388e4fd1cc25f21a8076e918"),
				common.HexToAddress("0x4e399f27537bc480bdd66f304f117655b4bc55ff"),
				common.HexToAddress("0x86677398b2822b636848c1f196dc3885145d979c"),
				common.HexToAddress("0x91bec3b4d64db7cb5036f8574a23676c487d1467"),
				common.HexToAddress("0x9ad0e3224df03708bf82ba40092315951358e579"),
				common.HexToAddress("0xac1f9fadc33cc0799cf7e3051e5f6b28c98966ee"),
				common.HexToAddress("0xcd824bbcc9c5a308f74d46892c7376dc5aa968c7"),
				common.HexToAddress("0xe085e9462ae0111886c1f9cd11150fe0c71fa634"),
				common.HexToAddress("0xf41622fdcfaa8965110c9fd14291ba08b7069d28"),
			},
			Topics: [][]common.Hash{
				{
					common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
					common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
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
