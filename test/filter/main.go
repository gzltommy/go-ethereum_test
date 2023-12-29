package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	if err != nil {
		fmt.Printf("Dial fail.err:%v", err)
		return
	}
	transaction, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash("0x13be57d4511a56b3cf16379d3518bff1f3285b3c70aebf0ada72a47adbd647e4"))
	if err != nil {
		fmt.Printf("TransactionReceipt fail.err:%v", err)
		return
	}
	if isPending {
		return
	}

	header, err := client.HeaderByNumber(context.Background(), big.NewInt(2393693))
	if err != nil {
		fmt.Printf("HeaderByNumber fail.err:%v", err)
		return
	}

	fmt.Println(header.Number.String())
	fmt.Println(time.Unix(int64(header.Time), 0))

	msg, err := transaction.AsMessage(types.LatestSignerForChainID(big.NewInt(int64(324))), nil)
	if err != nil {
		fmt.Printf("AsMessage fail.err:%v", err)
		return
	}
	_ = header
	fmt.Println(msg.From().Hex())
	fmt.Println(msg.To().Hex())
}
