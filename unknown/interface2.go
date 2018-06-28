package main 

// 指针nb些
import (
	"fmt"
)

type A interface {
	Name() string
	Age() int 
}

// 当method has pointer receiver时 
// 接口赋值必须是指针 不然会报错 这些很正常
type person struct {
	name string 
	age int 
}

func (p person) Name () string {
	fmt.Println("person name is :", p.name)
	return ""
}
func (p *person) Age() int {
	return p.age
}

// 两次打印的name都是yang
// 这就尴尬了， 那要是我向要跟着变怎么搞
// 那就取指针撒 
func main() {
	p := person{name : "yang", age :1}
	// 将person强制转为A
	var a A = &p
	a.Name()
	a.Age()
	p.name = "sssss"
	a.Name()
}