基本结构：
* 定义一个消息结构Message
	* 属性
		* 消息ID
		* 消息长度
		* 消息的内容
	 * Setter/Getter方法

抽象层实现/Hua/hiface/imessage.go
```go
package hiface  
  
/*  
将请求的消息封装在一个Message中，定义抽象的接口  
*/  
  
type IMessage interface {  
	//获取消息的ID  
	GetMsgId() uint32  
	//获取消息的长度  
	GetMsgLen() uint32  
	//获取消息的内容  
	GetData() []byte  
	  
	//设置消息的ID  
	SetMsgId(uint32)  
	//设置消息的内容  
	SetData([]byte)  
	//设置消息的长度  
	SetDataLen(uint32)  
}
```
实现层代码/Hua/hnet/message.go
```go
package hnet  
  
type Message struct {  
	Id uint32 //消息id  
	DataLen uint32 //消息长度  
	Data []byte //消息内容  
}  
  
// 获取消息的ID  
func (m *Message) GetMsgId() uint32 {  
	return m.Id  
}  
  
// 获取消息的长度  
func (m *Message) GetMsgLen() uint32 {  
	return m.DataLen  
}  
  
// 获取消息的内容  
func (m *Message) GetData() []byte {  
	return m.Data  
}  
  
// 设置消息的ID  
func (m *Message) SetMsgId(id uint32) {  
	m.Id = id  
}  
  
// 设置消息的内容  
func (m *Message) SetData(data []byte) {  
	m.Data = data 
}  
  
// 设置消息的长度  
func (m *Message) SetDataLen(len uint32) {  
	m.DataLen = len  
}
```