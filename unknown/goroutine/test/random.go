package main 

import (
	"fmt"
)

func main () {
	c := make(chan int)

	go func () {
		for v := range c {
			fmt.Println(v)
		} 
	}()

	// 随机向c中写入0或者1
	// 测试select 随机写入功能
	for {
		select {
		case c<-0:
		case c<-1:
		}
	}
}