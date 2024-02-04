package znet

import "zinx-demo/ziface"

// 封装了IConnection和请求数据
type Request struct {
	// 连接
	conn ziface.IConnection
	// 数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
