package main 

import (
	"fmt"
)

/*
	6 :   0110
	11 :  1011

	&    0010 = 2  (全是1才是1 否则就是0)
	|    1111 = 15 (有1就是1 )
	^    1101 = 13 (只有一个1 才成立)
	&^   0100 = 4 (第二个操作数是1 就将第一个操作数改为0 否则不变)

*/
func main () {
	a := 6
	b := 11
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}