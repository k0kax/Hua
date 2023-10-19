package main

import "Hua/hnet"

/*
基于Hua开发的 服务器端应用程序
*/
func main() {
	//1.新建server句柄，使用Hua的api
	s := hnet.NewServer("[Hua V0.1]")

	//2.启动server
	s.Serve()
}
