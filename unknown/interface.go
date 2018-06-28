package main 

import (
	"fmt"
)
// 接口就是一个方法或者多个方法的集合
// Structural Type 只要一个结构时间了某个接口所有的方法那么就算实现了该接口
// 所以接口不需要显示的声明

type USB interface {
	Name() string
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	testInterBase()

}

func testInterBase() {
	var usb USB = PhoneConnector{name:"iphone"}
	fmt.Println(usb.Name())
	usb.Connect()
}