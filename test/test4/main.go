package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 替换为你的以太坊节点 URL
	ethereumURL := "YOUR_ETHEREUM_NODE_URL"

	// 替换为 ERC1155 合约地址和要查询的代币 ID
	contractAddress := common.HexToAddress("CONTRACT_ADDRESS")
	tokenID := big.NewInt(TOKEN_ID)

	// 替换为要查询的地址
	address := common.HexToAddress("ADDRESS_TO_CHECK")

	// 创建以太坊客户端
	client, err := ethclient.Dial(ethereumURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 创建一个新的上下文
	ctx := context.Background()

	// 调用 ERC1155 合约的 balanceOf 函数来查询代币余额
	balance, err := callBalanceOf(ctx, client, contractAddress, address, tokenID)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	fmt.Printf("Address %s has a balance of %s tokens for token ID %s\n", address.Hex(), balance.String(), tokenID.String())
}

func callBalanceOf(ctx context.Context, client *ethclient.Client, contractAddress common.Address, owner common.Address, tokenID *big.Int) (*big.Int, error) {
	// 替换为 ERC1155 合约的 ABI
	// 你需要使用合约的 ABI 来调用合约函数
	// 这是一个示例 ABI，请根据你的合约进行替换
	abi := "[合约ABI]"

	// 解析合约 ABI
	parsedABI, err := ethereum.JSONStringToABI(abi)
	if err != nil {
		return nil, err
	}

	// 构造调用数据
	data, err := parsedABI.Pack("balanceOf", owner, tokenID)
	if err != nil {
		return nil, err
	}

	// 构造调用消息
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// 发起调用
	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	// 解析调用结果
	var balance big.Int
	err = parsedABI.Unpack(&balance, "balanceOf", result)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}


instance,err := InitInstance(url, contract)
if err!=nil{
t.Log(err)
return
}
result,err := instance.BalanceOf(&bind.CallOpts{},common.HexToAddress(""),big.NewInt(0))
t.Log(result)