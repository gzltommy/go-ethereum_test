package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum_test/04.smart_contract/02.deploy_contract/contract/store" // 导入包
	"log"
	"math/big"
	"strings"
)

var (
	netWork   = "kovan"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial(fmt.Sprintf("wss://%s.infura.io/ws/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	contractAddress := common.HexToAddress("0x3f6c4f9FDd3AAdfE545D45221477A1F553404Ae5")

	// 构造一个过滤查询
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(32085286),
		ToBlock:   big.NewInt(32085598),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 为了解码日志，我们需要导入我们智能合约的 ABI。
	contractAbi, err := abi.JSON(strings.NewReader(store.StoreMetaData.ABI)) // 该函数返回一个我们可以在 Go 应用程序中使用的解析过的 ABI 接口
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		// 读取日志结构体中的附件信息（区块摘要，区块号和交易摘要）
		fmt.Println(vLog.BlockHash.Hex()) // 0x98576e84c24ba628a41f7f7eb6f24c1fbb74a527995de981a9f54b21cb5ca04e
		fmt.Println(vLog.BlockNumber)     // 32085598
		fmt.Println(vLog.TxHash.Hex())    // 0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b

		// 读取日志结构体中的数据信息
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data) // 解码日志的数据字段
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // tommy
		fmt.Println(string(event.Value[:])) // 汤米

		// 读取日志结构体中的 topics
		// 若您的 solidity 事件包含 indexed 事件类型，那么它们将成为主题而不是日志的数据属性的一部分。
		// 在 solidity 中您最多只能有 4 个主题，但只有 3 个可索引的事件类型。第一个主题总是事件的签名。
		// 我们的示例合约不包含可索引的事件
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	// 正如您所见，首个主题只是被哈希过的事件签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
