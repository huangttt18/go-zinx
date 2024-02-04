package main

import (
	"zinx-demo/znet"
)

func main() {
	server := znet.NewServer("[zinx-v0.1]")
	server.Serve()
}
