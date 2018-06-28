package main  

import (
	"fmt"
	"reflect"
)
type User struct {
	Id int 
	Name string 
	Age int 
}

func (u User) Hello () {
	fmt.Println("Hello World")
}

type Manager struct {
	User
	title string
}

// 反射
func main() {
	// u := User{1, "ok", 13}
	// Info(&u)

	m := Manager{User:User{2, "no", 3}, title:"title是xxx"}

	t := reflect.TypeOf(m)

	// 获取user
	fmt.Printf("%#v\n", t.Field(0))
	// 获取id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0}) )
}

func Info2 (o interface{}) {

}

// 反射值只能反射
func Info (o interface{}) {
	// 获取接口的类型
	t := reflect.TypeOf(o)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("is not struct")
		return
	}

	// 打印类型的名字
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o);
	fmt.Println("Fields:")
	fmt.Println("v is:", v)
	// 反射成员变量
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()

		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	// 反射方法信息
	for i:= 0; i< t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}