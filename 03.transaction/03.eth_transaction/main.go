package main

import (
	"crypto/ecdsa"
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

	/*
		2.之后我们需要获得帐户的随机数(nonce)。 每笔交易都需要一个nonce。 根据定义，nonce是仅使用一次的数字。 如果是发送交易的新帐户，
		则该随机数将为“0”。 来自帐户的每个新事务都必须具有前一个nonce增加1的nonce。很难对所有nonce进行手动跟踪，
		于是ethereum客户端提供一个帮助方法PendingNonceAt，它将返回你应该使用的下一个nonce。
	*/
	// 发送的帐户的公共地址 - 这个我们可以从私钥派生。
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // PendingNonceAt，它将返回你应该使用的下一个 nonce。
	if err != nil {
		log.Fatal(err)
	}

	/*
		3.下一步是设置我们将要转移的ETH数量。 但是我们必须将ETH以太转换为wei，因为这是以太坊区块链所使用的。 以太网支持最多18个小数位，
		因此1个ETH为1加18个零。
	*/
	value := big.NewInt(310000000000000000) // in wei (0.31 eth)
	gasLimit := uint64(21000)               // ETH 转账的燃气应设上限为“21000”单位。

	// 燃气价格必须以 wei 为单位设定。 在撰写本文时，将在一个区块中比较快的打包交易的燃气价格为30 gwei。
	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	//然而，燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。
	//go-ethereum客户端提供SuggestGasPrice函数，用于根据’x’个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	/*
		4.现在我们最终可以通过导入go-ethereumcore/types包并调用NewTransaction来生成我们的未签名以太坊事务，这个函数需要接收 nonce，
		地址，值，燃气上限值，燃气价格和可选发的数据。 发送ETH的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。
	*/
	toAddress := common.HexToAddress("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	/*
		5. 下一步是使用发件人的私钥对事务进行签名。 为此，我们调用 SignTx 方法，该方法接受一个未签名的事务和我们之前构造的私钥。 SignTx 方法需要 EIP155 签名者，这个也需要我们先从客户端拿到链 ID。
	*/
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	/*
		6.最后通过在客户端上调用“SendTransaction”来将已签名的事务广播到整个网络。
	*/
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // 0xacd776e998fd27a30bafd927b36158af43889d52da9815135fc6a8cef95682a3

	// 然后你可以去 Etherscan 看交易的确认过程: https://kovan.etherscan.io/tx/0xacd776e998fd27a30bafd927b36158af43889d52da9815135fc6a8cef95682a3
}
