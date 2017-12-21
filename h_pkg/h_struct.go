package h_pkg

import (
	"reflect"
	"fmt"
)

type Man struct {
	Name   string "user name " //这引号里面的就是tag
	Passwd string "user passsword"
}

func StructTest1() {
	user := &Man{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}

type S struct {
	A string `hy:"hehe" color:"red" `
	F string `hy:"haha" color:"red" fit:"-"`
}

func StructTest2() {
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
