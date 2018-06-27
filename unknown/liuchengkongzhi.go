package main

import (
	"fmt"
)

/*
	go中保留了指针 但是不支持 -> 而是用.访问
	&取地址 *通过指针间接访问对象

	默认值 nil 不是 null
*/

func main () {
	// testPointer()
	testIf()
}


func testIf() {
	if a := 2; a > 3 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}
}


func testPointer() {
	a := 1
	b := &a
	fmt.Println(*b)
}