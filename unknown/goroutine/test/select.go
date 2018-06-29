package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	// 在一个goroutine中不断的使用select取值
	go func() {
		for {
			// select 一次代码就会结束所以外面要套一个死循环
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1 is not ok")
					break
				}
				time.Sleep(2 * time.Second)
				fmt.Println("c1 is ok and value = ", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2 is not ok")
					break
				}
				fmt.Println("c2 is ok and value = ", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	// 不close会造成for死循环 这是正常人的想法 实际上它正常退出了
	// 他退出不是因为死循环活了 而是主线程结束了
	// close(c1)
	// close(c2)
}
