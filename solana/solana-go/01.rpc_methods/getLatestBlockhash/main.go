package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	// 此方法仅在 solana-core v1.9 或更高版本中可用。solana-core v1.8 请使用 getRecentBlockhash
	example, err := client.GetLatestBlockhash(
		context.Background(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(example)
}
