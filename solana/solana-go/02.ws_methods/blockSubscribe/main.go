package main

import (
	"context"
	"github.com/gagliardetto/solana-go"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

// 此订阅被视为不稳定，仅当验证程序使用 --rpc-pubsub-enable-block-subscription 标志启动时才可用。此订阅的格式将来可能会更改。
func main() {
	client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
	if err != nil {
		panic(err)
	}
	program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin") // serum

	// todo:实现了吗？？？？？？？？？？？？
	option := ws.BlockSubscribeOpts{
		Commitment:                     rpc.CommitmentFinalized,
		Encoding:                       solana.EncodingBase64,
		TransactionDetails:             rpc.TransactionDetailsFull,
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
