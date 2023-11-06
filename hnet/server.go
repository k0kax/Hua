package hnet

import (
	"Hua/hiface"
	"Hua/utils"
	"fmt"
	"net"
)

// iserver的接口实现，定义一个Server的服务模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPversion string
	//服务器绑定的ip
	IP string
	//服务器监听的端口
	Port int

	//当前的server添加router，server注册的链接对应的处理业务
	Router hiface.IRouter
}

// 启动服务器
func (s *Server) Start() {
	//0.日志记录
	fmt.Printf("[Hua] Server Name:%s,Listenner at IP:%s,port:%d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Start] Server  Listenner at IP:%s,port:%d is starting!!!\n", s.IP, s.Port)
	go func() {
		//1.获取一个TCP的Addr句柄 更像是，重新设置以下本地监听的信息
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}

		//2.监听服务器的地址 获取监听 根据上面给的信息在本地监听
		listener, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("Listern", s.IPversion, "err", err)
			return
		}
		fmt.Println("Start Hua Server succ name:", s.Name, "succ,Listerning....")

		var cid uint32
		cid = 0 //链接id

		//3.阻塞的等待客户端连接，处理客户端业务（读写）
		for {
			//如果有客户连接进来，阻塞会返回
			conn, err := listener.AcceptTCP() //套接字句柄 进行监听返回一个conn连接
			if err != nil {
				fmt.Println("Accept err", err)
				continue //继续执行
			}

			delConn := NewConnection(conn, cid, s.Router)
			cid++

			//启动当前的链接业务处理
			go delConn.Start()
		}
	}()
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()

	//todo 启动服务器的其他业务
	//阻塞状态
	select {}
}

// 停止服务器
func (s *Server) Stop() {
	//todo 停止服务器的其他业务
}

func (s *Server) AddRouter(router hiface.IRouter) {
	s.Router = router
	fmt.Println("Add router Succ!!")
}

// 初始化服务器模块
func NewServer(name string) hiface.IServer {

	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPversion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
