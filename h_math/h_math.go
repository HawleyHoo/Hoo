package main

import (
	"fmt"
	//"time"
	"math"
)

// 向上取整 向下取整
func ceilFloor() {
	x := 1.1
	fmt.Println(math.Ceil(x))  // 2
	fmt.Println(math.Floor(x)) // 1

}

// 四舍五入方法
func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func fansuanjiecheng(n int) (res int) {
	len := n
	res = 0
	for i := 2; i < len; i++ {
		remainder := n % i
		if remainder > 0 { // n 不能被整除  则n不是某个数阶乘的乘
			return 0
		}
		fmt.Println("n:", n, "i:", i, "resmainder:", remainder)
		n = n / i
		if n == 1 {
			res = i
			return res
		}
	}
	return res
}

func main() {
	//fmt.Println(120 / 2)
	fmt.Println(fmt.Sprintf("120 阶乘反算为：%d", fansuanjiecheng(120)))
	fmt.Println(fmt.Sprintf("122 阶乘反算为：%d", fansuanjiecheng(122)))
	fmt.Println(fmt.Sprintf("720 阶乘反算为：%d", fansuanjiecheng(720)))
	fmt.Println(fmt.Sprintf("721 阶乘反算为：%d", fansuanjiecheng(721)))
	//fmt.Println("n:", n, "i:", i, "resmainder:", remainder)

	return

	//fmt.Println(fmt.Sprintf("15的阶乘末尾有%d个零", find(15)))
	//fmt.Println(fmt.Sprintf("25的阶乘末尾有%d个零", find(25)))
	//fmt.Println(fmt.Sprintf("50的阶乘末尾有%d个零", find(50)))
	//utils.Trace("查找素数")()
	//start := time.Now()
	//ii := 23
	fmt.Println("10 % 10", 10%10)
	for ii := 8; ii < 23; ii++ {
		//fmt.Println("shuzi", ii)
		fmt.Println(ii, "以内有", findNumberOfContainOne(ii), "个含有1的整数, :", count1(ii))
	}

	fmt.Println(10, "以内有", findNumberOfContainOne(10), "个含有1的整数")
	//findPrimeNumber(1000)
	//duration := time.Since(start)
	//fmt.Println("耗时：", duration)

}

// 设计一个算法，计算出n阶乘中尾部零的个数
func find(n int) (res int) {
	for n > 0 {
		res += n / 5
		n = n / 5
	}
	return
}

// 1~n 里有多少个包含1的整数
func findNumberOfContainOne(n int) (res int) {
	//fmt.Println("n", n)

	offset := 1
	for n > 0 {
		a := n / 10
		b := n % 10
		if b > 0 {
			if a == 0 {
				res += 1
			} else if a == 1 {
				res += a + b + 1
			} else {
				res += a + b + offset
			}
		} else {
			if a > 1 {
				res += a + 1
			} else {
				res += offset + a + 1
			}
		}
		offset *= 10
		n = n / 10
	}
	//res--
	return res
}

func count1(n int) int {
	count := 0
	j := 0
	val := 0
	for i := 1; i <= n; i++ {
		val = i
		for val > 0 {
			j = val % 10
			if j == 1 {
				count++
				break
			}
			val = val / 10
		}
	}
	return count
}

func countDigitOne(n int) int {
	ones := 0
	for m := 1; m <= n; m *= 10 {
		a := n / m
		b := n % m
		if a%10 == 1 {
			ones += (a+8)/10*m + 1*(b+1)
		} else {
			ones += (a+8)/10*m + 0*(b+1)
		}
	}
	return ones
}

//func count_one_bits(n int) (int)  {
//
//}

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
			fmt.Println(n, "以内的素数：", i)
			//fmt.Println(n, "以内最大的素数是：", i)
			//return
		} else {
			//fmt.Println("---", i)
		}
	}
}

func isPrinme(m int) bool {
	if m == 2 {
		return true
	}
	if m%2 == 0 {
		return false
	}
	for index := 3; index*index <= m; index += 2 {
		if m%index == 0 {
			return false
		}
	}
	return true
}

// 大数相加， 从末尾开始相加
func multiAdd(str1, str2 string) (res string)  {
	if len(str1) == 0 && len(str2) == 0 {
		res = "0"
		return
	}

	index1 := len(str1) - 1
	index2 := len(str2) - 1
	left := 0

	for index1 >= 0 && index2 >= 0  {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		index1--
		index2--
		fmt.Println(c1, c2, c3, res)
	}
	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		index1--
	}
	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		index2--
	}
	fmt.Println(index1, index2)

	return
}