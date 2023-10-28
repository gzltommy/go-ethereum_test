package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

var cs = []string{
	"0x24495c7A24B41686215558bbAB0bE2172bf0Dc1e",
	"0xd2DF3189E276CAdEC88c1928f99f30E01B24935f",
	"0x47F8131e6b6041eE87C9C38aE8C804AF6e0B3eF9",

	//"0xAC1f9Fadc33cC0799Cf7e3051E5f6b28C98966EE",//Factory
	"0x2aDd449582863e75B81338244546B886724AB8b9",
	"0x33bE2fb2F6D0476105ec155D0A2319E8074B54E9",
	"0x7D27D520ACFc0C5EEa1e25277C3eA0b284f75937",
	"0xf05405f0f10cEA70Ae2aa7BD799f2eB59a3b871e",
	"0x481a22A95ACb664A574dBc959a1D6aEc7E245Cdd",
	"0x23acd4c7c736bb97438475f3d5343a5097581129",
	// todo:外加 imo 中的合约地址
}

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://arbitrum.llamarpc.com")
	if err != nil {
		log.Fatal(err)
	}

	contract := make([]common.Address, 0, len(cs))
	for _, v := range cs {
		contract = append(contract, common.HexToAddress(v))
	}

	// 构造一个过滤查询
	// 指定我们想过滤的区块范围并指定从中读取此日志的合约地址
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(118975472),
		ToBlock:   big.NewInt(118976472),
		Addresses: contract,
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		transaction, isPending, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			fmt.Printf("TransactionReceipt fail.err:%v", err)
			return
		}
		if isPending {
			return
		}

		header, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Printf("HeaderByNumber fail.err:%v", err)
			return
		}

		fmt.Println(header.Number.String())
		fmt.Println(time.Unix(int64(header.Time), 0))

		msg, err := transaction.AsMessage(types.LatestSignerForChainID(big.NewInt(int64(42161))), nil)
		if err != nil {
			fmt.Printf("AsMessage fail.err:%v", err)
			return
		}
		_ = header
		fmt.Println(msg.From().Hex())
		fmt.Println(msg.To().Hex())
	}
}
