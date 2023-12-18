package main

import (
	"Hua/hnet"
	"fmt"
	"io"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {

	fmt.Println("Client start....")
	time.Sleep(1 * time.Second)
	//1.连接远程服务器，得到conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error", err)
		return
	}
	for {
		//2.发送封包的message消息 MsgId = 0
		dp := hnet.NewDataPack()
		binaryMsg, err := dp.Pack(hnet.NewMsgPackage(0, []byte("Hua_V0.5 client Test Message")))
		if err != nil {
			fmt.Println("Pack error:", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("Write error:", err)
			return
		}

		//服务器应该返回message数据，msgId = 1

		//3.先读取流的head,得到Id,dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("Read head error", err)
			break
		}
		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println(" client Unpack msgHead error:", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//msgHead有数据
			//4.再根据dataLen进行二次读取，将data读出来
			msg := msgHead.(*hnet.Message) //断言
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error:", err)
				return
			}

			fmt.Println("------>Recv Server Msg : ID = ", msg.Id, ",len = ", msg.DataLen, ",data = ", string(msg.Data))
		}

		//cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
