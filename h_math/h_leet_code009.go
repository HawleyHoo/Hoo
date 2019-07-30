package main

import (
	"fmt"
	"strconv"
)

/*
num. 93
给一个由数字组成的字符串。求出其可能恢复为的所有IP地址。

样例
给出字符串 "25525511135"，所有可能的IP地址为：

[
"255.255.11.135",
"255.255.111.35"
]

分析
ip地址分为四节，每节取值范围在[0,255]
对于取出的字符串，若它的第一位为’0’，那么它的长度只能为1，多位字符不能以0作为开头

明白这两点，我们可以每次拿出[1,3]个字符，判断是否符合ip地址的条件，
若符合，则递归剩下的字符串，直到四节都符合，那么就得到了一个解，这是分治算法思想
*/

// 暴力穷举法
func RestoreIPAddress(s string) (res []string) {
	for i := 1; i < len(s)-2 && i < 4; i++ {
		for j := i + 1; j < len(s)-1 && j < 4+i; j++ {
			for k := j + 1; k < len(s) && k < 4+j; k++ {
				v1 := s[:i]
				v2 := s[i:j]
				v3 := s[j:k]
				v4 := s[k:]

				if isValid(v1) && isValid(v2) && isValid(v3) && isValid(v4) {
					res = append(res, v1+"."+v2+"."+v3+"."+v4)
				}

			}
		}
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

// 方法二：
//深度搜索，回溯
func RestoreIPAddress2(s string) (res []string) {
	return restore(s, 4)
}

func restore(s string, require int) []string {
	if 1 == require {
		if len(s) > 1 && '0' == s[0] {
			return []string{}
		}
		if v, _ := strconv.Atoi(s); v < 256 {
			return []string{s}
		}
		return []string{}
	}

	var r []string
	for i := 1; i < 4 && i+require-1 <= len(s); i++ {
		prefix := s[:i]
		//fmt.Println("prefix:", prefix)
		if v, _ := strconv.Atoi(prefix); v < 256 {
			for _, j := range restore(s[i:], require-1) {
				r = append(r, prefix+"."+j)
			}
		}
		if '0' == s[0] {
			break
		}
	}
	return r
}

// 方法三
