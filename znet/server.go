package znet

import (
	"fmt"
	"net"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
	Router    ziface.IRouter
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      utils.GlobalObject.Name,
		IpVersion: "tcp4",
		Ip:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
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

			// 处理请求并返回响应
			connection := NewConnection(conn, cid, s.Router)
			cid++
			go connection.Start()
		}
	}()
}

// Stop 停止服务器，释放资源
func (s *Server) Stop() {

}

// Serve 运行服务器，初始化
func (s *Server) Serve() {
	s.Start()

	// TODO: 初始化资源等

	// 阻塞
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}
