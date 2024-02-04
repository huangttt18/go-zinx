package ziface

type IServer interface {
	// 启动服务器
	Start()
	// 关闭服务器
	Stop()
	// 运行服务器，初始化资源
	Serve()
	// 增加路由
	AddRouter(router IRouter)
}
