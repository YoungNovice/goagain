package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	// 在一个goroutine中不断的使用select取值
	// 只要chan 不被close掉就不会有问题
	o := make(chan bool, 2)
	go func() {
		for {
			// select 一次代码就会结束
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1 is not ok")
					o <- true
					break
				}
				fmt.Println("c1 is ok and value = ", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2 is not ok")
					o <- true
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

	// 现在只close了c1

	// 不管chan是否被close case始终会运行到 这他么就尴尬了
	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}
