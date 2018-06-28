package main 

import (
	"fmt"
)

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

type TZ int

func (a *TZ) Print (xxx int) {
	fmt.Println("TZ = ", xxx)
}

func (a TZ) print2(xxx int) {
	fmt.Println("TZ  value = ", xxx)
}

// TZ 是int的别名 但是TZ没有int的方法
// 可以给TZ添加方法
// TZ 不可以转换成int 他们方法都不一样怎么可能转
func testTZ() {
	var a TZ
	// method value 调用方式
	a.Print(1)
	// method expression 调用方式
	(*TZ).Print(&a, 1)

	// method value 调用方式
	a.print2(3)
	// method expression 调用方式
	TZ.print2(a, 3)
}

// go的method 挂在到strut
// 直接收者 指针接收者
func main() {
	// testA()
	// testX()
	testTZ()
}

// 通过指针接收着
func testX() {
	a := A{}
	a.PrintX()
	fmt.Println(a.Name)
}

// 通过值接收着
func testA() {
	a := A{}
	a.Name = "yang"
	a.Print()
}

// 值接收着
func (a A) Print() {
	fmt.Println(a.Name)
} 

// 指针接收着
func (a *A) PrintX() {
	a.Name = "AA"
	fmt.Println("A")
}