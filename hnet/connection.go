package hnet

import (
	"Hua/hiface"
	"errors"
	"fmt"
	"io"
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

		//创建一个拆包解包的对象
		dp := NewDataPack()

		//读取客户端的Msg Head 二进制流 8 字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}
		//拆包 得到msgId msgDataLen 放在msg消息里
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		//根据dataLen 再次读取Data,放在msg.Data里
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}

		msg.SetData(data)
		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
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

// 连接提供一个发包的方法：将发送的信息进行打包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when Send msg")
	}

	//将data进行封包 MsgDataLen/MsgId/Data
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id=", msgId)
		return errors.New("Pack error mag")
	}

	//将数据发送给客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write msg id:", msgId, "error:", err)
		return errors.New("conn Write error")
	}

	return nil
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
