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

	//等待连接被动退出的channel管道
	ExitChan chan bool

	//该链接处理的方法
	Router hiface.IRouter
}

// 初始化链接的方法
func NewConnection(conn *net.TCPConn, connID uint32, router hiface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 读业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running....")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit,remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的业务到buf中，最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf) //cnt用不上了
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		//不匹配问题 使用
		//c.Router.Handle(req)
		//执行注册的路由方法
		go func(request hiface.IRequest) {
			//从路由中找到注册绑定的Conn对应的router调用
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// 启动链接
func (c *Connection) Start() {
	fmt.Println("Conn Start()..ConnID=", c.ConnID)
	//启动从当前链接的读数据业务
	go c.StartReader()
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
