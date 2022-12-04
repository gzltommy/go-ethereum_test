package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	netWork   = "goerli"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("https://%s.infura.io/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6")

	// 获取最新的余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 1190850263533757080

	//读取该区块时的账户余额
	blockNumber := big.NewInt(31229914)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 1195032592258669923

	// wei/10^18
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 1.195032592258669923

	//待处理的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 1190850263533757080
}
