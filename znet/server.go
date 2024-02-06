package znet

import (
	"fmt"
	"net"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

type Server struct {
	Name        string
	IpVersion   string
	Ip          string
	Port        int
	MsgHandler  ziface.IMsgHandler
	ConnManager ziface.IConnManager
	OnConnStart func(ziface.IConnection)
	OnConnStop  func(ziface.IConnection)
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:        utils.GlobalObject.Name,
		IpVersion:   "tcp4",
		Ip:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}
}

// Start 启动服务器，创建连接
func (s *Server) Start() {
	fmt.Printf("[Zinx]ServerName: %s\n", utils.GlobalObject.Name)
	fmt.Printf("[Zinx]Host: %s\n", utils.GlobalObject.Host)
	fmt.Printf("[Zinx]Port: %d\n", utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx]Version: %s\n", utils.GlobalObject.Version)
	fmt.Printf("[Zinx]MaxConn: %d\n", utils.GlobalObject.MaxConn)
	fmt.Printf("[Zinx]MaxPackageSize: %d\n\n", utils.GlobalObject.MaxPacketSize)

	fmt.Printf("[Server]Listenner on %s:%d is starting...\n", s.Ip, s.Port)

	go func() {
		// 初始化WorkerPool
		s.MsgHandler.InitWorkerPool()

		// 创建addr、绑定ip、端口
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("[Server]Resolve tcp addr error ", err)
			return
		}

		// 监听
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("[Server]Listen ", s.IpVersion, " error", err)
			return
		}

		fmt.Println("[Server]Zinx server", s.Name, "started, listenning...")
		cid := uint32(0)
		// 接受请求
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[Server]Accept error", err)
				continue
			}

			// 检查当前连接数是否超过最大连接数
			if s.ConnManager.Len() >= int(utils.GlobalObject.MaxConn) {
				// 关闭当前连接
				conn.Close()
				fmt.Println("[Server]Build connection failed, too many connections")
				continue
			}

			// 处理请求并返回响应
			connection := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			go connection.Start()
		}
	}()
}

// Stop 停止服务器，释放资源
func (s *Server) Stop() {
	fmt.Println("[Server]Stop...")
	// 清理连接
	s.ConnManager.Clear()
}

// Serve 运行服务器，初始化
func (s *Server) Serve() {
	s.Start()

	// TODO: 初始化资源等

	// 阻塞
	select {}
}

// AddRouter 添加路由
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

// GetConnManger 获取连接管理器
func (s *Server) GetConnManger() ziface.IConnManager {
	return s.ConnManager
}

// SetOnConnStart 注册建立连接后执行的Hook函数
func (s *Server) SetOnConnStart(hook func(conn ziface.IConnection)) {
	s.OnConnStart = hook
}

// SetOnConnStop 注册销毁连接后执行的Hook函数
func (s *Server) SetOnConnStop(hook func(conn ziface.IConnection)) {
	s.OnConnStop = hook
}

// CallOnConnStart 调用Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart == nil {
		return
	}

	fmt.Println("[Server]Execute OnConnStart Hook")
	s.OnConnStart(conn)
}

// CallOnConnStop 调用Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop == nil {
		return
	}

	fmt.Println("[Server]Execute OnConnStop Hook")
	s.OnConnStop(conn)
}
