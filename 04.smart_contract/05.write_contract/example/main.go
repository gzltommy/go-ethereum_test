package example

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

type reqUserMedalMintTT struct {
	UserMedalID int64 `form:"user_medal_id" binding:"required"` // 传 -1 表示一键 mint 所有
	UserID      int64 `form:"user_id" binding:"required"`
}

func UserMedalMintTT(c *gin.Context) {
	var req reqUserMedalMintTT
	if err := c.ShouldBind(&req); err != nil {
		render.Json(c, render.ErrParams, err)
		return
	}

	userAddress, err := model.FindUserAddressByUID(req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Json(c, render.NoBindingWallet, err.Error())
			return
		}
		render.Json(c, render.Failed, err.Error())
		return
	}

	var userMedalList []*model.UserMedal
	if req.UserMedalID == -1 {
		var err error
		userMedalList, err = model.FindUserNoMintMedal(req.UserID)
		if err != nil {
			render.Json(c, render.Failed, err.Error())
			return
		}
	} else {
		userMedal, err := model.FindUserMedalById(req.UserMedalID)
		if err != nil {
			render.Json(c, render.Failed, err.Error())
			return
		}
		if userMedal.MintState == code.UserMedalMintState_Commit {
			// 已经 mint 过了
			render.Json(c, render.Failed, "already minted")
			return
		}
		userMedalList = append(userMedalList, userMedal)
	}

	nftContract, err := model.FindMarketContractByUsageV3("medal_nft")
	if err != nil {
		render.Json(c, render.NotFound, err)
		return
	}

	mintContract, err := model.FindMarketContractByUsageV3("medal_mint")
	if err != nil {
		render.Json(c, render.NotFound, err)
		return
	}

	chain := model.GetChainByChainId(nftContract.ChainId)
	client, err := ethclient.Dial(chain.Http)
	if err != nil {
		log.Log.Errorf("Init ethClient[%s] err:%v", chain.Http, err)
		render.Json(c, render.Failed, err)
		return
	}

	mintInstance, err := SecondLiveMedal.NewSpaceStation(common.HexToAddress(mintContract.Address), client)
	if err != nil {
		render.Json(c, render.Failed, err)
		return
	}

	txList := make([]userMedalTx, 0, len(userMedalList))
	defer func() {
		// mint 接口调用完成的勋章都要处理
		go waitHandTx(client, txList) // 异步等待处理
	}()
	for _, userMedal := range userMedalList {
		dummyId := uuid.GenerateUUID() // 生成一个唯一数
		err, txHash := userMedalMint(client, mintInstance, userMedal, dummyId,
			nftContract.ChainId, nftContract.Address, mintContract.Address, userAddress)
		if err != nil {
			render.Json(c, render.Failed, "mint fail")
			return
		}

		userMedal.DummyId = dummyId
		userMedal.MintState = code.UserMedalMintState_Commit
		userMedal.ChainId = nftContract.ChainId
		userMedal.Contract = nftContract.Address
		err = model.SaveUserMedal(userMedal)
		if err != nil {
			log.Log.WithAlarm().Errorf("SaveUserMedal fail:%+v,%v", userMedal, err)
			render.Json(c, render.Failed, err)
			return
		}
		txList = append(txList, userMedalTx{
			NftContract: nftContract.Address,
			ID:          userMedal.ID,
			TxHash:      txHash,
		})
	}
	render.Json(c, render.Ok, nil)
	return
}

const (
	medalCastPrev = "40e96e37bd2df8abcac12d6016fa4b87f05d9137065ec3af2cf325dee5c29ed9"
	canTransfer   = true
)

func userMedalMint(
	client *ethclient.Client,
	mintInstance *SecondLiveMedal.SpaceStation,
	userMedal *model.UserMedal,
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
			Name:              "SecondLive",
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

	signature, err := eip712_sign.SignWithEip721(privateKey, typedData)
	if err != nil {
		log.Log.Errorf("SignWithEip721 err:%v", err)
		return err, common.Hash{}
	}

	gasPrice := int64(0)
	if config.GetConfig().App.Env == "prod" {
		gasPrice = 5
	}

	auth, err := contract.Author(privateKey, client, 400000, gasPrice)
	if err != nil {
		log.Log.Errorf("Create Auth err:%v", err)
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
		log.Log.WithAlarm().Errorf("mint medal err:%v", err)
		return err, common.Hash{}
	}

	return nil, tx.Hash()
}
