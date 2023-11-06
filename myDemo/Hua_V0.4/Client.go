package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {

	fmt.Println("Client start")
	//1.连接远程服务器，得到conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start error", err)
		return
	}
	for {
		//2.连接调用Write写数据
		_, err := conn.Write([]byte("Hello Hua V0.2...."))
		if err != nil {
			fmt.Println("write conn error", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back:%s,cnt:%d\n", string(buf), cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
