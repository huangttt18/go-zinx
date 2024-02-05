package ziface

type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 关闭服务器
	Stop()
	// Serve 运行服务器，初始化资源
	Serve()
	// AddRouter 增加路由
	AddRouter(uint32, IRouter)
	// GetConnManger 获取连接管理器
	GetConnManger() IConnManager
	// SetOnConnStart 注册建立连接后执行的Hook函数
	SetOnConnStart(func(IConnection))
	// SetOnConnStop 注册销毁连接后执行的Hook函数
	SetOnConnStop(func(IConnection))
	// CallOnConnStart 调用Hook函数
	CallOnConnStart(IConnection)
	// CallOnConnStop 调用Hook函数
	CallOnConnStop(IConnection)
}
