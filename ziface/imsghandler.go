package ziface

// IMsgHandler 消息处理器
type IMsgHandler interface {
	// DoMsgHandle 根据MsgId获取对应的消息处理路由并处理
	DoMsgHandle(IRequest)
	// AddRouter 添加路由到Handler中
	AddRouter(uint32, IRouter)
	// InitWorkerPool 初始化workerPool
	InitWorkerPool()
	// SubmitTask 提交任务到Worker
	SubmitTask(IRequest)
}
