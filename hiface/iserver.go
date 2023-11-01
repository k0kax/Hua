package hiface

/*
	定义一个服务器接口
*/

type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//停止服务器
	Stop()

	//路由功能：给当前的服务注册一个路由方法，供给客户端链接使用
	AddRouter(router IRouter)
}
