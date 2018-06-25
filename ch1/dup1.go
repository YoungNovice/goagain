package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	// 构造一个map
 	counts := make(map[string]int)
 	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("scanTest is ", scanner.Text())
		if scanner.Text() == "" {
			break
		}
		counts[scanner.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
