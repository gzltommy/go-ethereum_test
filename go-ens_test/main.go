package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

func main() {
	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("https://eth-mainnet.nodereal.io/v1/fa6e1351d6064b41b40cfd6c0500d0ec")
	if err != nil {
		panic(err)
	}

	// Resolve a name to an address.
	domain := "zorromiaotommy.eth"
	address, err := ens.Resolve(client, domain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

	//address := common.HexToAddress("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6")

	// Reverse resolve an address to a name.
	reverse, err := ens.ReverseResolve(client, address)
	if err != nil {
		panic(err)
	}
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
}
