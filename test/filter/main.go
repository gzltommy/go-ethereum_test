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
	client, err := ethclient.Dial("https://mainnet.era.zksync.io")
	//client, err := ethclient.Dial("https://zksync-era.blockpi.network/v1/rpc/public")
	if err != nil {
		fmt.Printf("Dial fail.err:%v", err)
		return
	}
	transaction, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash("0x1c5055142cfaf9c651923fda213510a3af513db8b239a28002ac26214bbc7504"))
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
