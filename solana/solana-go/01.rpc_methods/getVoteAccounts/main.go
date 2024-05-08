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

	votePubkey := solana.MustPublicKeyFromBase58("vot33MHDqT6nSwubGzqtc6m16ChcUywxV7tNULF19Vu")
	out, err := client.GetVoteAccounts(
		context.TODO(),
		&rpc.GetVoteAccountsOpts{
			VotePubkey: &votePubkey,
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}