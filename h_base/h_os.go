package main

import (
	"fmt"
	"os"
)

func testExpand() {
	//Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	s := "hello,world"              //  hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s1 := "${hello},$world $finish" //  hi,kitty finish，finish没有在map规则中，所以还是返回原来的值
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
}

func main() {
	testExpand()
	s := "hello $GOPATH"
	// ExpandEnv根据当前环境变量的值来替换字符串中的${var}或者$var。
	// 如果引用变量没有定义，则用空字符串替换。
	fmt.Println(os.ExpandEnv(s))

	var sep string
	if os.IsPathSeparator('\\') {
		sep = "\\"
	} else {
		sep = "/"
	}
	pwd, _ := os.Getwd()
	path := pwd + sep + "tmp"
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(path)

}
