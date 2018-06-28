package main

import (
	"fmt"
	"reflect"
)

func main() {
	// testInt()
	
}

// 修改一个基本类型int的值
func testInt() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println("x = ", x)
}