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

// Test Handle
func (this *PingRouter) Handle(request hiface.IRequest) {
	fmt.Println("Call Router Handle....")
	//1.先读取客户端数据，再回写
	fmt.Println("recv from client: msgID =", request.GetMsgID(), ",data =", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1.新建server句柄，使用Hua的api
	s := hnet.NewServer("[Hua V0.5]")

	//2.给当前的框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//3.启动server
	s.Serve()
}
