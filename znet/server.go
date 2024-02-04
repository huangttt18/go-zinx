package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx-demo/ziface"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}
}

// 启动服务器，创建连接
func (s *Server) Start() {
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
			connection := NewConnection(conn, cid, Callback)
			cid++
			go connection.Start()
		}
	}()
}

func Callback(conn *net.TCPConn, data []byte, bytes int) error {
	if _, err := conn.Write(data); err != nil {
		fmt.Println("Callback error", err)
		return errors.New("Callback Error")
	}

	return nil
}

// 停止服务器，释放资源
func (s *Server) Stop() {

}

// 运行服务器，初始化
func (s *Server) Serve() {
	s.Start()

	// TODO: 初始化资源等

	// 阻塞
	select {}
}
