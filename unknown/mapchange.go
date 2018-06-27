package main

import (
	"fmt"
)

func main() {
	 m := map[int]string{1:"a", 2:"b"}
	 m1 := make(map[string]int)
	for k, v := range m {
		m1[v] = k
	}

	fmt.Println(m)
	fmt.Println(m1)
}

