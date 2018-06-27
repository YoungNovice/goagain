package main 

import (
	"fmt"
)


// 
func main() {
	// testArr()
	// testArrPointer()
	// testArrCompare()
	// testArrNew()
	testSort()
}

func testSort() {
	a := [...]int{5,2,6,3,0}
	fmt.Println("before sort", a)

	for i, num := 0, len(a); i < num; i++ {
		for j := i+1 ; j < num ; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println("after sort", a)
}

// 使用数组的指针和数组本身都可以通过[]索引赋值
func testArrNew() {
	// 用new 也可以创建数组 返回的是指向数组的指针
	p := new([10]int)
	p[2] = 10
	(*p)[3] = 20
	fmt.Println(p) 
}


/*
	数组比较 只能同种类型 同长度比较 == != 两种
*/
func testArrCompare() {
	a := [2]int{1,2}
	b := [2]int{1,2}
	fmt.Println(a == b)
}


/*
	数组指针问题
*/
func testArrPointer() {
	a := [...]int{99:1}
	// 数组的指针 一个指向数组的指针
	var p *[100]int = &a
	fmt.Println(p)

	// 指针数组 一个数组保存者别人的指针
	x, y := 1, 2
	b := [...]*int{&x, &y}
	for _, v := range b {
		fmt.Println(*v)
	}
}

/*
	数组初始化问题
*/
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