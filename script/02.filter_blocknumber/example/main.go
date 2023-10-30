package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum_test/test/test2/NFT721Bridge"
	"log"
	"math/big"
	"strings"
)

var (
	nftBridgeReceiveNftContract = []string{
		"0x844F5756e4a588532bff5D22443284e694b46ff8",
		"0x49586e8e216671308cC3F04AFBBd46394E28B6Ce",
		"0x81efE18B9aa586CC9678DCb80c62E06fDb321D76",
		"0x7cFf558ebe0eA7335edBA8f60aA2Ee9762d1F233",
		"0xb82E3E96fF8EfEFd5369586b62b207b459cE97c4",
		"0x518F3bb208DE96e9Aa6388ca8621981BB6d0db9b",
		"0x4df882D8ad4003C94C912d29b3F84101ce095c81",
		"0x5e26e82319041FB505878a11EB6eE17E02dAe379",
		"0x6457A211c447FEcb37c08127671CdCD2E75E28f9",
		"0xfd8eCD74f523517863f186d60e38fF36548d3f3A",
		"0x8Cc033FECC160eAe95c084B1783EB6528168e75C",
		"0xC2DE540c008D2f8B0107383e15bE075f0CA1Cb06",
		"0x5eD79df8D6876482ca05C5cb07d3C6DB7c99744D",
		"0x2029ec97f205DCbd7BcC4976f9f573D272aef036",
		"0xDeaFE008fCF700D5DcCbc01E60Dc3aD81A7A4A0C",
		"0x01F49CE796436bEf82Fa9F7B515Ff5576BBe9f51",
		"0xc675bA2ce965dcE39662DdA741a394f6FA12190c",
	}
)

func main() {
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	if err != nil {
		log.Fatal(err)
	}

	var (
		fromBlock = int64(32851288)
		endBlock  = int64(32851388)
		//endBlock = int64(33043388)
		bmap    = make(map[uint64]struct{}, 100000)
		s       = make([]uint64, 0, len(bmap))
		step    = int64(4000)
		toBlock = fromBlock + step
	)

	// 如果只是用于解析日志，client 参数可以不传（nil）
	ins, err := NFT721Bridge.NewNFT721BridgeFilterer(common.HexToAddress("0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8"), nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock), // 开始
			ToBlock:   big.NewInt(toBlock),   // 结束
			Addresses: []common.Address{
				common.HexToAddress("0xCE0e4e4D2Dc0033cE2dbc35855251F4F3D086D0A"),
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
			return
		}

		for _, vLog := range logs {
			eventLog, err := ins.ParseReceiveNFT(vLog)
			if err != nil {
				log.Fatal(err)
				return
			}

			yes := false
			for _, v := range nftBridgeReceiveNftContract {
				if strings.EqualFold(v, eventLog.Token.Hex()) {
					yes = true
					break
				}
			}
			if !yes {
				continue
			}

			if _, ok := bmap[vLog.BlockNumber]; !ok {
				bmap[vLog.BlockNumber] = struct{}{}
				s = append(s, vLog.BlockNumber)
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
	fmt.Println(string(buf))
}
