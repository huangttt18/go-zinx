package znet

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"zinx-demo/ziface"
)

// ConnManager 连接管理器
type ConnManager struct {
	// 所有已建立的连接
	connections map[uint32]ziface.IConnection
	// 读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

// Add 添加连接
func (cm *ConnManager) Add(conn ziface.IConnection) {
	// 加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 检查连接是否已经存在
	if _, ok := cm.connections[conn.GetConnId()]; ok {
		fmt.Printf("[ConnManager]Connection[%d] already existed\n", conn.GetConnId())
		return
	}

	// 不存在则新建连接
	cm.connections[conn.GetConnId()] = conn
	fmt.Printf("[ConnManager]Connection[%d] added, current connection num = [%d]\n", conn.GetConnId(), len(cm.connections))
}

// Remove 删除连接
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	// 加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 检查连接是否存在
	if _, ok := cm.connections[conn.GetConnId()]; !ok {
		fmt.Printf("[ConnManager]Connection[%d] not exists\n", conn.GetConnId())
		return
	}

	// 连接存在则清理
	delete(cm.connections, conn.GetConnId())
	fmt.Printf("[ConnManager]Connection[%d] removed, current connection num = [%d]\n", conn.GetConnId(), len(cm.connections))
}

// Get 获取连接
func (cm *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	// 加锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	conn, ok := cm.connections[connId]
	if !ok {
		fmt.Printf("[ConnManager]Connection[%d] not exists\n", connId)
		return nil, errors.New("cannot find connection[" + strconv.Itoa(int(connId)) + "]")
	}

	return conn, nil
}

// Len 获取连接数量
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

// Clear 清理全部连接
func (cm *ConnManager) Clear() {
	// 加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connId, conn := range cm.connections {
		// 停止连接
		conn.Stop()
		// 删除连接
		delete(cm.connections, connId)
	}

	fmt.Println("[ConnManager]Clear all connections successfully!")
}
