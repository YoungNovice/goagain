package main 

// 对象赋值给接口的时候是值拷贝
import (
	"fmt"
)
// empty接口中没有方法那么是不是可以看作所有的接口都实现了这个empty接口
// 那么记住go中所有的类型都实现了空接口
type empty interface {
}


// 接口就是一个方法或者多个方法的集合
// Structural Type 只要一个结构时间了某个接口所有的方法那么就算实现了该接口
// 所以接口不需要显示的声明
type USB interface {
	Name() string
	Connecter
}
// Connecter 接口被USB组合
type Connecter interface {
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
	// testInterBase()
	testTransform()
} 

// 类型转换
func testTransform() {
	var pc PhoneConnector 
	pc = PhoneConnector{"PhoneConnector"}

	var c Connecter
	// phoneConnector 转成了Connector
	c = Connecter(pc)
	c.Connect()
}


func testInterBase() {
	var usb USB = PhoneConnector{name:"iphone"}
	fmt.Println(usb.Name())
	usb.Connect()
	// Disconnect(usb)
	Disconnect1(usb)
}


func Disconnect1(usb interface{}) {
	// type switch
	switch v:= usb.(type) {
	case PhoneConnector :
		fmt.Println("Disconnnected:", v.name)
	default:
		fmt.Println("Disconnected.:")
	}
}

func Disconnect(usb empty) {
	// 类型断言
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("Disconnnected:", pc.name)
		return
	}
	fmt.Println("Disconnected.:")
}