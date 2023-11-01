package hnet

import "Hua/hiface"

/*
实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写
*/
type BaseRouter struct{}

//方法为空，并没有写死，不需要实现，主要用于被其他继承重写

// 在处理conn业务之前的钩子方法
func (br *BaseRouter) PreHandle(request hiface.IRequest) {}

// 在处理conn业务的主方法
func (br *BaseRouter) Handle(request hiface.IRequest) {}

// 在处理conn业务之后的钩子方法
func (br *BaseRouter) PostHandle(request hiface.IRequest) {}
