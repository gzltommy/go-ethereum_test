package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	netWork   = "kovan"
	projectID = "b2c8412b5acc4138a27f524ee4d6d18f"
)

func main() {
	client, err := ethclient.Dial(fmt.Sprintf("https://%s.infura.io/v3/%s", netWork, projectID))
	if err != nil {
		log.Fatal(err)
	}

	// 1.加载您的私钥
	privateKey, err := crypto.HexToECDSA("debe1d34d97ad7ad00d4eaff4c291c9261a11c7f9432b867d7a987f698bbc33b")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts[0].Data()
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf(rawTxHex) // f86...772
}
