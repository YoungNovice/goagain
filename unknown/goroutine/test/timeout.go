package main 

import (
	"fmt"
	"time"
)

// 第一个case从c中取数组 但是我们并没有向c中放入任何值
// 第二个case中是一个定时器
func main () {
	c := make(chan bool)
	select {
	case v := <-c :
		fmt.Println("v = ", v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
		 
}