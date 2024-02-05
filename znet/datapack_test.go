package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	// 构建服务端
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Listen to 127.0.0.1:8999 error", err)
		return
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			return
		}

		go func(conn net.Conn) {
			// 服务端拆包
			dp := NewDataPack()
			for {
				// 拆头
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("Server read head error", err)
					return
				}

				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("Server unpack head error", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					// 拆Data
					msg := msgHead.(*Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("Server unpack data error", err)
						return
					}

					// 消息读完了，输出一下
					fmt.Println("Server unpack message, MsgId", msg.GetMsgId(), " DataLen=", msg.GetDataLen(), " Data=", string(msg.GetData()))
				}
			}
		}(conn)
	}()

	// 构建客户端
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:8999")
		if err != nil {
			fmt.Print("Connect to server error", err)
			return
		}

		// 客户端封包
		dp := NewDataPack()
		msg1 := &Message{
			Id:      1,
			DataLen: 4,
			Data:    []byte{'z', 'i', 'n', 'x'},
		}

		msg1Bytes, err := dp.Pack(msg1)
		if err != nil {
			fmt.Println("Pack msg1 error", err)
			return
		}

		msg2 := &Message{
			Id:      2,
			DataLen: 10,
			Data:    []byte{'h', 'e', 'l', 'l', 'o', ',', 'z', 'i', 'n', 'x'},
		}
		msg2Bytes, err := dp.Pack(msg2)
		if err != nil {
			fmt.Println("Pack msg2 error", err)
			return
		}

		sendBytes := append(msg1Bytes, msg2Bytes...)

		_, err = conn.Write(sendBytes)
		if err != nil {
			fmt.Println("Write msg error", err)
		}
	}()

	select {}
}
