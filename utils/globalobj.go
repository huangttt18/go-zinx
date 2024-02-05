package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"zinx-demo/ziface"
)

// GlobalObj 全局配置文件
type GlobalObj struct {
	// 全局Server对象
	TcpServer ziface.IServer
	// 服务器名
	Name string
	// 服务器ip/域名
	Host string
	// 服务器端口
	TcpPort int

	// zinx版本号
	Version string
	// 最大连接数
	MaxConn uint32
	// 每次请求最大处理字节数
	MaxPacketSize uint32
	// 工作池worker数量
	WorkerPoolSize uint32
	// 每个worker最大等待处理的任务数量
	MaxWorkerTaskSize uint32
}

// Reload 从配置文件中读取配置
func (globalObj *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("Load conf/zinx.json error", err)
		panic(err)
	}

	err = json.Unmarshal(data, globalObj)
	if err != nil {
		fmt.Println("Unmarshal conf/zinx.json error", err)
		panic(err)
	}
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:              "ZinxServerApp",
		Host:              "0.0.0.0",
		TcpPort:           8999,
		Version:           "zinx-v0.8",
		MaxConn:           1024,
		MaxPacketSize:     4096,
		WorkerPoolSize:    0,
		MaxWorkerTaskSize: 1024,
	}

	GlobalObject.Reload()
}
