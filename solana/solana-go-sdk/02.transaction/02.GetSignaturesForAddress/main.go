package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"log"
)

func main() {
	// create a RPC client
	//c := client.NewClient(rpc.DevnetRPCEndpoint)
	c := client.NewClient("https://rpc.ankr.com/solana_devnet")
	sigs, err := c.GetSignaturesForAddressWithConfig(context.Background(), "FFP7XhWSL4SKWdS6SrNhA6Ef9hVxUqdgRH3g1MErViRR", client.GetSignaturesForAddressConfig{
		Limit:      10,
		Before:     "",
		Until:      "",
		Commitment: "",
	})

	//sigs, err := c.GetSignaturesForAddress(context.Background(), "FFP7XhWSL4SKWdS6SrNhA6Ef9hVxUqdgRH3g1MErViRR")
	if err != nil {
		log.Fatalf("GetSignaturesForAddressWithConfig fail.err:%v", err)
	}

	buf, _ := json.Marshal(sigs)
	fmt.Println(string(buf))
}
