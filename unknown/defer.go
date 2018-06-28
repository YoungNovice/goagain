package main

import (
	"fmt"
)

func main() {
	// testfor()
	// 闭包
	// testf()
	// 传递参数过去就不算闭包了
	testf2()
}

func testf2() {
	for i := 0; i < 3; i++ {
		defer func(j int) {
			fmt.Println(j)
		}(i)
	}
}

// 相当于一个闭包 一直引用i的地址
func testf() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// 这个输出2 1 0 没毛病
func testfor() {
	for i:=0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func test1() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}