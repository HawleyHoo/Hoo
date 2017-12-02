package main

import (
	//"Hoo/h_math"
	//"fmt"
	//"math"
	//"Hoo/h_pkg"
	"fmt"
	"Hoo/h_pkg"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"nursing/model"
	"fit"
	"reflect"

	"time"
)

type Man struct {
	Name string
}

type Studnt struct {
	Man
	Code int
}

func main() {
	fmt.Println(fmt.Sprintf("a%.0f", 0.0/24.0))
	fmt.Println(fmt.Sprintf("%.0f", 0.0/24))


	for index := 1;index < 36 ;index++  {
		switch index {
		case 18:
			fmt.Println("  hehehe ", index)
		case 31, 32, 33:
			fmt.Println(" hahaha + ", index)
		}
	}

	t1str := "2018-12-08 12:00:00"
	t1time, _ := time.ParseInLocation("2006-01-02 15:04:05", t1str, time.Local)
	if t1time.Before(time.Now()) {
		fmt.Println("t1time has arrived")
	} else {
		fmt.Println("t1time hasn't come yet")
	}
	fmt.Println("2016-05-20 09:30:29" > "2016-05-20 08:50:12")
	fmt.Println("2016-05-21 09:30:29" > "2016-05-20 08:50:12")
	fmt.Println("08:50" > "")

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
	engine.Sync(new(model.NRL7Title))
	// new(model.NRL3), new(model.NRL4), new(model.NRL5), new(model.NRL6), new(model.NRL7),
	//{{if eq $val.NRL01 "1"}}A{{else if eq $val.NRL01 "2"}}B{{else if $val.NRL01 "3"}}C{{end}}
}


