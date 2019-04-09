package main

import (
	"fmt"
	//"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"
	//"nursing/model"

	"time"
	//"github.com/astaxie/beego/session"

	"encoding/json"

	"errors"
	"math"
	"nursing/utils"
	"strconv"
	"strings"
	"sync"

	"runtime"
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

//判断a是否等于b
func IsEqualA(a, b float64) bool {
	var r = a - b
	if r == 0.0 {
		return true
	} else if r < 0.0 {
		return r > -0.0001
	}
	return r < 0.0001
}

func parseIP(s string, require int) []string {
	if require == 1 {
		return []string{s}
	}
	r := []string{}
	for i := 1; i < 4 && i+require-1 < len(s); i++ {
		pre := s[:i]
		if v, _ := strconv.Atoi(pre); v < 256 {
			r = append(r, pre)
			r = append(r, parseIP(s[i:], require-1)...)
		}
		if '0' == s[0] {
			fmt.Println("break:", len(s), s)
			break
		}
	}

	return r
}

func addDot(s string, k int) (res string) {
	if len(s) < k {
		return ""
	}
	if len(s) > 3*(k+1) {
		return ""
	}
	return res
}

func isValid(s string) bool {
	if ('0' == s[0] && len(s) > 1) || len(s) > 3 || len(s) == 0 {
		return false
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("strconv err:%s , str:%s", err.Error(), s)
		return false
	}
	if val < 256 && val >= 0 {
		return true
	}
	return false
}
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var y int
	var r int
	for x > y {
		r = x % 10
		x = x / 10
		y = y*10 + r

		if x == y || x/10 == y {
			return true
		}
	}
	return false
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	l := make([]int, 0)

	for x > 9 {
		t := x % 10
		l = append(l, t)
		x = x / 10
	}
	l = append(l, x)
	//fmt.Println(x)
	//for _, v := range l {
	//	fmt.Println(v)
	//}

	length := len(l)
	l2 := length / 2
	for i := 0; i < l2; i++ {
		fmt.Println("v", i, length-1-i, l[i], l[length-1-i])
		if l[i] != l[length-1-i] {
			return false
		}
	}

	return true
}

type Sl []S

func (s *Sl) Remove(value interface{}) error {

	for i, v := range *s {
		fmt.Println("sl:", i, v)
		if v.v == value {

			*s = append((*s)[:i], (*s)[i+1:]...)

			return nil

		}

	}

	return errors.New("not found")

}

func fm(arr []string) {
	for _, v := range arr {
		fmt.Println("v:", v)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

var wg sync.WaitGroup

type Integer int

//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
func (a Integer) Add(b Integer) Integer {
	return a + b
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println("prime:", prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
				//fmt.Println("out", num)
			}
		}
		close(out)
	}()

}

var quit chan int // 只开一个信道

func foo() {
	//fmt.Println(id)
	// ok, finished
	go func() {
		fmt.Println("quit:", <-quit)
		//time.Sleep(1 * time.Second)
	}()
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func say(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		runtime.Gosched()
		fmt.Println(s)
	}
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		//runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	//runtime.GOMAXPROCS(4)
	go say2("1 world")
	say("hello")
}

func main0406() {
	//fmt.Println(multiAdd("1111111", "222222"))

	return

	c := make(chan int)

	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

	return
	count := 10
	quit = make(chan int, 0) // 无缓冲

	for i := 0; i < count; i++ {
		//foo()
		go func(ii int) {
			quit <- ii
		}(i)
	}

	for i := 0; i < count; i++ {
		fmt.Println("quit:", <-quit)
	}

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	go func() {
		//ch <- 4
	}()
	//const zero = 0.0
	//fmt.Println(reflect.TypeOf(zero))
	// 显式地关闭信道
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
	//for v := range ch {
	//	fmt.Println(v)
	//	if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
	//		break
	//	}
	//}
	//fmt.Println(<-ch) // 2
	//fmt.Println(<-ch) // 3
	//fmt.Println(<-ch) // 3
	return

	orgin, wait := make(chan int), make(chan struct{})
	Processor(orgin, wait)
	for num := 2; num < 1000; num++ {
		orgin <- num
	}
	close(wait)
	<-wait

	return
	var ss []int

	ss = append(ss, 1)
	fmt.Println("s:", ss, 100/0.000000000000000001)
	//var m map[string]int
	//m["one"]= 1
	//fmt.Println("m:", m)

	var a Integer = 1

	var b Integer = 2

	var i interface{} = a
	aa := 3
	fmt.Println(&aa)
	sum := i.(Integer).Add(b)

	fmt.Println(sum)

	ar := make([]string, 10)
	for i := 0; i < 10; i++ {
		ar[i] = strconv.Itoa((i + 1) * 10)
	}

	//ar1 := ar[:5]
	//ar2 := ar[5:]
	//wg.Add(2)
	//go fm(ar1)
	//go fm(ar2)
	//wg.Wait()

	sl := Sl{S{0}, S{1}, S{2}, S{3}, S{4}, S{5}, S{6}}
	fmt.Println(sl[:len(sl)-1])
	fmt.Println(sl.Remove(2))
	fmt.Println(sl)
	ii := 1
	fmt.Println(ii)
	ii++
	fmt.Println(ii)
	fmt.Println(ii)
	return

	//defer func() { //必须要先声明defer，否则不能捕获到panic异常
	//	fmt.Println("2")
	//	if err := recover(); err != nil {
	//		fmt.Println(err) //这里的err其实就是panic传入的内容，bug
	//	}
	//	fmt.Println("3")
	//}()
	//s := "25525511135"
	s := "127001"
	for i := 1; i < len(s)-2 && i < 4; i++ {
		for j := i + 1; j < len(s)-1 && j < 4+i; j++ {
			for k := j + 1; k < len(s) && k < 4+j; k++ {
				v1 := s[:i]
				v2 := s[i:j]
				v3 := s[j:k]
				v4 := s[k:]

				if isValid(v1) && isValid(v2) && isValid(v3) && isValid(v4) {
					fmt.Println("-------------------------:" + v1 + "." + v2 + "." + v3 + "." + v4)
				}
			}
		}
	}
	//fmt.Println("IP:", h_math.RestoreIPAddress2("25525511135"))
	//fmt.Println("IP:", h_math.RestoreIPAddress2("4710620350"))
	//fmt.Println("IP:", h_math.RestoreIPAddress2("2552551113"))
	//fmt.Println("IP:", h_math.RestoreIPAddress2("127001"))
	//for i := 1; i <= 30; i++ {
	//	fmt.Println("Fibonacci1:", i, "    ", h_math.Fibonacci(i))
	//
	//}
	//fmt.Println("Fibonacci2:", h_math.Fibonacci2(100))
	//fmt.Println("Fibonacci3:", h_math.Fibonacci3(100))
	//return
	//s1 := "4710620350"
	r := []string{}

	for i := 1; i < len(s); i++ {
		pre := s[:i]
		if v, _ := strconv.Atoi(pre); v < 256 {
			r = append(r, pre)
		}
		s = s[i:]
	}
	fmt.Println(r)
	//fmt.Println("parse:", parseIP(s1, 4))
	return
	//sr := []rune(s)
	l := len(s)
	for i := 0; i < l; i++ {
		fmt.Printf(" %s ", s[:i])
	}
	fmt.Println("\n")
	for _, v := range s {
		fmt.Printf(" %c ", v)
	}
	return
	//fmt.Println(restore(s, 4))
	//fmt.Println(restore(s1, 4))

	fmt.Printf("%0.10f \n", math.Abs(0.00000001-0.000000026))

	arr1 := []string{"1", "2"}
	arr2 := []string{"21", "22"}
	arr3 := []string{}

	arr1 = append(arr1, arr3...)
	arr1 = append(arr1, arr3...)
	arr1 = append(arr1, arr3...)
	arr1 = append(arr1, arr3...)
	arr1 = append(arr1, arr3...)
	fmt.Println("srr1:", arr1, len(arr1))
	arr1 = append(arr1, arr2...)
	fmt.Println("srr1:", arr1, len(arr1))

	//tt := "ab&&2&&fasdfa&&1802395"
	//str := tt
	//n := strings.Count(str, "&&") + 1
	//a := make([]string, n)
	//n--
	//i := 0
	//for i < n {
	//	m := strings.Index(str, "&&")
	//	if m < 0 {
	//		break
	//	}
	//	a[i] = str[:m]
	//	str = str[m+len("&&"):]
	//	i++
	//}
	//a[i] = str
	//fmt.Println("Hoo:", a[:], len(a))
	//res := strings.SplitN(tt, "&&", -1)
	//fmt.Println("res:", res[:3], len(res))
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
	//h_pkg.ReflectTest3()
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

	//h_pkg.Test1()
	//
	//h_pkg.ReflectTest()
	//h_pkg.ReflectTest2()
	//h_pkg.ReflectTest3()

	fmt.Println("------------- 分割线 ---------------")
}
