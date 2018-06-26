package main

import "fmt"

/**
	执行结果为什么是这样的
*/
func main() {
	var fs = [4]func(){}
	for i:=0 ; i<4 ; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			// 获取i的引用
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() {
			// 获取i的引用
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}
}