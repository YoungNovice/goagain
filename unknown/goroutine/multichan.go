package main 

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	// 在一个goroutine中不断的使用select取值
	// 只要chan 不被close掉就不会有问题
	go func() {
		for {
			// select 一次代码就会结束
			select {
			case v, ok := <- c1:
				if !ok {
					break
				}
				fmt.Println("c1 v = ", v)
			case v, ok := <- c2:
				if !ok {
					break
				}
				fmt.Println("c2 v = ", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1)
	close(c2)
}