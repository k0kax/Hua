package hnet

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

}

// 运行服务器
func (s *Server) Serve() {

}

// 停止服务器
func (s *Server) Stop() {

}
