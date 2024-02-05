package main

import (
	"fmt"
	"io"
	"net"
	"time"
	zinx "zinx-demo/znet"
)

func main() {
	fmt.Println("[Client]Start client1...")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("[Client]Connect to remote server error", err)
		return
	}

	for {
		msg := zinx.NewMessage(1, []byte("Hello, Zinx-v0.9"))
		dp := zinx.NewDataPack()
		// 封包
		binaryData, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("[Client]Pack data error", err)
			break
		}
		// 发包
		_, err = conn.Write(binaryData)
		if err != nil {
			fmt.Println("[Client]Send message error", err)
			break
		}

		// 收包、拆包
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("[Client]Recv head error", err)
			break
		}

		msgRecv, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("[Client]Unpack head error", err)
			break
		}

		if msgRecv.GetDataLen() > 0 {
			msgData := make([]byte, msgRecv.GetDataLen())
			_, err = io.ReadFull(conn, msgData)
			if err != nil {
				fmt.Println("[Client]Read msgData error", err)
				break
			}

			msgRecv.SetData(msgData)
		}

		fmt.Println("[Client]Receive message from server, msgId=", msgRecv.GetMsgId(), "msgLen=", msgRecv.GetDataLen(), "msgData=", string(msgRecv.GetData()))

		time.Sleep(2 * time.Second)
	}
}
