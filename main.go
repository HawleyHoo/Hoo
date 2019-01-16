package main

import (
	"fmt"
	//"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"
	//"nursing/model"

	"Hoo/h_pkg"
	"time"
	//"github.com/astaxie/beego/session"
	"encoding/json"
	"homedoctor/utils"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"net/url"
	"regexp"
	"errors"
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

func (s Studnt) String() string {
	return fmt.Sprintf("%s, %d")
}

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}

//func filter(r rune) rune {
//	if strings.IndexRune() {
//
//	}
//}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func testsss(str string) string {

	str = "hehe"
	return str
}
func testsssss() {
	fmt.Println("ehhehehehehehehehehehehheheheh")
	var str string = "12345"
	tmpstr := str
	fmt.Printf("%p %p\n", &str, &tmpstr)
	tmpstr = tmpstr + "x"
	fmt.Printf("%p %p\n", &str, &tmpstr)

	var b []string = []string{"b"}
	a := "a"
	fmt.Printf("%p %p\n", &a, &b)
	b = append(b, a)
	fmt.Printf("%p %p\n", &a, &b[1])
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

func testhehe(arr interface{}) [][]interface{} {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return nil
	}
	arrValue := reflect.ValueOf(arr)
	step := 2
	retArr := make([][]interface{}, 0)
	for k := 0; k < arrValue.Len(); k = k + step {
		temp := make([]interface{}, 0)

		temp = append(temp, arrValue.Index(k).Interface())
		if k+1 < arrValue.Len() {
			temp = append(temp, arrValue.Index(k+1).Interface())
		}

		retArr = append(retArr, temp)
	}
	return retArr
}

type Test struct {
	Name string
}

func main() {
	var err error
	err = errors.New("hehe")
	fmt.Printf("%s\v",err)

	phone := "18682169331"
	match,err := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", phone)
	fmt.Println("regexp:", match, err)
	fmt.Println(phone[:8])
	//sbsb := 1
	//fmt.Println(float64(sbsb))
	timestamp := time.Now().Format("0102150405")
	values := url.Values{"access_token": {"token"}}
	values.Add("name", "aaa")
	values.Add("name", "bbb")
	fmt.Println(timestamp, values)
	ts := "20181203 163920"
	t1, err := time.ParseInLocation("20060102 150405", ts, time.Local)
	t2, err := time.Parse("20060102 150405", ts)
	fmt.Println("t1", t1)
	fmt.Println("t2", t2)
	//t3 := time.Now().Add(time.Second * 18).Format("20060102 150405")
	fmt.Println(time.Unix(0,0))
	//schedule.Start()
	//schedule.Every(1).AtOnce(t3).Do(testsssss)
	time.Sleep(200 * time.Second)
	local := time.Now().Local()
	fmt.Println(local)
	//arr1 := []string{"a", "b", "c", "d", "e"}
	//arr1 := []int{1,2, 3, 4, 5, 6}
	arr1 := []Test{Test{"a"}, Test{"b"}, Test{"c"}, Test{"d"}, Test{"e"}, Test{"f"}}
	arr2 := testhehe(arr1)
	fmt.Printf("%+v  \n", arr2)

	return
	//rad := rand.New(rand.NewSource(time.Now().Unix()))
	//for i := 0; i < rad.Intn(9)+1; i++ {
	//	fmt.Println(rad.Intn(50))
	//}
	//testsssss()

	workPath, err := os.Getwd()
	fmt.Println(workPath)
	workPath2, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(workPath2)

	now := time.Now()
	timenow := now.Format("20060102") + " 000000"
	fmt.Println(timenow)
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		list = append(list, i*2)
		if i == 9 {
			h := list[i]
			h = 168
			fmt.Println(h, list[i])
			//list[i] = 168
		}
	}
	fmt.Println(list)

	aaaaaa := "haha"

	fmt.Println("200025893" < "200025892")
	fmt.Println("200025893" <= "200025893")
	fmt.Println("200025893" < "200025894")

	fmt.Println(testsss(aaaaaa))
	fmt.Println(aaaaaa)
	return
	fmt.Println(strings.Replace("10-17-18", "-", "", -1))
	fmt.Println("123456789"[:8])
	for i := 0; i < 100; i++ {
		fmt.Println(RandInt(1000, 9999))
	}
	SaleTime := func(input string, characters string) string { //
		filter := func(r rune) rune {
			if strings.IndexRune(characters, r) < 0 {
				return r
			}
			return -1
		}
		return strings.Map(filter, input)
	}("2016-05-08 23:00:00", "-:")
	fmt.Println(SaleTime)
	//fmt.Println(strings.Map(":","2016-05-08 23:00:00"))
	ru := []rune(":")
	fmt.Println("ele:", strings.IndexRune("20asdf16:88", ru[0]))

	he := ";asdfasdfa;sldkfj;"
	strhe := strings.TrimSuffix(he, ";")

	fmt.Println(he, len("哈哈"))
	fmt.Println(strhe)
	fmt.Println(time.Now().Unix())
	a1 := 0.0000123
	b2 := 0.000012234
	if IsEqual(a1, b2) {
		fmt.Println("a < b", 7/2, h_pkg.Reverse("Hello world!"))
	}
	lsit := make([]string, 0)
	fmt.Println("list :", lsit, "len:", len(lsit))
	storeList := []string{"204149", "204146", "204251", "204086", "204024", "204081", "204220", "204121", "204190", "204007", "204191", "204005", "204010", "204100", "204209", "204221", "204002", "204004", "204008", "204204", "204241", "204206", "204335"}
	for _, v := range storeList {
		if v != "" {
			lsit = append(lsit, v)
		}
	}

	fmt.Println("list :", lsit, "len:", len(lsit))
	return

	hhe := map[string]int{"a": 2, "b": 3}
	if aaaaaaa, ok := hhe["c"]; ok {
		fmt.Println("c:", aaaaaaa)
	} else {
		fmt.Println("no c:", aaaaaaa)
	}

	s := Studnt{}
	fmt.Println(s)
	//t1 := time.Now()
	//t2 := time.Unix(0, 0)

	fmt.Println(t1)
	fmt.Println(t2)

	//aaa := [20]int{1}
	for i := 0; i < 20; i += 3 {
		fmt.Println("hehe", i)
	}

	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	go fmt.Println("hehe")
	fmt.Println("2")
	fmt.Println("22")

	return

	data, i := [3]int{0, 1, 2}, 0
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
