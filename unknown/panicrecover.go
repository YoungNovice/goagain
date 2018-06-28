package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

// panic之后程序就终止了
// 怎么恢复
func B() {
	// defer 要放在panic 之前
	// 因为panic之后的代码无法执行下去 是无效的
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B", err)
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}