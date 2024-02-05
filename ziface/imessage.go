package ziface

type IMessage interface {
	// GetMsgId 获取消息ID
	GetMsgId() uint32
	// GetDataLen 获取消息长度
	GetDataLen() uint32
	// GetData 获取消息
	GetData() []byte
	// SetMsgId 设置消息ID
	SetMsgId(uint32)
	// SetDataLen 设置消息长度
	SetDataLen(uint32)
	// SetData 设置消息
	SetData([]byte)
}
