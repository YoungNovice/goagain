package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	// 相对路径
	// files , err := os.Open("fileopen/aa.txt")
	// 绝对路径
	files , err := os.Open("D:\\ideaGit\\goagain\\fileopen\\aa.txt")
	if err != nil {
		fmt.Println("file not exists")
		return 
	}
	scanner := bufio.NewScanner(files)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}