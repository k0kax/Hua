package hnet

import (
	"Hua/hiface"
	"Hua/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

// 封包，拆包的具体模块（并没有实现，只是提供一个承载）
type DataPack struct{}

// 拆包封包的初始化实例方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的head的方法
func (dp *DataPack) GetHeadLen() uint32 {
	//datalen uint32(4字节)+ID uint32(4字节)
	return 8
}

// 封包的方法 dataLen ID data （tlv ---> 二进制）
func (dp *DataPack) Pack(msg hiface.IMessage) ([]byte, error) {

	//创建一个存放字节流的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将dataLen写入dataBuff(二进制写法),小端
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//将dataLen写入dataBuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//将dataLen写入dataBuff
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// 拆包的方法 (二进制---> tlv)
// 将包的Head信息读出来，再将head的data的长度（dataLen）,再读一次
func (dp *DataPack) Unpack(binaryData []byte) (hiface.IMessage, error) {
	//创建一个从输入二进制的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息的带dataLen，MsgId
	msg := &Message{} //它承载

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读MsgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen是否已经超过允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("Too Large msg data recv!!")
	}

	return msg, nil

}
