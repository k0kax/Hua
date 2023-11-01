package main

import (
	"Hua/hiface"
	"Hua/hnet"
	"fmt"
)

/*
基于Hua开发的 服务器端应用程序
*/

// ping test 自定义路由，继承baseRouter
type PingRouter struct {
	hnet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request hiface.IRequest) {
	fmt.Println("Call Router PreHandle....")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping ..."))
	if err != nil {
		fmt.Println("Call back before ping error", err)
	}
}

// Test Handle
func (this *PingRouter) Handle(request hiface.IRequest) {
	fmt.Println("Call Router Handle....")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ... ping....ping"))
	if err != nil {
		fmt.Println("Call back  ping...ping...ping.. error", err)
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request hiface.IRequest) {
	fmt.Println("Call Router PostHandle....")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte(" After ping"))
	if err != nil {
		fmt.Println("Call back   After ping error", err)
	}
}

func main() {
	//1.新建server句柄，使用Hua的api
	s := hnet.NewServer("[Hua V0.3]")

	//2.给当前的框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//3.启动server
	s.Serve()
}
