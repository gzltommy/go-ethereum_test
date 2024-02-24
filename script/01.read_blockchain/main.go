package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://scroll-alphanet.public.blastapi.io")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(2043125)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 通过调用 Transactions 方法来读取块中的事务，该方法返回一个 Transaction 类型的列表
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// 我们从客户端拿到链ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("network id", chainID) // 1

		// 为了读取发送方的地址，我们需要在事务上调用 AsMessage，它返回一个 Message 类型，其中包含一个返回 sender（from）地址的函数
		////if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		//if msg, err := tx.AsMessage(types.LatestSignerForChainID(chainID), nil); err == nil {
		//	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		//}

		// 最新
		//signer := types.MakeSigner(params.MainnetChainConfig, block.Number(), block.Time())
		//from, err := signer.Sender(tx)
		//if err != nil {
		//	log.Fatal("获取交易的发送者失败:", err)
		//}
		//fmt.Println("交易的发送者是:", from.Hex())

		// 每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志，以及为“1”（成功）或“0”（失败）的事件结果状态。
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1

		logs, _ := json.Marshal(receipt.Logs)
		fmt.Println("--------", string(logs)) // 1
	}
}
