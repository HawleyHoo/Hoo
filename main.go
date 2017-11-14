package main

import (
	//"Hoo/h_math"
	//"fmt"
	//"math"
	//"Hoo/h_pkg"
	"time"
	"fmt"
	//"Hoo/h_pkg"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"nursing/model"
	"fit"
	"reflect"

	"strings"
)

func main() {
	var date int64 = 1501632000

	datestr := time.Unix(date, 0).Format("2006-01-02 15:04:05")
	fmt.Println("date str :", datestr)

	date21, _ := time.Parse("2006-01-02 15:04:05", "2017-08-02 02:00:00")
	var date22 int64 = date21.Unix()
	time22 := time.Unix(date22, 0)

	datestr2 := time22.Format("2006-01-02 15:04:05")

	fmt.Println(date21, "---", date22)
	fmt.Println("date str2:", datestr2)

	var t int64 = time.Now().Unix()
	var s string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	fmt.Println(t)
	fmt.Println(s)

	aa := strings.Split("1,2", ",")
	fmt.Println(aa)
	/*y0 := math.Sqrt(2)
	y1 := h_math.SqrtByNewton(2)
	y2 := h_math.SqrtNewton(2)
	y3 := h_math.InvSqrt(2)

	fmt.Println(y0)
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)


	h_pkg.Interfacetest()
	//h_pkg.ReflectTest()
	//h_pkg.ReflectTest2()
	//h_pkg.ReflectTest3()
	//h_pkg.ReflectTest5()
	*/
	/*var str = ""
	val, err := strconv.ParseInt(str, 10, 64)
	fmt.Println(val, err)
	//timetest()
	weeks, dates1, dates2, weeknum, err := GetWeeksByDate("2017-10-10", "2017-10-12", 1)
	fmt.Println(weeks)
	fmt.Println(dates1, dates2)
	fmt.Println(weeknum, err)

	str1 := "qwert"
	result := h_pkg.Substr(str1,  1, 3)
	fmt.Println(result)


	nrl3 := model.NRL3{ID:1}

	nrlre, err := QueryNRL("1", nrl3)
	fmt.Println("error", err)
	fmt.Printf("nrl3 %+v\n", nrl3)
	fmt.Printf("result %+v\n", nrlre)

	str11 := fmt.Sprintf("hehe %+v", nrl3)
	fmt.Println("-------", str11)*/
	enginesync()
}

func QueryNRL(rid string, nrl interface{}) (interface{}, error)  {
	var nr3 interface{}
	switch nrl.(type) {
	case model.NRL1:
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

func GetWeeksByDate(datestr, hosdate string, weekindex int) (weeks []time.Time, dates1, dates2 []int, weeknum int, err error) {
	loc := time.Now().Location()
	t, err := time.ParseInLocation("2006-01-02", datestr, loc)
	if err != nil {
		return nil, nil, nil, 0,err
	}

	var hostime time.Time
	if hosdate != "" {
		hostime, _ = time.ParseInLocation("2006-01-02", hosdate, loc)
		//hostime.Add(time.Duration(60))
	}


	// 入院日期到今天的总周数
	weeknum = int(time.Since(t).Hours()/24) / 7 + 1


	//weekoffset := int(time.Since(t) / 24) / 7
	fmt.Println("weeknum", weeknum, t)

	offset := weekindex * 7 - 7
	if weekindex == 0 {
		offset = weeknum * 7 - 7
	}


	t1 := time.Date(t.Year(), t.Month(), t.Day() + offset, 0, 0, 0, 0, loc)
	for i := 0; i < 7; i++ {
		t2 := t1.AddDate(0, 0, +i)
		fmt.Println("time ---", t2.String())
		fmt.Println(t2.Sub(hostime).Hours() / 24)
		weeks = append(weeks, t2)
		dates1 = append(dates1, offset + i + 1)
		// 手术后或产后日期
		hosoffset := t2.Sub(hostime).Hours()
		if hosoffset >= 0 {
			dates2 = append(dates2, int(hosoffset / 24) + 1)
		} else {
			dates2 = append(dates2, 0)
		}
	}
	return weeks, dates1, dates2, weeknum, nil
}

func timetest() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.ISOWeek())
	fmt.Println(t.Weekday())

	t1, err := time.ParseInLocation("2006-01-02 15:04", "2017-10-10 00:00", t.Location() )


	fmt.Println(time.Since(t1).Hours() / 24, time.Since(t1))
	fmt.Println(int(time.Since(t1).Hours() / 24) / 7)
	fmt.Println(int(t.Sub(t1).Hours() / 24) / 7, int(t.Sub(t1).Hours() / 24) % 7)
	fmt.Println("t1:", t1.String())
	if err != nil {
		fmt.Println(err.Error())
	}
	index := t.Weekday()

	weeknum := int(time.Since(t1).Hours() / 24) / 7
	//weekoffset := int(time.Since(t) / 24) / 7
	fmt.Println("weeknum", weeknum)
	t3 := time.Date(t1.Year(), t1.Month(), t1.Day() + 7 * weeknum, 0, 0, 0, 0, t.Location())
	var weeks []time.Time
	for i := 0; i < 7; i++ {
		t2 := t3.AddDate(0, 0, -int(index)+i)
		fmt.Println("time:", t2.String())
		weeks = append(weeks, t2)
	}
	fmt.Println(weeks)
}
