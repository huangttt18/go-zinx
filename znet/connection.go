package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

type Connection struct {
	// 当前连接
	Conn *net.TCPConn
	// 当前连接ID
	ConnId uint32
	// 当前连接状态
	isClosed bool
	// 当前连接的业务处理逻辑的MsgHandler
	MsgHandler ziface.IMsgHandler
	// 告知当前连接已经退出/停止的 channel
	ExitChan chan bool
	// Reader和Writer之间通信的 channel
	// client -> server -> reader -> writer -> client
	msgChan chan []byte
}

func NewConnection(conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		MsgHandler: msgHandler,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
}

func (conn *Connection) StartReader() {
	fmt.Println("[Server]Start reader")
	fmt.Println("[Server]Read message, connId =", conn.ConnId)
	defer fmt.Println("[Server]Read message finished, connId =", conn.ConnId, " RemoteAddr =", conn.RemoteAddr().String())
	defer conn.Stop()

	for {
		// 拆包，将二进制数据拆包为Message，再传递给router进行处理
		dp := NewDataPack()
		// HeadData
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn.GetTcpConnection(), headData); err != nil {
			fmt.Println("[Server]Read head error", err)
			break
		}

		// 拆包为Message
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("[Server]Unpack head error", err)
			break
		}

		msgData := make([]byte, msg.GetDataLen())
		// 读取Data
		if msg.GetDataLen() > 0 {
			if _, err = io.ReadFull(conn.GetTcpConnection(), msgData); err != nil {
				fmt.Println("[Server]Unpack data error", err)
				break
			}
		}
		// 将读取到的data放入Message中
		msg.SetData(msgData)

		request := Request{
			Conn: conn,
			Msg:  msg,
		}

		// 如果开启了workerPool，则用workerPool来处理请求
		if utils.GlobalObject.WorkerPoolSize > 0 {
			// 将请求提交到workerPool中，由相应的worker来处理
			conn.MsgHandler.SubmitTask(&request)
		} else {
			// 未开启workerPool，则手动开Goroutine来处理请求
			go conn.MsgHandler.DoMsgHandle(&request)
		}
	}
}

// StartWriter 启动一个处理写业务逻辑的Goroutine
// 可以在写前、写后做一些额外的操作
func (conn *Connection) StartWriter() {
	fmt.Println("[Server]Start writer")
	fmt.Println("[Server]Write message, connId =", conn.ConnId)
	defer fmt.Println("[Server]Write message finished, connId =", conn.ConnId, " RemoteAddr =", conn.RemoteAddr().String())

	for {
		// 写业务逻辑
		select {
		case binaryData := <-conn.msgChan:
			if _, err := conn.GetTcpConnection().Write(binaryData); err != nil {
				fmt.Println("[Server]Write message error", err)
				return
			}
		case <-conn.ExitChan:
			return
		}
	}
}

func (conn *Connection) Start() {
	fmt.Println("[Server]Connection start... connId =", conn.ConnId)
	// 处理读业务
	go conn.StartReader()
	// 处理写业务
	go conn.StartWriter()
}

func (conn *Connection) Stop() {
	fmt.Println("[Server]Connection stop... connId =", conn.ConnId)

	if conn.isClosed {
		return
	}

	conn.isClosed = true
	conn.Conn.Close()
	conn.ExitChan <- true
	close(conn.msgChan)
	close(conn.ExitChan)
}

func (conn *Connection) GetTcpConnection() *net.TCPConn {
	return conn.Conn
}

func (conn *Connection) GetConnId() uint32 {
	return conn.ConnId
}

func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

// SendMsg 发包，将数据封包，再发送出去
func (conn *Connection) SendMsg(msgId uint32, data []byte) error {
	if conn.isClosed {
		return errors.New("Connection closed")
	}

	// 封包
	dp := NewDataPack()
	binaryData, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("[Server]Pack message error, msgId=", msgId)
		return err
	}

	// 发包, 将数据通过msgChan发送给Writer Goroutine
	conn.msgChan <- binaryData

	return nil
}
