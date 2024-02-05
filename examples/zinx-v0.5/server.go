package main

import (
	"fmt"
	"zinx-demo/ziface"
	"zinx-demo/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Receive message from client, msgId=", request.GetMsgId(), "data=", string(request.GetData()))
	request.GetConnection().SendMsg(1, []byte("Copy!"))
}

func main() {
	server := znet.NewServer("[zinx-v0.5]")
	// 注册路由
	server.AddRouter(&PingRouter{})
	server.Serve()
}
