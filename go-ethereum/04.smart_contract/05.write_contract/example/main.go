package main

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"go-ethereum_test/go-ethereum/04.smart_contract/05.write_contract/example/const"
	"go-ethereum_test/go-ethereum/04.smart_contract/05.write_contract/example/constract"
	"log"
	"math/big"
	"strings"
	"time"
)

func main() {
	userMedal := &UserMedal{
		ID:         0,
		UserId:     0,
		MedalId:    0,
		MedalLevel: 0,
		IsWearing:  0,
		HaveRead:   0,
		MintState:  0,
		DummyId:    0,
		MintHash:   "",
		ChainId:    0,
		Contract:   "",
		TokenId:    "",
		CreateTime: 0,
		UpdateTime: 0,
	}

	nodeUrl := ""
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	chainId := 97
	nftContract := ""
	mintContract := ""
	userAddress := ""
	mintInstance, err := sl.NewSpaceStation(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatal(err)
	}

	mint := func() error {
		txList := make([]userMedalTx, 0, 1)
		defer func() {
			// mint 接口调用完成的勋章都要处理
			go waitHandTx(client, txList) // 异步等待处理
		}()

		//dummyId := uuid.GenerateUUID() // 生成一个唯一数
		dummyId := int64(123456) // 生成一个唯一数
		err, txHash := userMedalMint(client, mintInstance, userMedal, dummyId,
			chainId, nftContract, mintContract, userAddress)
		if err != nil {
			return err
		}

		userMedal.DummyId = dummyId
		userMedal.MintState = code.UserMedalMintState_Commit
		userMedal.ChainId = chainId
		userMedal.Contract = nftContract
		err = SaveUserMedal(userMedal)
		if err != nil {
			return err
		}
		txList = append(txList, userMedalTx{
			NftContract: nftContract,
			ID:          userMedal.ID,
			TxHash:      txHash,
		})
		return nil
	}

	err = mint()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Hour)
}

const (
	medalCastPrev = "12e96e37bd2df8abcac12d6016fa4b87f05d9137065ec3af2cf325dee5c29ed9"
	canTransfer   = true
)

func userMedalMint(
	client *ethclient.Client,
	mintInstance *sl.SpaceStation,
	userMedal *UserMedal,
	dummyId int64,
	chainId int,
	nftContract string,
	mintContract string,
	userAddress string) (error, common.Hash) {

	// 签名
	typedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Claim": {
				{Name: "pid", Type: "uint256"},
				{Name: "claimNFT", Type: "address"},
				{Name: "dummyId", Type: "uint256"},
				{Name: "level", Type: "uint256"},
				{Name: "mintTo", Type: "address"},
				{Name: "royaltyReceiver", Type: "address"},
				{Name: "royaltyFeeNumerator", Type: "uint96"},
				{Name: "canTranster", Type: "bool"},
			},
		},
		Domain: apitypes.TypedDataDomain{
			Name:              "SL",
			Version:           "1.0.0",
			ChainId:           math.NewHexOrDecimal256(int64(chainId)),
			VerifyingContract: mintContract,
			Salt:              "",
		},
		Message: map[string]interface{}{
			"pid":                 math.NewHexOrDecimal256(userMedal.MedalId),
			"claimNFT":            nftContract,
			"dummyId":             math.NewHexOrDecimal256(dummyId),
			"level":               math.NewHexOrDecimal256(int64(userMedal.MedalLevel)),
			"mintTo":              userAddress,
			"royaltyReceiver":     code.NullAddress,
			"royaltyFeeNumerator": math.NewHexOrDecimal256(0),
			"canTranster":         canTransfer,
		},
		PrimaryType: "Claim",
	}

	privateKey, err := crypto.HexToECDSA(medalCastPrev)
	if err != nil {
		return err, common.Hash{}
	}

	signature, err := SignWithEip721(privateKey, typedData)
	if err != nil {
		return err, common.Hash{}
	}

	gasPrice := int64(0)
	//if config.GetConfig().App.Env == "prod" {
	gasPrice = 5
	//}

	auth, err := Author(privateKey, client, 400000, gasPrice)
	if err != nil {
		return err, common.Hash{}
	}

	tx, err := mintInstance.Claim(auth,
		big.NewInt(userMedal.MedalId),           // 勋章 id
		common.HexToAddress(nftContract),        // 勋章合约地址
		big.NewInt(dummyId),                     // 勋章唯一数
		big.NewInt(int64(userMedal.MedalLevel)), // 勋章等级
		common.HexToAddress(userAddress),        // 用户地址
		common.HexToAddress(code.NullAddress),   // 版税接收地址
		big.NewInt(0),                           // 版税
		canTransfer,                             // 是否可转移
		signature,                               // 签名
	)
	if err != nil {
		return err, common.Hash{}
	}

	return nil, tx.Hash()
}

type userMedalTx struct {
	NftContract string
	ID          int64
	TxHash      common.Hash
}

func waitHandTx(client *ethclient.Client, mtxList []userMedalTx) {
	getMintReceipt := func(mTx *userMedalTx) (error, string) {
	Retry:
		receipt, err := client.TransactionReceipt(context.Background(), mTx.TxHash)
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second * 1)
				goto Retry
			}
			return err, ""
		}
		if receipt.Status == 1 {
			// mint 成功
			tokenId := ""
			for _, eLog := range receipt.Logs {
				if strings.EqualFold(eLog.Address.Hex(), mTx.NftContract) &&
					len(eLog.Topics) > 3 && eLog.Topics[0].Hex() == code.Sig721Transfer {
					tokenId = eLog.Topics[3].Big().String()
					break
				}
			}
			return nil, tokenId
		}

		// mint 失败了
		return errors.New("mint fail"), ""
	}

	for i := 0; i < len(mtxList); i++ {
		upData := make(map[string]interface{}, 2)
		mTx := mtxList[i]
		err, tokenId := getMintReceipt(&mTx)
		if err != nil {
			log.Printf("getMintReceipt fail.err:%v", err)
			upData["mint_state"] = code.UserMedalMintState_Fail
		} else {
			upData["mint_state"] = code.UserMedalMintState_Succeed
			upData["token_id"] = tokenId
		}
		err = UpdateUserMedal(mTx.ID, upData)
		if err != nil {
			log.Printf("UpdateUserMedal fail:%v,%+v,%v", mTx.ID, upData, err)
		}
	}
}

type UserMedal struct {
	ID         int64  `gorm:"column:id" json:"id"`
	UserId     int64  `gorm:"column:user_id" json:"user_id"`         //  用户 id
	MedalId    int64  `gorm:"column:medal_id" json:"medal_id"`       //  勋章 id
	MedalLevel uint16 `gorm:"column:medal_level" json:"medal_level"` //  勋章等级（0 表示没有等级区分）
	IsWearing  uint8  `gorm:"column:is_wearing" json:"is_wearing"`   //  是否佩戴（0:未佩戴；1:已佩戴）
	HaveRead   uint8  `gorm:"column:have_read" json:"have_read"`     //  是否阅读（0:未读；1:已读）
	MintState  int8   `gorm:"column:mint_state" json:"mint_state"`   //  mint 状态（0:未上链；1:已 mint）
	DummyId    int64  `gorm:"column:dummy_id" json:"-"`              //  勋章上链时分配的唯一 id
	MintHash   string `gorm:"column:mint_hash" json:"-"`             //  MintHash 值

	ChainId    int    `gorm:"column:chain_id" json:"-"`              //  链
	Contract   string `gorm:"column:contract" json:"-"`              //  合约
	TokenId    string `gorm:"column:token_id" json:"-"`              //  tokenId 值
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` //  创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"` //  更新时间
}

func SaveUserMedal(userMedal *UserMedal) error {
	userMedal.UpdateTime = time.Now().Unix()
	//return DB().Model(&UserMedal{}).Where("id=?", userMedal.ID).Save(&userMedal).Error
	return nil
}

func UpdateUserMedal(id int64, data map[string]interface{}) error {
	data["update_time"] = time.Now().Unix()
	//return DB().Model(&UserMedal{}).Where("id=?", id).Updates(data).Error
	return nil
}
