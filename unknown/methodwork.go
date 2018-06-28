package main 

import (
	"fmt"
)

type myint int

func (my *myint) Increase() {
	(*my) += 100
}

func main() {
	var a myint = 1
	a.Increase()
	a.Increase()
	fmt.Println(a)	
}