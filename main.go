package main

import (
	"fmt"
	//"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"
	//"nursing/model"

	"Hoo/h_pkg"
	"nursing/utils"
	"time"
	//"github.com/astaxie/beego/session"
	"encoding/json"
	"homedoctor/utils"
	"strings"
)

type Man struct {
	Name string
}

type Studnt struct {
	Man
	Code int
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

func trimTest1() {
	defer utils.Trace("hoo test1")()
	tagstr := " hy pk - autoincre"
	for index := 0; index < 10000; index++ {
		tags := splitTag(tagstr)
		if index == 10 {
			fmt.Println("tags:", tags, "len:", len(tags))
		}
	}
}

func trimTest2() {
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
	TempOther []string   // 体温事件
	Pulse     []string   // 脉搏
	Heartrate []string   // 心率
	Breathe   []string   // 呼吸
	Intake    []string   // 输入液量
	Output1   []string   // 排出其他
	Output2   []string   // 排出尿量
	Output3   []string   // 排出大便
	Pressure  []string   // 血压
	Weight    []string   // 体重
	Skin      []string   // 皮试
	Other     []string   // 其他
	Incident  []string   // 事件
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
	AccessTypeBack    AccessType = 1 << iota
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

	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	go fmt.Println("hehe")
	fmt.Println("2")
	fmt.Println("22")

	return

	data, i := [3]int{0,1,2}, 0
	i, data[i] = 2, 100
	fmt.Println(data, i)
	data[i], i = 100, 2
	fmt.Println(data, i)

	return

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

	list1 := [8]string{"a", "b", "c", "d", "e", "f"}
	list2 := make([]string, 6)

	fmt.Println("list:", list1, list2)
	str := "ab&&2"
	n := strings.Count(str, "&&") + 1
	a := make([]string, n)
	n--
	i = 0
	for i < n {
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
	fmt.Println("-----", date2, date2*1000)
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

	str11 := ""
	fmt.Println("------", utils.Substr(str11, 0, 10))
	for index := 1; index < 36; index++ {
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
}
