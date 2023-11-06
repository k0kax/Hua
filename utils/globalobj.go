package utils

import (
	"Hua/hiface"
	"encoding/json"
	"io/ioutil"
)

/*
储存一切有关框架的全局参数，供其他模块使用
一切参数皆可通过hua.json实现
*/
type GlobalObj struct {
	/*
		Server的配置信息
	*/

	TcpServer hiface.IServer
	Host      string
	TcpPort   int
	Name      string

	/*
		hua框架
	*/
	Version        string //hua版本号
	MaxConn        int    //当前服务器允许的最大连接数
	MaxPackageSize uint32 //当前hua框架数据包的最大值
}

/*
定义一个全局的对外GlobalObj
*/
var GlobalObject *GlobalObj

// 提供一个初始化当前GlobalObject对象
func init() {
	//如果配置文件没有加载的默认值
	GlobalObject = &GlobalObj{
		Name:           "HuaServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//加载
	GlobalObject.Reload()

}

// 从conf/hua.json文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/hua.json")
	if err != nil {
		panic(err)
	}
	//将json解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
