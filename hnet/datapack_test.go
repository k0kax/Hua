package hnet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 只曾是拆包封包的单元测试
func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	//1.创建socketTCP
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	//创建一个goroutine负责承载从客户端的处理业务
	go func() {
		//2.从客户端读取数据，拆包处理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error:", err)
			}

			go func() {
				//处理客户端请求
				//---->拆包过程<-------
				//定义一个拆包的对象
				dp := NewDataPack()
				for {
					//1.第一次从conn，把包的Head读出来
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData) //全读出来了
					/*
						io.ReadFull是连续读取的

						假定file数据为：aaabbbb
						buf := make([]byte, 3)
						n, err := io.ReadFull(file, buf)
						则buf:aaa

						buf2 := make([]byte, 4)
						n2, err2 := io.ReadFull(file, buf2)
						buf:bbbb

					*/
					if err != nil {
						fmt.Println("read head error")
						return
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						//说明Msg有数据，需要二次读取
						//2.第二次从conn，根据head的dataLen,再读出data的内容
						msg := msgHead.(*Message) //+++++++++++++++++++++++++++++++++++++++++++++++断言++++++++++
						msg.Data = make([]byte, msg.GetMsgLen())

						//根据dataLen的长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)

						if err != nil {
							fmt.Println("server unpack data err:", err)
							return
						}
						//完整的消息读取完毕
						fmt.Println("------> Recv MsgId:", msg.Id, "dataLen=", msg.DataLen, "data=", string(msg.Data))
					}
				}
			}()
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	//创建一个封包对象 dp
	dp := NewDataPack()
	//模拟粘包过程，封装两个msg一同发送
	//封装msg1
	msg1 := &Message{
		Id:      1,
		Data:    []byte{'h', 'u', 'a'},
		DataLen: 3,
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack error:", err)
		return
	}
	//封装msg2

	msg2 := &Message{
		Id:      2,
		Data:    []byte{'c', 'a', 'o'},
		DataLen: 3,
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack error:", err)
		return
	}
	//粘包
	sendData1 = append(sendData1, sendData2...) //+++++++++++++++++++++++++++++++++++++++切片的连接问题+++++++++++++++++++++++

	//一股脑发给客户端
	conn.Write(sendData1)

	//客户端阻塞
	select {}
	//取消GlobalObject.Reload()
}
