package hnet

import (
	"Hua/hiface"
	"fmt"
	"net"
)

/*
	具体实现层
*/

// 链接模块
type Connection struct {

	//socketTCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnID uint32

	//当前连接的状态（是否已经关闭）
	isClosed bool

	//与当前连接绑定的处理业务和方法
	handleAPI hiface.HandleFunc

	//等待连接被动退出的channel管道
	ExitChan chan bool
}

// 初始化链接的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api hiface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 读业务
func (c *Connection) StartRead() {
	fmt.Println("Reader Goroutine is running....")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的业务到buf中，最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		//调用当前业务绑定的API
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID=", c.ConnID, "handle is error", err)
			break
		}

	}
}

// 启动链接
func (c *Connection) Start() {
	fmt.Println("Conn Start()..ConnID=", c.ConnID)
	//启动从当前链接的读数据业务
	go c.StartRead()
	//TODO 启动从当前链接写数据的业务

}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()..ConnID=", c.ConnID)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭socket链接
	c.Conn.Close()

	//回收资源
	close(c.ExitChan)
}

// 获取当前连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态、IP、端口
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
