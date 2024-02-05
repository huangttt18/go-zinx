package ziface

// IRouter 路由，用于处理请求
type IRouter interface {
	// PreHandle 处理请求之前需要做的工作
	PreHandle(request IRequest)
	// Handle 处理请求
	Handle(request IRequest)
	// PostHandle 处理请求之后需要做的工作
	PostHandle(request IRequest)
}
