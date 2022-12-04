package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethereum_test/04.smart_contract/02.deploy_contract/contract/store" // 导入包
	"log"
	"math/big"
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

	// 1.加载您的私钥。写入智能合约需要我们用私钥来对交易事务进行签名。
	privateKey, err := crypto.HexToECDSA("6cdd2cd6a8f813c33ac0c59a92798da9c0e490e17da96fbb9dc86fd1e77d23e2")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 查询当前的 nonce 值
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 查询当前的燃气价格
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
	authTransactor.GasLimit = uint64(3000000) // in units
	authTransactor.GasPrice = gasPrice

	// 加载合约
	contractAddress := common.HexToAddress("0x3f6c4f9FDd3AAdfE545D45221477A1F553404Ae5")
	instance, err := store.NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	//copy(key[:], []byte("tommy"))
	copy(key[:], []byte("miao"))
	copy(value[:], []byte("苗"))
	//copy(value[:], []byte("汤米"))

	// 调用合约的写方法
	tx, err := instance.SetItem(authTransactor, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b
	// 你可以用事务哈希来在Etherscan上查询合约的部署状态 https://goerli.etherscan.io/tx/0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
