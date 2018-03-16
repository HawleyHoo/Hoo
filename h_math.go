package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(fmt.Sprintf("15的阶乘末尾有%d个零", find(15)))
	//fmt.Println(fmt.Sprintf("25的阶乘末尾有%d个零", find(25)))
	//fmt.Println(fmt.Sprintf("50的阶乘末尾有%d个零", find(50)))
	//utils.Trace("查找素数")()
	start := time.Now()
	findPrimeNumber(2348153532453467788)
	duration := time.Since(start)
	fmt.Println("耗时：", duration)
}

// 设计一个算法，计算出n阶乘中尾部零的个数
func find(n int) (res int) {
	for n > 0 {
		res += n / 5
		n = n / 5
	}
	return
}

// 给出两个整数a和b, 求他们的和, 但不能使用 + 等数学运算符。
func aplusb(a int, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return aplusb(sum, carry)
}
func Add(a, b int) int {
	c := 0
	d := 0
	for {
		c = a ^ b
		d = (a & b) << 1
		a = c
		b = d
		if c == 0 {
			break
		}
	}
	return d
}

// 素数（质数）：只能被1和自身整除的数*
func findPrimeNumber(n int) {
	for i := n; i >= 2; i-- {
		if isPrinme(i) {
			fmt.Println(n, "以内最大的素数是：", i)
			return
		} else {
			//fmt.Println("---", i)
		}
	}
}

func isPrinme(m int) bool {
	if m % 2 == 0 {
		return false
	}
	for index := 3; index * index < m; index += 2 {
		if m % index == 0 {
			return false
		}
	}
	return true
}
