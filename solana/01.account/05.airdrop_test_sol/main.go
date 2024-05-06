package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"log"
)

// 获取开发用的 Devnet 币
func main() {
	// create a RPC client
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// create a new wallet
	wallet := types.NewAccount()

	fmt.Println("account:", wallet.PublicKey.String())

	// request for 1 SOL airdrop using RequestAirdrop()
	txhash, err := c.RequestAirdrop(
		context.TODO(),              // request context
		wallet.PublicKey.ToBase58(), // wallet address requesting airdrop
		1e9,                         // amount of SOL in lamport
	)
	/*
		RequestAirdrop() 函数接受以 lamport 为单位的 amount 参数，这是 SOL 的最小分数单位，类似于比特币的 Satoshi
		1 lamport ~ 0.000000001 SOL
	*/

	// check for errors
	if err != nil {
		log.Fatalf("RequestAirdrop failed.err: %v", err)
	}

	fmt.Println("Transaction Hash:", txhash)
}
