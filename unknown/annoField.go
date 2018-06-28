package main

import (
	"fmt"
)
// 匿名字段
type p struct {
	string
	int
}
func main() {
	testCompare()

}

type person struct {
	sex string
	age string
}

// 内容比较 前提是结构类型一样 不然就会报错
func testCompare() {
	a := person{sex:"aa", age:"nv"}
	b := person{sex:"aa", age:"nv"}
	fmt.Println(a == b)
}

func testp() {
	a := &p{"aa", 12}
	fmt.Println(a)
	fmt.Println("========")

	// 匿名字段 匿名结构 真是不知道这个有卵子用
	b := struct{
		int
		string
	}{1, "1"}
	fmt.Println(b)
}