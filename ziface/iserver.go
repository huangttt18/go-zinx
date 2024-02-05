package ziface

type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 关闭服务器
	Stop()
	// Serve 运行服务器，初始化资源
	Serve()
	// AddRouter 增加路由
	AddRouter(router IRouter)
}
