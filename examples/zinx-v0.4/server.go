package main

import (
	"fmt"
	"zinx-demo/ziface"
	"zinx-demo/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PreHandle...")
	request.GetConnection().GetTcpConnection().Write([]byte("Before ping...\n"))
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle...")
	request.GetConnection().GetTcpConnection().Write([]byte("Ping...\n"))
}

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PostHandle...")
	request.GetConnection().GetTcpConnection().Write([]byte("After ping...\n"))
}

func main() {
	server := znet.NewServer("[zinx-v0.4]")
	// 注册路由
	server.AddRouter(&PingRouter{})
	server.Serve()
}
