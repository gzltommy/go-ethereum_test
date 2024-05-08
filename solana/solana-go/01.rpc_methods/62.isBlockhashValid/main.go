package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	blockHash := solana.MustHashFromBase58("J7rBdM6AecPDEZp8aPq5iPSNKVkU5Q76F3oAV4eW5wsW")
	out, err := client.IsBlockhashValid(
		context.TODO(),
		blockHash,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Value) // true or false

	fmt.Println("is blockhash valid:", out.Value)
}
