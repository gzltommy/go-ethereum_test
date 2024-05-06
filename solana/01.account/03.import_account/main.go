package main

import (
	"fmt"
	"github.com/blocto/solana-go-sdk/types"
	"log"
)

func main() {
	// import a wallet with base58 private key
	account, err := types.AccountFromBase58("8EcaHUYSj85efPS2peniMnobL5UhbNuvC6PSVbbdY5AgzFU9aFTgNfAhgfi56Zb3b6vjspxjWQXjB7bBkfi9d2Q")
	if err != nil {
		log.Fatalf("AccountFromBase58, err: %v", err)
	}
	fmt.Printf("account: %v\n", account.PublicKey.String())

	// import a wallet with bytes slice private key
	accountN := types.NewAccount()
	fmt.Printf("accountN: %v\n", accountN.PublicKey.String())

	account, err = types.AccountFromBytes([]byte(accountN.PrivateKey))
	if err != nil {
		log.Fatalf("AccountFromBytes, err: %v", err)
	}
	fmt.Printf("account: %v\n", account.PublicKey.String())

	// import a wallet with hex private key
	//types.AccountFromHex(accountN.PrivateKey.)
}
