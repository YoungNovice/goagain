package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	bytes, e := ioutil.ReadFile("fileopen/aa.txt")
	if e != nil {
		return
	}
	fmt.Println("%s", bytes)

}
