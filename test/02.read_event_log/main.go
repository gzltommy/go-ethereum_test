package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var AvatarMintABIJson = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nftID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dummyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_mintTo\",\"type\":\"address\"}],\"name\":\"EventClaim\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dummyId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

//var AvatarMintABIJson = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"_nftID","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_dummyId","type":"uint256"},{"indexed":false,"internalType":"address","name":"_mintTo","type":"address"}],"name":"EventClaim","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"signer","type":"address"}],"name":"UpdateSigner","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"_startTime","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_endTime","type":"uint256"}],"name":"UpdateTime","type":"event"},{"inputs":[{"internalType":"uint256[][]","name":"_attributes","type":"uint256[][]"}],"name":"addSuits","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_dummyId","type":"uint256"},{"internalType":"address","name":"_mintTo","type":"address"},{"internalType":"bytes","name":"_signature","type":"bytes"}],"name":"claim","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"editorNFT","outputs":[{"internalType":"contractISecondLiveEditor","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"endTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_owner","type":"address"},{"internalType":"address","name":"_signer","type":"address"},{"internalType":"contractISecondLiveEditor","name":"_editorNFT","type":"address"},{"internalType":"uint256","name":"_startTime","type":"uint256"},{"internalType":"uint256","name":"_endTime","type":"uint256"},{"internalType":"uint256[][][]","name":"_attributes","type":"uint256[][][]"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"isClaimed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"numClaimed","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"resetSuits","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_signer","type":"address"}],"name":"setSinger","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"signer","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"suitIndex","type":"uint256"}],"name":"singleCountBySuitIndex","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"startTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"suitIndex","type":"uint256"},{"internalType":"uint256","name":"singleIndex","type":"uint256"}],"name":"suitByIndex","outputs":[{"components":[{"internalType":"uint256","name":"rule","type":"uint256"},{"internalType":"uint256","name":"quality","type":"uint256"},{"internalType":"uint256","name":"format","type":"uint256"},{"internalType":"uint256","name":"extra","type":"uint256"}],"internalType":"structISecondLiveEditor.Attribute","name":"_attribute","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"suitLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_startTime","type":"uint256"},{"internalType":"uint256","name":"_endTime","type":"uint256"}],"name":"updateTime","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"userClaimed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`

var AvatarMintABI abi.ABI

type AvatarMintEventClaim struct {
	MintTo  common.Address
	NftID   *big.Int
	DummyId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

func main() {
	// 注意国内要设置代理才能连接
	// 拨打启用 websocket 的以太坊客户端
	client, err := ethclient.Dial("https://test-rpc.combonetwork.io")
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个过滤查询
	// 指定我们想过滤的区块范围并指定从中读取此日志的合约地址
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(4561778), // FIXME:这个区块的日志居然超过了 10000 条，无法正常取回
		ToBlock:   big.NewInt(4561778),
		Addresses: []common.Address{
			common.HexToAddress("0x2aDd449582863e75B81338244546B886724AB8b9"),
		},
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x7cf77399f3e9777f2c76c7cbcf2faabddcfc6a99908aed9cc5ad9a99762bc65f"),
			},
		},
	}

	// 查询将返回所有的匹配事件日志，返回的所有日志将是 ABI 编码的
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, eLog := range logs {
		// 读取日志结构体中的附件信息（区块摘要，区块号和交易摘要）
		//fmt.Println(vLog.BlockNumber)     // 32085598
		//fmt.Println(vLog.TxHash.Hex())    // 0xa766aa0a9fc5e5e969b2e02182b8f04115de8ccb380913b903f696055a51ef5b
		fmt.Println(eLog.Topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

		var event AvatarMintEventClaim
		err = AvatarMintABI.UnpackIntoInterface(&event, "EventClaim", eLog.Data)
		if err != nil {
			fmt.Println("UnpackIntoInterface fail.", err)
			return
		}
		buf, _ := json.Marshal(event)

		fmt.Println(string(buf))
	}
}
