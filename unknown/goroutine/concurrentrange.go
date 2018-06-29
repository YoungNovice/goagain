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
		// 关闭chan 不关闭chan range那边会死锁
		close(c)
	}()
	// 使用for range 循环取值 死循环 直到chan被关闭
	// 但是当chan被close时候会跳出循环
	for v := range c {
		fmt.Println(v)
	}
}
