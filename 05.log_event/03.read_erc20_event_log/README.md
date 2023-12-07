# 解析日志共有2种方式

## 1.使用 ABI 数据 + UnpackIntoInterface 方法解析

```azure
contractAbi, err := abi.JSON(strings.NewReader(token.ERC20MetaData.ABI))
if err != nil {
	log.Fatal(err)
}

...

var transferEvent LogTransfer
err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data) // 参数是 vlog.Data
if err != nil {
	log.Fatal(err)
}

// UnpackIntoInterface 方法不会解析出 index 类参数，需要用如下方式 
transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
```

## 2.使用 NewXXXFilter + ParseXXX 方式解析

```azure
// 都传 0 值
inst, err := token.NewERC20Filterer(common.HexToAddress("0x"), nil)
if err != nil {
	log.Fatal(err)
}

...

// 该方式会自动解析出 index 参数
transferEvent,err := inst.ParseTransfer(vLog) //参数是 vLog
if err != nil {
	log.Fatal(err)
}    
```