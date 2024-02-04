package znet

import "zinx-demo/ziface"

// 封装了IConnection和请求数据
type Request struct {
	// 连接
	Conn ziface.IConnection
	// 数据
	Data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}
