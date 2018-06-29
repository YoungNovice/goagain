package main

import (
	"fmt"
	"time"
)
// 通过通信共享内存 而不是通过共享内存来通信
func main () {
	// 配置一个chan
	c := make(chan bool)
	go func() {
		// 睡眠三秒
		time.Sleep(3 * time.Second)
		fmt.Println("GO GO GO")
		// 向通道中传递值
		// 当通道中有值的时候会一直阻塞
		c <- true
	}()
	// 取出通道中的值
	// 当通道中没有值的时候会一直阻塞
	<- c	
}
