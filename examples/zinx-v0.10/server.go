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

func onConnStart(conn ziface.IConnection) {
	fmt.Println("User OnConnStart Hook")
	conn.SetProperty("Name", "Daniel")
}

func onConnStop(conn ziface.IConnection) {
	fmt.Println("User OnConnStop Hook")
	name, err := conn.GetProperty("Name")
	if err == nil {
		fmt.Println("Name", name)
	}

	conn.RemoveProperty("Name")
	name, err = conn.GetProperty("Name")
	if err == nil {
		fmt.Println("Name", name)
	} else {
		fmt.Println("NoName")
	}
}

func main() {
	server := znet.NewServer("[zinx-v0.10]")
	// 注册Hook函数
	server.SetOnConnStart(onConnStart)
	server.SetOnConnStop(onConnStop)
	// 注册路由
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	server.Serve()
}
