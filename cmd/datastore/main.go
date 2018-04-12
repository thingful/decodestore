package main

import "github.com/thingful/datastore/pkg/server"

func main() {
	server := server.NewServer(":8080")

	server.Start()
}
