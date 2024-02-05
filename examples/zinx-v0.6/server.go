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
	fmt.Println("[PingRouter]Receive message from client, msgId=", request.GetMsgId(), "data=", string(request.GetData()))
	request.GetConnection().SendMsg(0, []byte("Ping!"))
}

type HelloRouter struct {
	znet.BaseRouter
}

func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("[HelloRouter]Receive message from client, msgId=", request.GetMsgId(), "data=", string(request.GetData()))
	request.GetConnection().SendMsg(1, []byte("Hello!"))
}

func main() {
	server := znet.NewServer("[zinx-v0.6]")
	// 注册路由
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	server.Serve()
}
