package znet

import (
	"fmt"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

type MsgHandler struct {
	// Apis 存储消息Id -> 处理路由
	Apis map[uint32]ziface.IRouter
	// 消息队列
	TaskQueue []chan ziface.IRequest
	// 工作池的worker数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

// DoMsgHandle 根据MsgId获取对应的消息处理路由并处理
func (mh *MsgHandler) DoMsgHandle(request ziface.IRequest) {
	// 检查消息对应的路由是否已存在
	msgHandler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("[MsgHandler]Cannot find router for msgId=", request.GetMsgId())
		return
	}

	// 处理消息
	msgHandler.PreHandle(request)
	msgHandler.Handle(request)
	msgHandler.PostHandle(request)
}

// AddRouter 添加路由到Handler中
func (mh *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	// 检查路由是否已存在
	if _, ok := mh.Apis[msgId]; ok {
		fmt.Println("[MsgHandler]Router has already existed")
		return
	}

	// 添加路由
	mh.Apis[msgId] = router
}

// InitWorkerPool 初始化workerPool
func (mh *MsgHandler) InitWorkerPool() {
	fmt.Println("[MsgHandler]Initializing worker pool...")
	// 如果WorkerPoolSize <= 0，则表示当前不开启workerPool
	if mh.WorkerPoolSize <= 0 {
		fmt.Println("[MsgHandler]WorkerPoolSize smaller than 0, do not use workerPool")
		return
	}

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 创建消息队列，长度为MaxWorkerTaskSize
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskSize)
		// 启动worker
		go mh.execute(uint32(i), mh.TaskQueue[i])
	}
}

// execute 处理业务逻辑
func (mh *MsgHandler) execute(workerId uint32, taskQueue chan ziface.IRequest) {
	fmt.Printf("[MsgHandler]Worker[%d] start working\n", workerId)
	// 阻塞等待消息的到来
	for {
		request := <-taskQueue
		// 消息到来之后就执行处理逻辑
		mh.DoMsgHandle(request)
	}
}

// SubmitTask 提交任务到worker对应的taskQueue中，目前采用平均分配的机制
func (mh *MsgHandler) SubmitTask(request ziface.IRequest) {
	// 找到对应的workerId: connId % poolSize
	workerId := request.GetConnection().GetConnId() % mh.WorkerPoolSize
	fmt.Printf("[MsgHandler]Task distributed to worker[%d]\n", workerId)

	// 将请求发送到worker对应的channel
	mh.TaskQueue[workerId] <- request
}
