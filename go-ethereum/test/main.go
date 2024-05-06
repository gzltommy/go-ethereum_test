package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://rpc.bsquared.network/")
	if err != nil {
		log.Fatal("--1-", err)
	}

	// 还可以使用 TransactionByHash 在给定具体事务哈希值的情况下直接查询单个事务。
	header, err := client.HeaderByNumber(context.Background(), big.NewInt(129439))
	if err != nil {
		log.Fatal("--2-", err)
	}

	fmt.Println("-----ok-------", header.Time) // 1

}
