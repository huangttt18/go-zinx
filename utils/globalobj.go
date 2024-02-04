package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"zinx-demo/ziface"
)

// 全局配置文件
type GlobalObj struct {
	// 全局Server对象
	TcpServer ziface.IServer
	// 服务器名
	Name string
	// 服务器ip/域名
	Host string
	// 服务器端口
	TcpPort int

	// Zinx版本号
	Version string
	// 最大连接数
	MaxConn int
	// 每次请求最大处理字节数
	MaxPackageSize int
}

// 从配置文件中读取配置
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
		Name:           "ZinxServerApp",
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Version:        "zinx-v0.4",
		MaxConn:        1024,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
