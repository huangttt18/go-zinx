package ziface

// 路由，用于处理请求
type IRouter interface {
	// 处理请求之前需要做的工作
	PreHandle(request IRequest)
	// 处理请求
	Handle(request IRequest)
	// 处理请求之后需要做的工作
	PostHandle(request IRequest)
}
