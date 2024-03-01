# 解析日志共有2种方式

## 1.使用 ABI 数据 + UnpackIntoInterface 方法解析

```go
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

```go
// 都传 0 值
inst, err := token.NewERC20Filterer(common.HexToAddress("0x"), nil)
if err != nil {
    log.Fatal(err)
}

...

// 该方式会自动解析出 index 参数
transferEvent, err := inst.ParseTransfer(vLog) //参数是 vLog
if err != nil {
    log.Fatal(err)
}
```

> 注意：该方式不需要完整的合约来生成 go 文件，只要有相应的 event 签名即可，例如下面的 .sol 文件生成
> 的 go 文件就有 NewNFT721BridgeFilterer 全局函数，返回的实例就有 ParseReceiveNFT、ParseReceiveNFT
> 这两个方法，可以用于解析 TransferNFT 和 ReceiveNFT 这两个 event 参数的 log
```solidity
pragma solidity ^0.8.0;
contract NFT721Bridge {
    event TransferNFT(uint64 indexed nonce, address token, uint256 tokenID, uint16 dstChainId, address sender, address recipient);
    event ReceiveNFT(uint64 indexed nonce, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient);
}
```
