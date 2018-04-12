package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thingful/decodestore/pkg/rpc/datastore"
)

func main() {
	client := datastore.NewEncryptedStoreProtobufClient("http://localhost:8080", &http.Client{})

	_, err := client.WriteData(context.Background(), &datastore.WriteRequest{
		PublicKey: "foo",
		Data:      []byte("hello world"),
	})

	if err != nil {
		panic(err)
	}

	resp, err := client.ReadData(context.Background(), &datastore.ReadRequest{
		PublicKey: "foo",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Read data: %s\n", resp.Events)
}
