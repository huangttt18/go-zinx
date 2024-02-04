package ziface

type IRequest interface {
	// 获取连接
	GetConnection() IConnection
	// 获取数据
	GetData() []byte
}
