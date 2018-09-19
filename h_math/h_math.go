package h_math

import (
	"fmt"
	//"time"
)


func fansuanjiecheng(n int) (res int) {
	len := n
	res = 0
	for  i := 2; i < len; i++ {
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

func count_one_bits(n int) (int)  {
	return n;
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
