package main

import (
	"fmt"
)

// go 中没有class 没有
type Person struct {
	name string
	age int
} 

func main() {
	// test()
	test2()
}

func test2() {
	a := &Person{name:"joe", age:19}
	// 值复制大对象 耗内存 我们可以在初始化的时候就获取其指针
	// 这样后面就不用取地址了 而且就算是地址访问内容也可以用点操作
	// a := &Person{name:"joe", age:19}
	a.name = "ok"
	fmt.Println(a)
}

func test() {
	a := Person{name:"joe", age:19}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)
}
// 地址 操作可以修改原始值
func B(per *Person) {
	per.age = 13
}

// 参数值复制
func A(per Person) {
	per.age = 13
	fmt.Println("A", per)
}