package code

const (
	UserMedalMintState_Fail    = -1
	UserMedalMintState_NoMint  = 0
	UserMedalMintState_Commit  = 1 // 提交
	UserMedalMintState_Succeed = 2 // 成功
)

const (
	Sig721Transfer        = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	SigTransfer2          = "0x8988d59efc2c4547ef86c88f6543963bab0cea94f8e486e619c7c3a790db93be"
	SigTransfer3          = "0xd5c97f2e041b2046be3b4337472f05720760a198f4d7d84980b7155eec7cca6f"
	Sig1155TransferSingle = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	Sig1155TransferBatch  = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"
	NullAddress           = "0x0000000000000000000000000000000000000000"
)
