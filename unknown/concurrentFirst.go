package main 

import (
	"fmt"
	"time"
)

// 主线程不会等待goroutine执行的 所以可能看不到打印的信息
// 可以让主线程sleep 2秒来看到日志
func main() {
	go Go()
	time.Sleep(2 * time.Second)
}

func Go() {
	fmt.Println("go go go")
}