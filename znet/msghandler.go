package znet

import (
	"fmt"
	"zinx-demo/ziface"
)

type MsgHandler struct {
	// Apis 存储消息Id -> 处理路由
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// DoMsgHandle 根据MsgId获取对应的消息处理路由并处理
func (mh *MsgHandler) DoMsgHandle(request ziface.IRequest) {
	// 检查消息对应的路由是否已存在
	msgHandler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("Cannot find router for msgId=", request.GetMsgId())
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
		fmt.Println("Router has already existed")
		return
	}

	// 添加路由
	mh.Apis[msgId] = router
}
