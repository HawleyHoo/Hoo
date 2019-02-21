package main

import (
	"fmt"
	"strings"
)

func nonRepeatingSubStr(s string) string {

	last := make(map[rune]int) //该字符上一次出现的位置

	startTemp := 0 //距离当前下标不含重复字符的起始下标

	start := 0 //当前最大不重复子串的起始下标

	maxLength := 0 //当前最大不重复子串的长度

	var res []rune //存储rune子串

	//获取最长不重复子串的起始位置和长度（start,maxLength）

	st := []rune(s) //这步很重要 见解释1

	for i, ch := range st {

		if lastI, ok := last[ch]; ok && lastI >= startTemp {

			startTemp = i

		}

		if i+1-startTemp > maxLength {

			maxLength = i + 1 - startTemp

			start = startTemp

		}

		last[ch] = i

	}

	//拼装成字符串并返回

	var t, j int

	for j = start; j < start+maxLength; j++ {

		temp := st[j]

		res = append(res, temp)

		t++

	}

	result := string(res)

	return result

}

func indexRu() {
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
}

func main() {

	fmt.Println(

		nonRepeatingSubStr("tcytcytcy1"))

	fmt.Println(

		nonRepeatingSubStr("ttttt"))

	fmt.Println(

		nonRepeatingSubStr("2221_22_23"))

	fmt.Println(

		nonRepeatingSubStr(""))

	fmt.Println(

		nonRepeatingSubStr("2"))

	fmt.Println(

		nonRepeatingSubStr("你好,tcy"))

	fmt.Println(

		nonRepeatingSubStr(

			"世界很大，出去看看"))

}
