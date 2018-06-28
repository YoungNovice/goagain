package main 

import (
	"fmt"
)

func main() {
	a := 1
	sum, slice := A(a,2,3,4);
	fmt.Println(sum)
	fmt.Println(slice)
	fmt.Println(a)

	x := 5
	// 将x的地址赋值给函数函数内通过地址改变值导致x内容变化
	intP(&x)
	fmt.Println("x = ", x)

	// 函数是一种类型
	var funa func(string) = FuncA
	TestFunc(funa)
}

func TestFunc (f func(string)) {
	f("testFunc")
}

func FuncA(s string) {
	fmt.Println("Func A", s)
}


// 普通类型的地址传递
func intP(a *int) {
	*a = 100
}

// 可变长参数实际上在函数里面表现为一个slice
// 那么在函数内修改slice的内容会改变么， 实际上不会 因为函数调用只是将参数拿去创建slice
// 属于值拷贝
func A(a ...int) (int, []int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	a[0] = 100
	return sum, a
}

func B(a , b ,c int) (x, y, z int) {
	x, y, z = a, b, c
	return x, y, z 
}