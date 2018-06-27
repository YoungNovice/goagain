package main 

import (
	"fmt"
)

const (
	// bit = 1
	// kb = 1 << 10
	// mb = 1 << 10 << 10
	// tb = 1 << 10 << 10 << 10
	// pb = 1 << 10 << 10 << 10 << 10
	_  = iota
	kb float64 = 1 << (iota * 10)
	mb
	gb 
	tb
	pb
	eb
	zb
	yb

)

func main () {
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(tb)
	fmt.Println(gb)
}