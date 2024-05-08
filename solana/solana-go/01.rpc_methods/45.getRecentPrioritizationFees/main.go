package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	out, err := client.GetRecentPrioritizationFees(
		context.TODO(),
		[]solana.PublicKey{
			solana.MustPublicKeyFromBase58("q5BgreVhTyBH1QCeriVb7kQYEPneanFXPLjvyjdf8M3"),
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
