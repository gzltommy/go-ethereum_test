package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "go-ethereum_test/go-ethereum/04.smart_contract/08.read_erc721_contract/erc721_contract/Erc721"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://forno.celo.org")
	if err != nil {
		log.Fatal(err)
	}

	// 获取一个合约实例
	tokenAddress := common.HexToAddress("0xb404e5233aB7E426213998C025f05EaBaBD41Da6")
	instance, err := token.NewErc721(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约的方法
	address := common.HexToAddress("0x068986b94531a5a13da224be26908d001030eae0")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("balance: %v", bal)
}
