package main

import (
	"fmt"
	"sort"
)

// map 的key 必须是可以 == ！= 比较的类型 
// 不可以是 func map slice
func main() {
	// testMapSimple()
	// testMapComplex()
	// testRange()
	testSort()
}


func testSort() {
	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	s := make([]int, len(m))
	i := 0
	// 取 map 的keyset 不同次执行结果可能不一样
	for k, _ := range m {
		s[i] = k
		i++
	}
	// 排序
	sort.Ints(s)
	
	fmt.Println(s)
}


func testRange() {
	// 构造一个slice 里面存的是一个map
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)

	// 上面sm 还是为空 这个因为v得到的是一份拷贝

	// 使用下面的就没有问题了
	for i, _ := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

}

func testMapComplex() {
	var m map[int]map[int]string
	// 初始化最外层map
	m = make(map[int]map[int]string)
	// 给key为1的内层map初始化
	m[1] = make(map[int]string)
	m[1][9] = "ok"

	fmt.Println(m)

}

func testMapSimple() {
	m := make(map[int]string)
	m[1] = "ok"
	m[2] = "aaa"
	fmt.Println(m)
	delete(m, 1)
	temp := m[1]
	if	temp ==  "" {
		fmt.Println("空")
	}
	// key 是 int  value是一个map  这个map的key是int value是string
	var mcom = make(map[int]map[int]string)

	mcom[100] = m

	fmt.Println(mcom)
}