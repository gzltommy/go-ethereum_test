package main

import (
	"fmt"
	util "go-ethereum_test/10.utils"
	"log"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/77055bf09f924b509eee20cfb4d5d5b9")
	//client, err := ethclient.Dial("https://arb1.arbitrum.io/rpc")
	if err != nil {
		fmt.Println("-1-", err)
		return
	}

	//blockNumber := big.NewInt(int64(202165))
	blockNumber := int64(38135097)
	block, err := util.GetBlockByBlockNumber(context.Background(), client, blockNumber)
	if err != nil {
		fmt.Println("-3-", err)
		return
	}
	fmt.Println("--time:", block.BlockTime)
	fmt.Printf("--,bnum:%v,txlen:%v\n", blockNumber, len(block.Transactions))

	// 通过调用 Transactions 方法来读取块中的事务，该方法返回一个 Transaction 类型的列表
	i := 0
	for _, tx := range block.Transactions {
		i++

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash)
		if err != nil {
			log.Fatal(err)
		}
		if receipt.Status == 1 {
			if err != nil {
				fmt.Printf("--i:%v,tx:%v,to:%v\n", i, tx.Hash.String(), tx.To)
			} else {
				fmt.Printf("==i:%v,tx:%v,from:%v,to:%v,data:%v\n", i, tx.Hash.String(), tx.From, tx.To, len(tx.Data)) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
			}
		}
	}
	return
}
