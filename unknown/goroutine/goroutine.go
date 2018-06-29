package main 

import (
	"fmt"
	"runtime"
)

// 通过缓存的chan提高多个goroutine并发问题
func main () {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i:= 0; i<10 ; i++ {
		go Go(c, i)
	}

	for i:=0; i<10; i++ {
		<- c
	}
}

func Go (c chan bool, index int) {
	a := 1
	for i:= 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	// 计算完成后向chan中写入一个true
	c <- true
}