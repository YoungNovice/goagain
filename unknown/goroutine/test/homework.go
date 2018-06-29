package main 

import (
	"fmt"
)

var c chan string
// 创建一个goroutine与主线程相互发送信息若干次打印
func main() {
	c = make(chan string)
	go Go()

	for i:=0; i<10; i++ {
		c <- fmt.Sprintf("From main: hello, #%d", i)
		fmt.Println(<-c)
	}
}

func Go() {
	i := 0
	for {
		// 接收信息
		fmt.Println(<-c) 
		// 发送信息
		c<- fmt.Sprintf("From pingpong hi, #%d", i)
		i++
	}
}