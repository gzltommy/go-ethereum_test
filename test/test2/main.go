package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	bscAPIURL       = "https://bsc-dataseed.binance.org/"
	contractAddress = "0xYourContractAddressHere" // 替换为您要查询的智能合约地址
)

func main() {
	// 获取智能合约的交易历史
	transactions, err := getContractTransactions(contractAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 提取交易中的钱包地址
	uniqueAddresses := make(map[string]bool)
	for _, tx := range transactions {
		fromAddress := tx.From
		toAddress := tx.To
		uniqueAddresses[fromAddress] = true
		uniqueAddresses[toAddress] = true
	}

	// 打印筛选后的钱包地址
	fmt.Println("钱包地址与合约交互的列表：")
	for address := range uniqueAddresses {
		fmt.Println(address)
	}
}

// 从 BSC 区块链获取智能合约的交易历史
func getContractTransactions(contractAddress string) ([]Transaction, error) {
	url := fmt.Sprintf("%s/api/v1/contract/%s?module=account&action=txlist&sort=desc", bscAPIURL, contractAddress)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result ContractTransactionsResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}

// 结构体定义用于解析 BSC API 响应
type ContractTransactionsResponse struct {
	Result []Transaction `json:"result"`
}

// 结构体定义用于解析交易数据
type Transaction struct {
	From string `json:"from"`
	To   string `json:"to"`
}
