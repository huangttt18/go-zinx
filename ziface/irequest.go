package ziface

type IRequest interface {
	// GetConnection 获取连接
	GetConnection() IConnection
	// GetData 获取数据
	GetData() []byte
	// GetMsgId 获取消息ID
	GetMsgId() uint32
}
