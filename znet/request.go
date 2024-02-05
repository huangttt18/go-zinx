package znet

import "zinx-demo/ziface"

// Request 封装了IConnection和请求数据
type Request struct {
	// 连接
	Conn ziface.IConnection
	// 数据
	Msg ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.Msg.GetMsgId()
}
