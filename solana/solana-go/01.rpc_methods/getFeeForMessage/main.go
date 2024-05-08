package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetFeeForMessage(
		context.Background(),
		"AQABAgIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEBAQAA",
		rpc.CommitmentProcessed,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(example)
}
