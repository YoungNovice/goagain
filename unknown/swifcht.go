package main

import "fmt"

func main() {
	//test1()
	//test2()
	test3()
}

func test3()  {
	// 这里为什么要加一个分号
	switch a := 1; {
	case a == 1:
		fmt.Println("a == 1")
	case a == 2:
		fmt.Println("a == 2")
	}

}

func test2()  {
	a := 1
	switch  {
	case a == 1:
		fmt.Println("a == 1")
	case a == 2:
		fmt.Println("a == 2")
	}
	fmt.Println(a)
}

func test1()  {

	a := 1
	switch a {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	}
	fmt.Println(a)

}
