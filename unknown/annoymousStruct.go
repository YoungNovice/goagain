package main

import (
	"fmt"
)

type person struct {
	Name string
	Age int
	// 在结构中声明匿名结构
	Contact struct {
		Phone,City string
	}
}
func main() {
	// 直接声明匿名结构 定义带初始化
	b := &struct {
		Name string
		Age int
	}{Name: "joe", Age:19}
	fmt.Println(b)
	fmt.Println("=====")

	a := person{Name:"a", Age:19}
	a.Contact.Phone = "122"
	a.Contact.City = "beijing"
	fmt.Println(a)
}