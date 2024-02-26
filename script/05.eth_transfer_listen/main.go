package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	util "go-ethereum_test/10.utils"
	"log"
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//client, err := ethclient.Dial("https://eth-mainnet.nodereal.io/v1/fa6e1351d6064b41b40cfd6c0500d0ec")
	client, err := ethclient.Dial("https://arb1.arbitrum.io/rpc")
	if err != nil {
		fmt.Println("-1-", err)
		return
	}

	blockNumber := big.NewInt(10001)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println("-2-", err)
		return
	}
	fmt.Println("--time:", header.Time)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println("-3-", err)
		return
	}

	fmt.Println("--time:", block.Header().Time)

	// 通过调用 Transactions 方法来读取块中的事务，该方法返回一个 Transaction 类型的列表
	for _, tx := range block.Transactions() {
		// 原生代币转账
		if !(len(tx.Data()) == 0 && !util.IsContractAddress(client, tx.To().Hex())) {
			continue
		}

		fmt.Println("-------------")                 // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("Hash():", tx.Hash().Hex())      // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("Value():", tx.Value().String()) // 10000000000000000
		fmt.Println("Data():", tx.Data())            // []
		fmt.Println("To():", tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// 我们从客户端拿到链ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// 最新
		signer := types.LatestSignerForChainID(chainID)
		from, err := signer.Sender(tx)
		if err != nil {
			log.Fatal("获取交易的发送者失败:", err)
		}
		fmt.Println("from:", from.Hex())

		//// 每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志，以及为“1”（成功）或“0”（失败）的事件结果状态。
		//receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//fmt.Println(receipt.Status) // 1
		//
		//// 才能够 receipt 中获取 log
		//logs, _ := json.Marshal(receipt.Logs)
		//fmt.Println("--------", string(logs))
	}

	return
}

func IsContractAddress(client *ethclient.Client, address string) bool {
	bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(address), nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	return len(bytecode) > 0
}
