package znet

type Message struct {
	// 消息ID
	Id uint32
	// 消息长度
	DataLen uint32
	// 消息内容
	Data []byte
}

func NewMessage(msgId uint32, data []byte) *Message {
	return &Message{
		Id:      msgId,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// GetMsgId 获取消息ID
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// GetDataLen 获取消息长度
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// GetData 获取消息
func (m *Message) GetData() []byte {
	return m.Data
}

// SetMsgId 设置消息ID
func (m *Message) SetMsgId(msgId uint32) {
	m.Id = msgId
}

// SetDataLen 设置消息长度
func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

// SetData 设置消息
func (m *Message) SetData(data []byte) {
	m.Data = data
}
