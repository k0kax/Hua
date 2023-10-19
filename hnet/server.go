package hnet

import (
	"Hua/hiface"
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
}

// 启动服务器
func (s *Server) Start() {
	//0.日志记录
	fmt.Println("[Start] Server  Listenner at IP:%s,port:%d is starting!!!", s.IP, s.Port)
	go func() {
		//1.获取一个TCP的Addr句柄 更像是，重新设置以下本地监听的信息
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
		}

		//2.监听服务器的地址 获取监听 根据上面给的信息在本地监听
		listernner, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("Listern", s.IPversion, "err", err)
			return
		}
		fmt.Println("Start Hua Server succ name:", s.Name, "succ,Listerning....")

		//3.阻塞的等待客户端连接，处理客户端业务（读写）
		for {
			//如果有客户连接进来，阻塞会返回
			conn, err := listernner.AcceptTCP() //套接字句柄 进行监听返回一个conn连接
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//客户端已经建立连接，做一些业务，此处做一个最大512字节回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf) //cnt 是 Read 方法的返回值，表示成功读取的字节数。
					if err != nil {
						fmt.Println("recv buf err:", err)
						continue
					}

					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err")
						continue
					}
				}
			}()
		}
	}()
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()
}

// 停止服务器
func (s *Server) Stop() {

}

// 初始化服务器模块
func NewServer(name string) hiface.IServer {

	s := &Server{
		Name:      name,
		IPversion: "TCP4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
