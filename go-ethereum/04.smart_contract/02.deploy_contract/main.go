package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	store "go-ethereum_test/go-ethereum/04.smart_contract/02.deploy_contract/contract/store"
	"log"
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

	// 1.加载您的私钥
	privateKey, err := crypto.HexToECDSA("6cdd2cd6a8f813c33ac0c59a92798da9c0e490e17da96fbb9dc86fd1e77d23e2")
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个有配置密匙的交易发送器（Transactor）
	authTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	authTransactor.Nonce = big.NewInt(int64(nonce))
	authTransactor.Value = big.NewInt(0)      // in wei
	authTransactor.GasLimit = uint64(3000000) // 设置高一点，不然很可能失败
	authTransactor.GasPrice = gasPrice

	input := "1.0"                                                                 // 合约中，构造函数初始化值
	address, tx, instance, err := store.DeployStore(authTransactor, client, input) // 生成的 Go 合约文件提供了部署方法。 部署方法名称始终以单词 Deploy 开头，后跟合约名称
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x3f6c4f9FDd3AAdfE545D45221477A1F553404Ae5
	fmt.Println(tx.Hash().Hex()) // 0x7530e697016d11e7683381e791091967d1f76f9ad96e186546aa5fc53f7338c6
	// 你可以用事务哈希来在 Etherscan 上查询合约的部署状态 https://goerli.etherscan.io/tx/0x7530e697016d11e7683381e791091967d1f76f9ad96e186546aa5fc53f7338c6
	_ = instance
}
