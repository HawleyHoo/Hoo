package main

import (
	"Hoo/h_pkg"
	"fmt"
)

/*
给定一个Excel表格中的列名称，返回其相应的列序号。

示例:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
*/

func convertToIndex(t string) (n int32) {
	r := []rune(t)
	l := len(r)
	//step := 1
	n = 0
	for i := 0; i < l; i++ {
		v := r[i]
		if v < 'A' || v > 'Z' {
			return -1
		}
		n = n*26 + v - 'A' + 1
		//v = v * exponent(26, int32(l - i))
		//fmt.Printf("%d   %d\n", v, exponent(26, int32(l - i)))
	}
	return n
}

func exponent(n, m int32) int32 {
	var res int32 = 1
	m -= 1
	for i := 0; i < int(m); i++ {
		res = res * n
	}
	return res
}

func main() {
	fmt.Println("index:", convertToIndex("AB"))
	fmt.Printf("%d %d  %c \n", 'A', 'Z', 68)
	fmt.Println(convertToTitle(28))

	for i := 1; i < 100; i++ {
		title := convertToTitle(int32(i))
		index := convertToIndex(title)
		fmt.Println("index:", index, " title:", title)
	}
}

func convertToTitle(n int32) (t string) {

	for n > 26 {
		s := n % 26
		n = n / 26
		t += fmt.Sprintf("%c", s+64)
	}
	t += fmt.Sprintf("%c", n+64)
	//if n < 26 {
	//	return fmt.Sprintf("%c", n)
	//}
	//s := n % 26
	return h_pkg.Reverse(t)
}
