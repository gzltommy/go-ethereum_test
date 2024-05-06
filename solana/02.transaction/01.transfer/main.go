package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/sysprog"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"log"
)

/*
在 Solana 上进行转账之前，你必须遵循 4 个步骤：
（1）：首先，使用 GetLatestBlockhash() 功能从网络上检索最新的区块哈希值。
（2）：用检索到的区块哈希值和交易签字人的公钥制作一个转账信息（转移信息包含交易发送方、接收方和金额的公钥。也可以在一个交易中进行多次转账）。
（3）：用转账信息和交易签字人（交易签字人是为该交易支付费用的账户。可以让另一个账户支付交易费用，但该账户也必须签署该交易）创建一个交易。
（4）：将交易发送到网络（将交易发送到网络后，你会收到它的哈希值，我们可以用它来追踪这笔交易）。
*/
func main() {
	// create a RPC client
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// import a wallet with Devnet balance
	wallet, err := types.AccountFromBytes([]byte{})
	if err != nil {
		log.Fatalf("AccountFromBase58, err: %v", err)
	}

	// （1）：首先，使用 GetLatestBlockhash() 功能从网络上检索最新的区块哈希值
	response, err := c.GetLatestBlockhash(context.TODO())
	if err != nil {
		log.Fatalf("GetLatestBlockhash, err: %v", err)
	}

	// （2）：用检索到的区块哈希值和交易签字人的公钥制作一个转账信息
	// 转移信息包含交易发送方、接收方和金额的公钥。也可以在一个交易中进行多次转账
	message := types.NewMessage(
		types.NewMessageParam{
			FeePayer: wallet.PublicKey, // public key of the transaction signer
			Instructions: []types.Instruction{
				sysprog.Transfer(
					sysprog.TransferParam{
						From:   wallet.PublicKey,                                                           // public key of the transaction sender
						To:     common.PublicKeyFromString("8t88TuqUxDMVpYGHcVoXnBCAH7TPrdZ7ydr4xqcNu2Ym"), // wallet address of the transaction receiver
						Amount: 1e9,                                                                        // transaction amount in lamport
					}),
			},
			RecentBlockhash: response.Blockhash, // recent block hash
		})

	// （3）：用转账信息和交易签字人创建一个交易。
	// 交易签字人是为该交易支付费用的账户。可以让另一个账户支付交易费用，但该账户也必须签署该交易
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: message,
		Signers: []types.Account{wallet, wallet},
	})
	if err != nil {
		log.Fatalf("NewTransaction, err: %v", err)
	}

	// （4）：将交易发送到网络（将交易发送到网络后，你会收到它的哈希值，我们可以用它来追踪这笔交易）。
	txhash, err := c.SendTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatalf("SendTransaction, err: %v", err)
	}

	fmt.Println("Transaction Hash:", txhash)
}
