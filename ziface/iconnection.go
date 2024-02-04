package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 关闭连接
	Stop()
	// 获取当前连接对应的connection socket
	GetTcpConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnId() uint32
	// 获取客户端的地址和端口
	RemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

// 当前连接处理业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
