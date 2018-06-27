package main 

import (
	"fmt"
)

func main() {
	// testSliceMake()
	testReslice()	
}

func testReslice() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s1 := a[2:5]
	fmt.Println(string(s1))
	s2 := s1[1:3]
	fmt.Println( string(s2) )

	// 问题 这个是可以的 只要不超过a的cap就可以
	s3 := s1[3:5]
	fmt.Println(string(s3))

}

func testSliceMake() {
	// type []int  len 4 cap 10
	s := make([]int, 4, 10)
	fmt.Println(s, len(s), cap(s))
}

func testSlice() {
	a := [10]int{1,2,3,4,5,6,7,8,9}
	s := a[1]
	var b int = s 
	fmt.Println(b)

	c := a[5:]
	fmt.Println(c)
}