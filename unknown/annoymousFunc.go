package main

import (
	"fmt"
)
// 匿名函数与闭包
func main() {
	a := func () {
		fmt.Println("func no name")
	}

	a()

	fmt.Println("===========")

	sss := 10
	// 测试闭包
	f := closure(sss)
	// closure 函数赋值10本来是一个值传递的
	// 但是调用 f(1) f(2) 打印的x地址都是同一个 说用了什么
	fmt.Println(f(1))
	fmt.Println(f(1))
	// 但是最后sss的值又没变说明了什么
	fmt.Println("sss = ", sss)
	// sss地址和closure内部x的地址不一样又说明了什么
	fmt.Printf("%p\n", &sss)

	// 应该是这样的闭包拷贝了值， 然后这个值在内部环境是可变的
	// 由于不是同一份值 故对外部环境没有影响
	// 但是在内部环境像是在做地址操作一样
}

// 闭包
// 函数的类型怎么确定 去掉名字就是类型了
func closure(x int) (func(int) int) {
	fmt.Printf("%p\n", &x)
	return func (y int) int {
		x = x + y
		fmt.Printf("%p\n", &x)
		return x
	}
}