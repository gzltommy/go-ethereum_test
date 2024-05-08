package main

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"
	"net"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
)

// 使用速率受限的 RPC 提供程序
func main() {
	cluster := rpc.MainNetBeta
	client := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		cluster.RPC,
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))

	out, err := client.GetVersion(
		context.TODO(),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

// 用于向 RPC 提供程序进行身份验证的自定义标头
func main2() {
	cluster := rpc.MainNetBeta
	client := rpc.NewWithHeaders(
		cluster.RPC,
		map[string]string{
			"x-api-key": "...",
		},
	)

	out, err := client.GetVersion(
		context.TODO(),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}

// 超时和自定义 HTTP 客户端
func main3() {
	cluster := rpc.MainNetBeta
	rpcClient := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(
		cluster.RPC,
		rate.Every(time.Second), // time frame
		5,                       // limit of requests per time frame
	))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	account := solana.PublicKey{}

	acc, err := rpcClient.GetAccountInfoWithOpts(
		ctx,
		account,
		&rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		panic(err)
	}

	spew.Dump(acc)
}

// 使用自定义 HTTP 客户端初始 rpc.NewWithCustomRPCClient 化 RPC 客户端
func NewHTTPTransport(
	timeout time.Duration,
	maxIdleConnsPerHost int,
	keepAlive time.Duration,
) *http.Transport {
	return &http.Transport{
		IdleConnTimeout:     timeout,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		Proxy:               http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
		}).Dial,
	}
}

// NewHTTP returns a new Client from the provided config.
func NewHTTP(
	timeout time.Duration,
	maxIdleConnsPerHost int,
	keepAlive time.Duration,
) *http.Client {
	tr := NewHTTPTransport(
		timeout,
		maxIdleConnsPerHost,
		keepAlive,
	)

	return &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
}

// NewRPC creates a new Solana JSON RPC client.
func NewRPC(rpcEndpoint string) *rpc.Client {
	var (
		defaultMaxIdleConnsPerHost = 10
		defaultTimeout             = 25 * time.Second
		defaultKeepAlive           = 180 * time.Second
	)
	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: NewHTTP(
			defaultTimeout,
			defaultMaxIdleConnsPerHost,
			defaultKeepAlive,
		),
	}
	rpcClient := jsonrpc.NewClientWithOpts(rpcEndpoint, opts)
	return rpc.NewWithCustomRPCClient(rpcClient)
}
