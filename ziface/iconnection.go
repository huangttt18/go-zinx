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
	// Send 发送数据
	SendMsg(msgId uint32, data []byte) error
}
