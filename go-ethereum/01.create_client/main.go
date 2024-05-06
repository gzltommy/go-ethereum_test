package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	netWork   = "mainnet"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	// https://mainnet.infura.io/v3/YOUR-PROJECT
	// https://sepolia.infura.io/v3/
	client, err := ethclient.Dial(fmt.Sprintf("https://%s.sepolia.infura.io/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
