package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-demo/ziface"
)

type Connection struct {
	// 当前连接
	Conn *net.TCPConn
	// 当前连接ID
	ConnId uint32
	// 当前连接状态
	IsClosed bool
	// 当前连接的业务处理逻辑的MsgHandler
	MsgHandler ziface.IMsgHandler
	// 告知当前连接已经退出/停止的 channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		MsgHandler: msgHandler,
		IsClosed:   false,
		ExitChan:   make(chan bool, 1),
	}
}

func (conn *Connection) StartReader() {
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

		// 根据MsgId找到对应的处理Handler并处理数据
		go conn.MsgHandler.DoMsgHandle(&request)
	}
}

func (conn *Connection) Start() {
	fmt.Println("Connection start... connId =", conn.ConnId)

	// 处理数据
	go conn.StartReader()
}

func (conn *Connection) Stop() {
	fmt.Println("Connection stop... connId =", conn.ConnId)

	if conn.IsClosed {
		return
	}

	conn.IsClosed = true
	conn.Conn.Close()
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
	if conn.IsClosed {
		return errors.New("Connection closed")
	}

	// 封包
	dp := NewDataPack()
	binaryData, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("[Server]Pack message error, msgId=", msgId)
		return err
	}

	// 发包
	if _, err = conn.GetTcpConnection().Write(binaryData); err != nil {
		fmt.Println("[Server]Send message error, msgId=", msgId)
		return err
	}

	return nil
}
