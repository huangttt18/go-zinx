package ziface

import "net"

type IConnection interface {
	// Start 启动连接
	Start()
	// Stop 关闭连接
	Stop()
	// GetTcpConnection 获取当前连接对应的connection socket
	GetTcpConnection() *net.TCPConn
	// GetConnId 获取当前连接ID
	GetConnId() uint32
	// RemoteAddr 获取客户端的地址和端口
	RemoteAddr() net.Addr
	// SendMsg 发送数据
	SendMsg(uint32, []byte) error
	// SetProperty 设置连接属性
	SetProperty(string, interface{})
	// GetProperty 获取连接属性
	GetProperty(string) (interface{}, error)
	// RemoveProperty 移除连接属性
	RemoveProperty(string) error
}
