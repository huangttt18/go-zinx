package ziface

// IConnManager 连接管理器
type IConnManager interface {
	// Add 添加连接
	Add(IConnection)
	// Remove 删除连接
	Remove(IConnection)
	// Get 获取连接
	Get(uint32) (IConnection, error)
	// Len 获取连接数量
	Len() int
	// Clear 清理全部连接
	Clear()
}
