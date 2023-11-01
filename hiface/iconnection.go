package hiface

import "net"

/*
	连接的抽象层
*/

type IConnection interface {
	//启动链接
	Start()

	//停止连接
	Stop()

	//获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	//获取当前连接的ID
	GetConnID() uint32

	//获取远程客户端的TCP状态、IP、端口
	RemoteAddr() net.Addr

	//发送数据
	Send(data []byte) error
}

// 定义一个处理链接业务的方法 链接/数据内容/数据长度
type HandleFunc func(*net.TCPConn, []byte, int) error
