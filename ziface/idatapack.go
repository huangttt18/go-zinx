package ziface

// IDataPack 数据的拆包、封包
type IDataPack interface {
	// GetHeadLen 获取包头的大小
	GetHeadLen() uint32
	// Pack 封包，将请求数据转换为二进制数据
	Pack(IMessage) ([]byte, error)
	// Unpack 拆包，将二进制数据转换为请求数据
	Unpack([]byte) (IMessage, error)
}
