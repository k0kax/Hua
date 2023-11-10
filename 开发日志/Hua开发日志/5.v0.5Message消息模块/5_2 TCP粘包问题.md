采用tlv的方式，对数据包进行处理，解决粘包问题，示意图如下
![](https://cdn.jsdelivr.net/gh/k0kax/PicGo@main/image/202311071755640.png)
将一个完整的消息分割为Head头和Body，Head包括数据的DataLen长度(也就是后面body里面的data的长度)和消息的ID，Body用来承载消息数据的内容
下面还需要进行封包拆包

举个例子

![](https://cdn.jsdelivr.net/gh/k0kax/PicGo@main/image/202311072115622.png)


下面还需要进行封包拆包