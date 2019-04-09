package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// 字符串反转
func Reverse(s string) string {
	r := []rune(s)
	count := len(s) / 2
	for i, j := 0, len(r)-1; i < count; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func Test() {
	//第一种连接方法（最快）

	var buffer bytes.Buffer
	s := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("test is here\n")
	}
	buffer.String() // 拼接结果
	e := time.Now()
	fmt.Println("1 time is ", e.Sub(s).Seconds())

	//第二种方法
	s = time.Now()
	var sl []string
	for i := 0; i < 100000; i++ {
		sl = append(sl, "test is here\n")
	}
	strings.Join(sl, "")
	e = time.Now()
	fmt.Println("2 time is", e.Sub(s).Seconds())

	//第三种方法
	s = time.Now()
	str := ""
	for i := 0; i < 100000; i++ {
		str += "test is here\n"
	}
	e = time.Now()
	fmt.Println("3 time is ", e.Sub(s).Seconds())

	//第四种方法
	s = time.Now()
	str4 := ""
	for i := 0; i < 100000; i++ {
		str4 = str4 + "test is here"
	}
	e = time.Now()
	fmt.Println("4 time is ", e.Sub(s).Seconds())

}

func Test2() {
	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}

func teststringsplitn() {
	fmt.Println(" SplitN 函数的用法")                                             // n 代表分隔成几份，-1是所有
	fmt.Printf("SplitN 0 : %q\n", strings.SplitN("/home/m_ta/src", "/", 0))  // []
	fmt.Printf("SplitN 1 : %q\n", strings.SplitN("/home/m_ta/src", "/", 1))  // ["/home/m_ta/src"]
	fmt.Printf("SplitN 2 : %q\n", strings.SplitN("/home/m_ta/src", "/", 2))  // ["" "home/m_ta/src"]
	fmt.Printf("SplitN 3 : %q\n", strings.SplitN("/home/m_ta/src", "/", 3))  // ["" "home" "m_ta/src"]
	fmt.Printf("SplitN 4 : %q\n", strings.SplitN("/home/m_ta/src", "/", 4))  // ["" "home" "m_ta" "src"]
	fmt.Printf("SplitN -1: %q\n", strings.SplitN("/home/m_ta/src", "/", -1)) // ["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))              // ["home" "m_ta,src"]
	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", 2))             // ["" "home#m_ta#src"]
}

//
func findMaxLength() {
	sss := []rune("pwwkew")
	fmt.Println("sli :", sss, len(sss))
	l := len(sss)
	maxLength := 1
	maxS := ""
	for k, _ := range sss {
		keyHash := make(map[rune]bool)
		maxl := 0
		maxStr := make([]rune, 0)
		//fmt.Println("key:", k, " val:", string(v))
		for i := k; i < l; i++ {
			v2 := sss[i]
			if _, ok := keyHash[v2]; ok {
				break
			} else {
				keyHash[v2] = true
				maxl++
				maxStr = append(maxStr, v2)
			}
		}
		fmt.Println("max length:", maxl, " max str:", string(maxStr))
		if maxl > maxLength {
			maxLength = maxl
			maxS = string(maxStr)
		}
	}
	fmt.Println("-------max length:", maxLength, " max str:", string(maxS))
}
