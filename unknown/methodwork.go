package main 

import (
	"fmt"
)

type myint int

func (my *myint) Increase(x int) {
	(*my) += myint(x)
}

func main() {
	var a myint = 1
	a.Increase(100)
	a.Increase(50)
	fmt.Println(a)	
}