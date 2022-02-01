package ziface

type IDatapack interface {
	GetHeadLen() uint32
	Pack(message IMessage) []byte
	Unpack([]byte) IMessage
}