package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
	"log"
	"math/big"
)

func main() {
	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("https://eth-mainnet.nodereal.io/v1/fa6e1351d6064b41b40cfd6c0500d0ec")
	if err != nil {
		fmt.Printf("ethclient.Dial.err:%v", err)
		return
	}

	// Resolve a name to an address.
	{
		domain := "zorromiaotommy.eth"
		address, err := ens.Resolve(client, domain)
		if err != nil {
			fmt.Printf("Resolve fail.err:%v", err)
			return
		}
		fmt.Printf("Address of %s is %s\n", domain, address.Hex())
	}

	// Reverse resolve an address to a name.
	{
		address := common.HexToAddress("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6")
		reverse, err := ens.ReverseResolve(client, address)
		if err != nil {
			fmt.Printf("ReverseResolve fail.err:%v", err)
			return
		}
		if reverse == "" {
			fmt.Printf("%s has no reverse lookup\n", address.Hex())
		} else {
			fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
		}
	}

	return

	// 创建子名称
	{
		// 创建一个有配置密匙的交易发送器（Transactor）
		// 1.加载您的私钥。写入智能合约需要我们用私钥来对交易事务进行签名。
		privateKey, err := crypto.HexToECDSA("asdf5783e8eb229fe4cb9e33afe76asdf719fe056b64294cf14e9fca35d8014bbac450")
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

		authTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Fatal(err)
		}

		authTransactor.Nonce = big.NewInt(int64(nonce))
		authTransactor.Value = big.NewInt(0)      // in wei
		authTransactor.GasLimit = uint64(3000000) // in units
		authTransactor.GasPrice = gasPrice

		registry, err := ens.NewRegistry(client)
		if err != nil {
			fmt.Printf("NewRegistry fail.err:%v", err)
			return
		}
		transaction, err := registry.SetSubdomainOwner(authTransactor, "zorromiaotommy.eth", "tommy", common.HexToAddress("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("SetSubdomainOwner tx:%s", transaction.Hash().Hex())
	}
	// 设置解析器
	{
		// 创建一个有配置密匙的交易发送器（Transactor）
		// 1.加载您的私钥。写入智能合约需要我们用私钥来对交易事务进行签名。
		privateKey, err := crypto.HexToECDSA("asdf5783e8eb229fe4cb9e33afe76asdf719fe056b64294cf14e9fca35d8014bbac450")
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

		authTransactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Fatal(err)
		}

		authTransactor.Nonce = big.NewInt(int64(nonce))
		authTransactor.Value = big.NewInt(0)      // in wei
		authTransactor.GasLimit = uint64(3000000) // in units
		authTransactor.GasPrice = gasPrice

		registry, err := ens.NewRegistry(client)
		if err != nil {
			fmt.Printf("NewRegistry fail.err:%v", err)
			return
		}
		// opts are go-ethereum's bind.TransactOpts
		transaction, err := registry.SetResolver(authTransactor, "tommy.zorromiaotommy.eth", common.HexToAddress("0x1234..."))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("SetResolver tx:%s", transaction.Hash().Hex())
	}

}
