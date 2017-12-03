package main

import (

	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"nursing/model"
	"fit"
	"time"
	"strings"
	"Hoo/h_pkg"
	"nursing/utils"
	"reflect"
)

type Man struct {
	Name string
}

type Studnt struct {
	Man
	Code int
}

func main() {


	teststr1 := "呵呵，"
	rs := []rune(teststr1)
	aaa := rs[0:len(rs) - 1]
	fmt.Println("test:", string(teststr1[0:len(teststr1) - 1]), string(aaa))
	fmt.Println("test1:", utils.Substr(teststr1, 0, 2))

	h_pkg.ReflectTest()
	h_pkg.ReflectTest2()
	h_pkg.ReflectTest3()
	h_pkg.ReflectTest4()



	h_pkg.Test1()

	h_pkg.ReflectTest()
	h_pkg.ReflectTest2()
	h_pkg.ReflectTest3()

	fmt.Println("------------- 分割线 ---------------")
	rt := reflect.TypeOf(model.NRL3{})
	rv := reflect.ValueOf(model.NRL3{})

	for index := 0; index < rt.NumField() ;index++  {
		fmt.Println("Name :", rt.Field(index).Name, rv.Field(index).Interface(), "field type:", rv.Field(index).Type(), "field kind:", rv.Field(index).Kind(), "kind:", rt.Kind(), "type:", rv.Type())
	}
	//Name : NRL01 0 field type: int field kind: int kind: struct type: model.NRL3
	//enginesync()
}

func QueryNRL(rid string, nrl interface{}) (interface{}, error)  {
	var nr3 interface{}
	switch nrl.(type) {

	case model.NRL3:
		//var nr3 model.NRL3
		nr3 = model.NRL3{}
		fmt.Println(reflect.TypeOf(nrl))
	default:
		break
	}
	//var nr3 model.NRL3
	fit.MySqlEngine().ShowSQL(true)
	_, err := fit.MySqlEngine().Table("NRL3").Where("ID = ?", rid).Get(&nr3)
	fit.MySqlEngine().ShowSQL(false)
	fmt.Println(nrl)
	if err != nil {
		return nil, err
	} else {
		fmt.Println(nrl)
		return nr3, nil
	}
}


func enginesync()  {
	engine,_ := xorm.NewEngine("mysql", "phpgroup:fitcome_meal1!qw2@tcp(114.119.10.182:1714)/nursing?charset=utf8")
	engine.SetMapper(core.SameMapper{})
	engine.Sync(new(model.NRL1Title))
	// new(model.NRL3), new(model.NRL4), new(model.NRL5), new(model.NRL6), new(model.NRL7),
	//{{if eq $val.NRL01 "1"}}A{{else if eq $val.NRL01 "2"}}B{{else if $val.NRL01 "3"}}C{{end}}
}


