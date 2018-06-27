package main 

import (
	"fmt"
)

func main() {
	// testSliceMake()
	// testReslice()	
	// testAppend()
	// testOther()
	// testReMake()
	testCopy()
}

func testCopy() {
	s1 := []int{1,2,3,4,5,6}
	s2 := []int{7,8,9,10,11,12,13,14}
	// 将s2 copy 到 s1 中
	copy(s2, s1)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
}

func testReMake() {
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	// s1 导致重新分配不再指向a了
	s1 = append(s1, 1,2,3,4,5,6,7,8,9)
	s2 := a[1:3]
	s1[0] = 9
	fmt.Println(s1, s2, a)
}

// 当s1中的第0个位置改变 会导致a的第二个位置改变 所以s2的值也会改变
// 但是如果中途使用了append而导致内存重新分配会怎样 
func testOther() {
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	s1[0] = 9
	fmt.Println(s1, s2, a)
}

func testAppend() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	// 不超过cap的个数使用的是同一个底层的数组
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	// 超过cap的个数会导致数组内存重新分配， 从而地址改变
	s1 = append(s1, 1,2,3)
	fmt.Printf("%v %p\n", s1, s1)
}

func testReslice() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s1 := a[2:5]
	fmt.Println(string(s1))
	s2 := s1[1:3]
	fmt.Println(string(s2))
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