# 第一步是安装 Solidity编译器 (solc)
Solc 在 Ubuntu 上有 snapcraft 包。
> sudo add-apt-repository ppa:ethereum/ethereum
>
> sudo apt-get update
> 
> sudo apt-get install solc

## 指定安装版本
如果需要指定安装的 solc 的版本的话那么还需要安装 brew
>sudo apt install linuxbrew-wrapper

安装完毕后可以
> brew install solidity@0.5.0

来安装指定版本。如果需要 0.4.24 版本可以
> brew install solidity@0.4.24


其他的平台或者从源码编译的教程请查阅官方 solidity 文档 install guide（https://solidity.readthedocs.io/en/latest/installing-solidity.html#building-from-source）.


# 安装 abigen 的工具（从 solidity 智能合约生成 ABI）

假设您已经在计算机上设置了 Go，只需运行以下命令即可安装 abigen 工具。
> go get -u github.com/ethereum/go-ethereum
> 
> cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v1.10.18
> 
> make
> 
> make devtools

执行完上述命令后，abigen 工具就被安装到了 $GOPATH/bin 目录下


## 从一个 solidity 文件生成 ABI（json 的，web 端可以用）
> sudo solc --abi Store.sol --pretty-json >Store_sol_Store.abi

它会将其写入名为“Store_sol_Store.abi”的文件中，删除文件开头的无用数据

## 用 abigen 将 ABI 转换为我们可以导入的 Go 文件
> abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go


## 将 solidity 智能合约编译为 EVM 字节码
为了从 Go 部署智能合约，我们还需要将 solidity 智能合约编译为 EVM 字节码。 EVM 字节码将在事务的数据字段中发送。 在 Go 文件上生成部署方法需要bin文件。
> sudo solc --bin Store.sol >Store_sol_Store.bin

它会将其写入名为“Store_sol_Store.bin”的文件中，删除文件开头的无用数据

## 编译 Go 合约文件
现在我们编译 Go 合约文件，其中包括 deploy 方法，因为我们包含了 bin 文件。
> abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go


## 最终
> sudo solc --abi --bin ERC20.sol  -o ./  --overwrite
> 
> abigen --bin=ERC20_sol.bin --abi=ERC20_sol.abi --pkg=erc20 --out=ERC20.go





