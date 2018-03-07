package main

import (

	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"nursing/model"
	"fit"

	"Hoo/h_pkg"
	"nursing/utils"
	"reflect"
	"time"
	"strconv"
	"database/sql"
	//"github.com/astaxie/beego/session"
	"strings"
	"encoding/json"
)

type Man struct {
	Name string
}

type Studnt struct {
	Man
	Code int
}

var timeChina =  []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
var dateChina =  []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}

func sinicizingDate(intval int)  {
	for index := len(strconv.Itoa(intval)); index >= 0 ; index++  {
		//ii := intval % 10
	}
}

func sinicizingTime(intval int) (datestr string) {
	if intval >= 0 && intval <= 60 {
		if intval == 0 {
			datestr = "零"
		} else if intval < 10 {
			datestr = timeChina[intval]
		} else {
			datestr = timeChina[intval / 10] + "十" + timeChina[intval % 10]
		}
	} else {
		datestr = ""
	}
	return
}

func initEngine() (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", "phpgroup:fitcome_meal1!qw2@tcp(39.108.133.131:1714)/nursing?charset=utf8")
	//db, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3307)/nuring?charset=utf8")
	db.TZLocation = time.Now().Location()
	db.DatabaseTZ = time.Now().Location() // Now().Location()
	//SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper
	//映射同名设置默认
	db.SetMapper(core.SameMapper{})
	fmt.Println("------hahaha")
	if err == nil {
		return db, err
	}
	return nil, err
}

func initmysqlEngine() (*sql.DB, error)  {
	db, err := sql.Open("mysql", "phpgroup:fitcome_meal1!qw2@tcp(39.108.133.131:1714)/nursing?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func splitTag(tag string) (tags []string) {
	tag = strings.TrimSpace(tag)
	var hasQuote = false
	var lastIdx = 0
	for i, t := range tag {
		if t == '\'' {
			hasQuote = !hasQuote
		} else if t == ' ' {
			if lastIdx < i && !hasQuote {
				tags = append(tags, strings.TrimSpace(tag[lastIdx:i]))
				lastIdx = i + 1
			}
		}
	}
	if lastIdx < len(tag) {
		tags = append(tags, strings.TrimSpace(tag[lastIdx:]))
	}
	return
}

func trimTest1()  {
	defer utils.Trace("hoo test1")()
	tagstr := " hy pk - autoincre"
	for index := 0; index < 10000; index++ {
		tags := splitTag(tagstr)
		if index == 10 {
			fmt.Println("tags:", tags, "len:", len(tags))
		}
	}
}

func trimTest2()  {
	defer utils.Trace("hoo test2")()
	tagstr := " hy pk - autoincre"
	tagstr = strings.TrimSpace(tagstr)
	for index := 0; index < 10000; index++ {
		tags := strings.Split(tagstr, " ")
		if index == 10 {
			fmt.Println("tags:", tags, "len:", len(tags))
		}
	}
}

type tempChart struct {
	Temp1     [42]string // 口表
	Temp2     [42]string // 腋表
	Temp3     [42]string // 肛表
	TempOther []string // 体温事件
	Pulse     []string // 脉搏
	Heartrate []string // 心率
	Breathe   []string // 呼吸
	Intake    []string // 输入液量
	Output1   []string // 排出其他
	Output2   []string // 排出尿量
	Output3   []string // 排出大便
	Pressure  []string // 血压
	Weight    []string // 体重
	Skin      []string // 皮试
	Other     []string // 其他
	Incident  []string // 事件
}

type Result struct {
	Status int
}

type S struct {
	v int
}
type AccessType int

const (
	//
	AccessTypeBack AccessType = 1 << iota
	AccessTypeOut
	AccessTypeAll
	AccessTypeUnknown = 0
)
func main123() {


	var data = []byte(`{"Status": 200}`)
	result := Result{}

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result=%+v", result)
	return
	//data, i := []int{0,1,2}, 0
	fmt.Println("---:", AccessTypeUnknown)
	fmt.Println("---:", AccessTypeBack)
	fmt.Println("---:", AccessTypeOut)
	fmt.Println("---:", AccessTypeAll)


	s := []S{{1}, {3}, {5}, {2}}
	fmt.Printf("%#v", s)
	// A

	fmt.Printf("sort：%#v", s)


}

func main() {
	//data, i := [3]int{0,1,2}, 0
	//i, data[i] = 2, 100
	//fmt.Println(data, i)
	//data[i], i = 100, 2
	//fmt.Println(data, i)

	fmt.Println(" SplitN 函数的用法")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1))

	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]

	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]
	sss := "a,b"
	ss := "a"
	sli1 := strings.Split(sss, ",")
	sli2 := strings.Split(ss, ",")
	fmt.Println("sli :", sli1, len(sli1))
	fmt.Println("sli :", sli2, len(sli2))


	fmt.Println("---:", 0.00001 > 0)
	fmt.Println("---:", -0.00001 > 0)
	datetime ,_ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-02 01:10:08", time.Local)

	datestr := fmt.Sprintln(datetime.Year(), "年", "月", "日", sinicizingTime(datetime.Hour()), "时", sinicizingTime(datetime.Minute()), "分", sinicizingTime(datetime.Second()), "秒")

	fmt.Println("time:", datestr, len("03"))

	list1 := [8]string{"a", "b", "c", "d", "e", "f"}
	list2 := make([]string, 6)

	fmt.Println("list:", list1, list2)
	str := "ab&&2"
	n := strings.Count(str, "&&") + 1
	a := make([]string, n)
	n--
	i := 0
	for i < n  {
		m := strings.Index(str, "&&")
		if m < 0 {
			break
		}
		a[i] = str[:m]
		str = str[m+len("&&"):]
		i++
	}
	a[i] = str
	fmt.Println("Hoo:", a[:i+1])
	return
	date2 := time.Now().Unix()
	fmt.Println("-----", date2, date2 * 1000)
	return

	//enginesync()
	chart := new(tempChart)
	chart.Temp1[2] = "24"
	fmt.Printf("hehe: %+v", chart)

	tagstr := " hy pk autoincre"
	fmt.Println("tag str 1:", tagstr)
	tagstr = strings.TrimSpace(tagstr)
	fmt.Println("tag str 2:", tagstr)
	for k, v := range tagstr {
		fmt.Println("range tag str :", k, "->", v)
	}
	h_pkg.ReflectTest3()
	return
	trimTest1()
	trimTest2()



	defer utils.Trace("hoo test")()

	engine, err := initEngine()
	//engine, err := initmysqlEngine()
	if err != nil {
		fmt.Println("err:",err)
	}
	//session := engine.NewSession()
	//
	//defer session.Close()
	//errsession := session.Begin()
	//if errsession != nil {
	//	fmt.Println("err:", errsession)
	//}
	for index := 0; index < 1; index++  {
		//resultmap, errengine := session.QueryString("select value from temperatrurechat where datetime = ? and patientid = ? and headtype in (1,4,5,8,9,10,13,14)", "2017-11-28 00:00:00", "806174")
		//var resultmap []map[string]interface{}
		//errengine := engine.SQL("select value from temperatrurechat where datetime = ? and patientid = ? and headtype in (1,4,5,8,9,10,13,14)", "2017-11-28 00:00:00", "806174").Find(&resultmap)
		resultmap, errengine := engine.Query("select value from temperatrurechat where datetime = ? and patientid = ? and headtype in (1,4,5,8,9,10,13,14)", "2017-11-28 00:00:00", "806174")
		if index == 400 {
			fmt.Println("index:", index, "result:", resultmap, "err:", errengine)
		}
	}
	//errcommit := session.Commit()
	//fmt.Println("err commit:", errcommit)





	str11 := ""
	fmt.Println("------", utils.Substr(str11, 0, 10))
	for index := 1;index < 36 ;index++  {
		switch index {
		case 18:
			fmt.Println("  hehehe ", index)
		case 31, 32, 33:
			fmt.Println(" hahaha + ", index)
		}
	}



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
	//engine,_ := xorm.NewEngine("mysql", "phpgroup:fitcome_meal1!qw2@tcp(114.119.10.182:1714)/nursing?charset=utf8")
	//engine,_ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3307)/nuring?charset=utf8")
	engine,_ := xorm.NewEngine("mysql", "youhao:youhao@tcp(192.168.0.126:3306)/nursing?charset=utf8")
	engine.SetMapper(core.SameMapper{})
	engine.Sync(new(model.NRL3), new(model.NRL4), new(model.NRL5), new(model.NRL6), new(model.NRL7), new(model.NRL8), new(model.NRL2), new(model.NRL1Title), new(model.NRL7Title), new(model.IOStatistics))
	// new(model.NRL3), new(model.NRL4), new(model.NRL5), new(model.NRL6), new(model.NRL7),
	//{{if eq $val.NRL01 "1"}}A{{else if eq $val.NRL01 "2"}}B{{else if $val.NRL01 "3"}}C{{end}}
}


