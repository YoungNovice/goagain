package main 

import (
	"fmt"
)


// 
func main() {
	// testArr()
	testArrPointer()
}

func testArrPointer() {
	a := [...]int{99:1}
	// 数组的指针 一个指向数组的指针
	var p *[100]int = &a
	fmt.Println(*p)

	// 指针数组 一个数组保存者别人的指针
	x, y := 1, 2
	b := [...]*int{&x, &y}
	for _, v := range b {
		fmt.Println(*v)
	}
}

func testArr () {

	var a [2]int 
	var b [2]int
	b = a
	// 索引赋值
	c := [20]int{19:3}
	d := [...]int{5,6,7}

	// 长度自动推断 索引赋值
	f := [...]int{0:1 , 3:4, 14:9}

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
}