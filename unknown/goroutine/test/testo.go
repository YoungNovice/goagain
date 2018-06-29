package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	// 在一个goroutine中不断的使用select取值
	o := make(chan bool)
	go func() {
		a := 1
		for {
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

	// 任意关闭一个chan 
	close(c1)
	<-o
	// 结论无论chan是否管理select的case都会进去 只是关闭后ok = false
	// 所以我们没法让c1 c2同时关闭
}
