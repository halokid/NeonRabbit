package main

import (
	"context"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
)

func C1() error {

	// create the client
	ctx := context.Background()
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// invoke a method called EchoMethod on another dapr enabled service
	content := &dapr.DataContent{
		ContentType: "text/plain",
		Data:        []byte("hellow"),
	}
	resp, err := client.InvokeMethodWithContent(ctx, "neon_broker", "ping", "post", content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("service method invoked response -->>> %s\n", string(resp))
	return nil
}
