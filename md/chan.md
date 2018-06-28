### Channel

channel 是goroutine的沟通桥梁 大都是阻塞同步的
通过make创建 close关闭
是引用类型
可以使用for range 来迭代不断操作channel
可以设置单向和双向
可以设置缓存大小 在未被填满是不会阻塞

Select 
可以处理一个或多个channel 的发送与接收
同时处理多个可用的channel时按照随机顺序处理
可用空的select 来阻塞main函数
可设置超时