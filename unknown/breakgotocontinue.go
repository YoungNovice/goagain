package main

import "fmt"

func main() {
	//testa()
	//testb()
	testc()
}

func testc() {
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//continue LABEL
			goto LABEL1
		}
	LABEL1:
	}

}

func testb() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
}

func testa() {
LABLE:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				break LABLE
			} else {
				fmt.Println(i)
			}
		}
	}

}
