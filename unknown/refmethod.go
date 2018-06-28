// 通过反射调用方法
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

func (u User) Hello (name string ) {
	fmt.Println("Hello ", name, "my name is ", u.Name)
}

func main() {
	u := User{1, "ok", 2}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	// 参数
	args := []reflect.Value{reflect.ValueOf("job")}
	mv.Call(args)
}
