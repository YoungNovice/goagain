package main

import "fmt"

func main() {
	a := true
	if a, b, c := 1, 2, 3; a+b+c >6 {
		fmt.Println("> 6")
	} else {
		fmt.Println("<= 6")
		fmt.Println(a)
	}
	fmt.Println(a)
}
