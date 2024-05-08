package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	pubKey := solana.MustPublicKeyFromBase58("7HZaCWazgTuuFuajxaaxGYbGnyVKwxvsJKue1W4Nvyro")
	out, err := client.GetTokenAccountsByOwner(
		context.TODO(),
		pubKey,
		&rpc.GetTokenAccountsConfig{
			Mint: solana.WrappedSol.ToPointer(),
		},
		&rpc.GetTokenAccountsOpts{
			Encoding: solana.EncodingBase64Zstd,
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)

	{
		tokenAccounts := make([]token.Account, 0)
		for _, rawAccount := range out.Value {
			var tokAcc token.Account

			data := rawAccount.Account.Data.GetBinary()
			dec := bin.NewBinDecoder(data)
			err := dec.Decode(&tokAcc)
			if err != nil {
				panic(err)
			}
			tokenAccounts = append(tokenAccounts, tokAcc)
		}
		spew.Dump(tokenAccounts)
	}
}
