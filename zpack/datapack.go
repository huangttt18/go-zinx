package zpack

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-demo/utils"
	"zinx-demo/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头的大小
func (dp *DataPack) GetHeadLen() uint32 {
	// | DataLen 4 bytes | MsgId 4 bytes | data
	return 8
}

// Pack 封包，将二进制数据转换为请求数据
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 初始化缓冲区，数据将会读到这里
	dataBuff := bytes.NewBuffer([]byte{})

	// 读取MsgId
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 读取DataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 读取数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包，将请求数据转换为二进制数据
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	// 拆MsgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 拆DataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 判断数据长度是否超过当前可读最大数据长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.GetDataLen() > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("packet data too large")
	}

	return msg, nil
}
