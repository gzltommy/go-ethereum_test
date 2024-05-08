package main

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// 创建钱包
func main() {
	// Create a new account:
	account := solana.NewWallet()
	fmt.Println("account private key:", account.PrivateKey)
	fmt.Println("account public key:", account.PublicKey())

	// Create a new RPC client:
	client := rpc.New(rpc.TestNet_RPC)

	// Airdrop 1 SOL to the new account:
	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("airdrop transaction signature:", out)
}

// LoadWallet 加载/解析私钥和公钥
func LoadWallet() {
	{
		// Load private key from a json file generated with
		// $ solana-keygen new --outfile=standard.solana-keygen.json
		privateKey, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/standard.solana-keygen.json")
		if err != nil {
			panic(err)
		}
		fmt.Println("private key:", privateKey.String())
		// To get the public key, you need to call the `PublicKey()` method:
		publicKey := privateKey.PublicKey()
		// To get the base58 string of a public key, you can call the `String()` method:
		fmt.Println("public key:", publicKey.String())
	}

	{
		// Load private key from base58:
		{
			privateKey, err := solana.PrivateKeyFromBase58("66cDvko73yAf8LYvFMM3r8vF5vJtkk7JKMgEKwkmBC86oHdq41C7i1a2vS3zE1yCcdLLk6VUatUb32ZzVjSBXtRs")
			if err != nil {
				panic(err)
			}
			fmt.Println("private key:", privateKey.String())
			fmt.Println("public key:", privateKey.PublicKey().String())
		}
		// OR:
		{
			privateKey := solana.MustPrivateKeyFromBase58("66cDvko73yAf8LYvFMM3r8vF5vJtkk7JKMgEKwkmBC86oHdq41C7i1a2vS3zE1yCcdLLk6VUatUb32ZzVjSBXtRs")
			_ = privateKey
		}
	}

	{
		// Generate a new key pair:
		{
			privateKey, err := solana.NewRandomPrivateKey()
			if err != nil {
				panic(err)
			}
			_ = privateKey
		}
		{
			{ // Generate a new private key (a Wallet struct is just a wrapper around a private key)
				account := solana.NewWallet()
				_ = account
			}
		}
	}

	{
		// Parse a public key from a base58 string:
		{
			publicKey, err := solana.PublicKeyFromBase58("F8UvVsKnzWyp2nF8aDcqvQ2GVcRpqT91WDsAtvBKCMt9")
			if err != nil {
				panic(err)
			}
			_ = publicKey
		}
		// OR:
		{
			publicKey := solana.MustPublicKeyFromBase58("F8UvVsKnzWyp2nF8aDcqvQ2GVcRpqT91WDsAtvBKCMt9")
			_ = publicKey
		}
	}
}
