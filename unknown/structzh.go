package main

import (
	"fmt"
)
//结构的组合

type human struct {
	Sex int
	Name string
	Age int
}
type teacher struct {
	// 嵌入结构 同时是匿名字段
	human
}
type student struct {
	h human
	score int 
}
func main() {
	t := teacher{human{Name:"t", Age:28, Sex:0}}
	s := student{h:human{Name:"s",Age:20, Sex:1}, score:3}
	fmt.Println(t)
	t.Sex = 10
	fmt.Println(t)
	s.h.Sex = 20
	fmt.Println(s)
}