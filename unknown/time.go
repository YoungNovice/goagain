package main

import (
	"fmt"
	"time"
)
func main() {
	// timek()
	closure()
}

func closure() {
	// 最后输出结果都是c 这个闭包造成的 
	// 解决方式 将v作为参数传递过去
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go 	func (){
			fmt.Println(v)
		}()
	}

	select {}
}


func timek() {
		// 获取当前时间
		t := time.Now()
		fmt.Println(t)
	
		// 格式化了
		fmt.Println(t.Format(time.ANSIC))
		// 那么坑在哪里呢， 这个ANSIC是一个字符串 而且还不能变化 
		// 如果取的时间跟他不一样 格式化出来的结果还不一样
}