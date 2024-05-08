package main

import (
	"context"
	"github.com/gagliardetto/solana-go"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
	if err != nil {
		panic(err)
	}
	program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin") // serum
	option := ws.BlockSubscribeOpts{
		Commitment:                     "",
		Encoding:                       solana.EncodingBase64,
		TransactionDetails:             "",
		Rewards:                        nil,
		MaxSupportedTransactionVersion: nil,
	}
	{
		sub, err := client.BlockSubscribe(
			ws.NewBlockSubscribeFilterMentionsAccountOrProgram(program),
			&option,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv()
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
}
