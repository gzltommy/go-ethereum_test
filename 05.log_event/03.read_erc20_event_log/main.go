package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	// 注意国内要设置代理才能连接
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	//client, err := ethclient.Dial("wss://data-seed-prebsc-1-s1.binance.org:8546")
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	contractAddress := common.HexToAddress("0x195c9831a0dac5f33a46f15b5d2fc3c80bf91f3a")

	// 构造一个过滤查询
	blockHash := common.HexToHash("0x73d7f99b23ea8d3a2ccc5394edca0075260c954bc4f15de61e418553197be26d")
	query := ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//// 为了解码日志，我们需要导入我们智能合约的 ABI。
	//contractAbi, err := abi.JSON(strings.NewReader(store.StoreMetaData.ABI)) // 该函数返回一个我们可以在 Go 应用程序中使用的解析过的 ABI 接口
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(len(logs))

	for _, vLog := range logs {
		//// 读取日志结构体中的附件信息（区块摘要，区块号和交易摘要）
		fmt.Println("BlockHash:", vLog.BlockHash.Hex()) // 0x98576e84c24ba628a41f7f7eb6f24c1fbb74a527995de981a9f54b21cb5ca04e
		fmt.Println("BlockNumber:", vLog.BlockNumber)   // 32085598
		fmt.Println("TxHash:", vLog.TxHash.Hex())       // 0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b

		// 读取日志结构体中的 topics
		// 若您的 solidity 事件包含 indexed 事件类型，那么它们将成为主题而不是日志的数据属性的一部分。
		// 在 solidity 中您最多只能有 4 个主题，但只有 3 个可索引的事件类型。第一个主题总是事件的签名。
		// 我们的示例合约不包含可索引的事件
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

		//tlog, err := logsTransfer(vLog)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
}

func logsTransfer(vLog types.Log) (TransferLog, error) {
	var logData TransferLog
	topicsLen := len(vLog.Topics)
	var tokenId *big.Int
	var from, to common.Address
	if topicsLen < 3 {
		return logData, errors.New(fmt.Sprintf("error %d < 3", topicsLen))
	} else {
		from = common.HexToAddress(vLog.Topics[1].Hex())
		to = common.HexToAddress(vLog.Topics[2].Hex())
		tokenId = vLog.Topics[3].Big()
	}
	logData.From = strings.ToLower(from.Hex())
	logData.To = strings.ToLower(to.Hex())
	logData.Ids = append(logData.Ids, tokenId)
	logData.Amounts = append(logData.Amounts, big.NewInt(1))
	return logData, nil
}

type TransferLog struct {
	From    string
	To      string
	Ids     []*big.Int
	Amounts []*big.Int
}
