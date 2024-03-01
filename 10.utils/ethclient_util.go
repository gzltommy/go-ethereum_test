package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type rpcBlock struct {
	Number       string           `json:"number"`
	Timestamp    string           `json:"timestamp"`
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
}

type Block struct {
	BlockNumber  int64
	BlockTime    int64
	Hash         common.Hash
	Transactions []transaction
}

type rpcTransaction struct {
	Hash             common.Hash `json:"hash"`
	BlockHash        common.Hash `json:"blockHash"`
	BlockNumber      string      `json:"blockNumber"`
	From             string      `json:"from"`
	Gas              string      `json:"gas"`
	GasPrice         string      `json:"gasPrice"`
	Input            string      `json:"input"`
	To               string      `json:"to"`
	TransactionIndex string      `json:"transactionIndex"`
	Value            string      `json:"value"`
}

type transaction struct {
	Hash             common.Hash
	BlockHash        common.Hash
	BlockNumber      int64
	From             string
	Gas              *big.Int
	GasPrice         *big.Int
	Data             []byte
	To               string
	TransactionIndex float64
	Value            *big.Int
}

func GetBlockByBlockNumber(ctx context.Context, client *ethclient.Client, blockNumber int64) (*Block, error) {
	var resultBlock Block
	var raw json.RawMessage
	err := client.Client().CallContext(ctx, &raw, "eth_getBlockByNumber", toBlockNumArg(big.NewInt(blockNumber)), true)
	if err != nil {
		return nil, err
	}

	var head *types.Header
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	var body rpcBlock
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	var txList []transaction
	for _, b := range body.Transactions {
		txList = append(txList, convertToTransaction(b))
	}
	number, _ := big.NewInt(0).SetString(body.Number[2:], 16)
	blockTime, _ := big.NewInt(0).SetString(body.Timestamp[2:], 16)
	resultBlock.BlockNumber = number.Int64()
	resultBlock.BlockTime = blockTime.Int64()
	resultBlock.Hash = body.Hash
	resultBlock.Transactions = txList
	return &resultBlock, nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}

func convertToTransaction(i rpcTransaction) transaction {
	blockNumber, _ := big.NewInt(0).SetString(i.BlockNumber[2:], 16)
	gas, _ := big.NewInt(0).SetString(i.Gas[2:], 16)
	gasPrice, _ := big.NewInt(0).SetString(i.GasPrice[2:], 16)
	txIndex, _ := big.NewInt(0).SetString(i.TransactionIndex[2:], 16)
	value, _ := big.NewInt(0).SetString(i.Value[2:], 16)
	result := transaction{
		Hash:             i.Hash,
		BlockHash:        i.BlockHash,
		BlockNumber:      blockNumber.Int64(),
		From:             i.From,
		Gas:              gas,
		GasPrice:         gasPrice,
		Data:             common.Hex2Bytes(i.Input[2:]),
		To:               i.To,
		TransactionIndex: float64(txIndex.Int64()),
		Value:            value,
	}
	return result
}
