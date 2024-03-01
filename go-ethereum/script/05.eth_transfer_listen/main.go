package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum_test/go-ethereum/utils"
)

func main() {
	client, err := ethclient.Dial("https://rpc.scroll.io")
	if err != nil {
		fmt.Println("-1-", err)
		return
	}

	blockNumber := int64(3754207)
	block, err := util.GetBlockByBlockNumber(context.Background(), client, blockNumber)
	if err != nil {
		fmt.Println("-3-", err)
		return
	}
	fmt.Println("--time:", block.BlockTime)
	fmt.Printf("--bnum:%v\n", blockNumber)

	return
}
