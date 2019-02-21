package h_math

import (
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

func testmain() {
	fmt.Println("index:", convertToIndex("EF"))
	fmt.Printf("%d %d  %c \n", 'A', 'Z', 68)
	fmt.Println(convertToTitle(28))

	for i := 1; i < 60; i++ {
		title := convertToTitle2(i)
		//index := convertToIndex(title)
		fmt.Println("index:", i, " title:", title)
	}
}

// letterOnlyMapF is used in conjunction with strings.Map to return only the
// characters A-Z and a-z in a string.
func letterOnlyMapF(rune rune) rune {
	switch {
	case 'A' <= rune && rune <= 'Z':
		return rune
	case 'a' <= rune && rune <= 'z':
		return rune - 32
	}
	return -1
}

// intOnlyMapF is used in conjunction with strings.Map to return only the
// numeric portions of a string.
func intOnlyMapF(rune rune) rune {
	if rune >= 48 && rune < 58 {
		return rune
	}
	return -1
}

/*
public String convertToTitle(int n) {
	String temp="";
	  while(n>0) {
		  char s=(char) ((n-1)%26+'A');
		  temp=s+temp;
		  n=(n-1)/26;
	  }
	return temp;
}

*/
func convertToTitle2(n int) (t string) {
	for n > 0 {
		t = fmt.Sprintf("%c", (n-1)%26+'A') + t
		n = (n - 1) / 26
	}
	//t = fmt.Sprintf("%c", n+'A') + t
	//return h_pkg.Reverse(t)
	return t
}

func convertToTitle(n int32) (t string) {
	for n > 26 {
		s := n % 26
		n = n / 26
		t += fmt.Sprintf("%c", s+'A')
	}
	t = fmt.Sprintf("%c", n+'A') + t
	//return h_pkg.Reverse(t)
	return t
}
