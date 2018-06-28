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

func main() {
	// testInt()
	testUser()
}

func testUser() {
	u := User{1, "ok", 2}
	Set(&u)
	fmt.Println(u)
}
func Set (o interface{}){
	v := reflect.ValueOf(o)
	// 判断传进来的是不是pointer-interface

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Filed Valid")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}

}


// 修改一个基本类型int的值
func testInt() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println("x = ", x)
}