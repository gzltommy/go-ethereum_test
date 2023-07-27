package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "go-ethereum_test/04.smart_contract/08.read_erc721_contract/erc721_contract/Erc721"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://bsc-mainnet.nodereal.io/v1/987c2644eafa4dbeba8155e0db5ce956")
	if err != nil {
		log.Fatal(err)
	}

	// 获取一个合约实例
	tokenAddress := common.HexToAddress("0x87a218Ae43C136B3148A45EA1A282517794002c8")
	instance, err := token.NewErc721(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约的方法
	address := common.HexToAddress("0x2b28628404df4941dcc08fe8efd76454f9473a0c")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("balance: %f", bal)
}
