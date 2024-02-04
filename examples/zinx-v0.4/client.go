package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("[Client]Start client...")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("[Client]Connect to remote server error", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello, Zinx-v0.4"))
		if err != nil {
			fmt.Println("[Client]Send message error", err)
			return
		}

		buf := make([]byte, 512)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("[Client]Read message error", err)
			return
		}

		fmt.Printf("[Client]Receive message from server: %s\n", buf[:read])

		time.Sleep(5 * time.Second)
	}
}
