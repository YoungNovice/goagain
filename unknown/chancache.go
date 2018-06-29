package main 

import (
	"fmt"
)
// 无缓存chan 取操作在前放操作在后
// 
func main() {
	// chan 有缓存 主线程中写入true并不会等待子线程取值就退出了
	var c chan bool = make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c
	}()
	c <- true
}