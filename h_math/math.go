package main

import (
	"math"
)

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < MIN
}

// 一些特殊数值（无穷（Inf）与非数值（NaN））

// 牛顿迭代法 z = z - (z * z - x) / (2 * z)
func SqrtNewton(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtByNewton(x float64) float64 {
	const E = 0.000001
	var z float64 = x
	var k float64 = 0.0
	for ; ; z = z - (z*z-x)/(2*z) {
		if z-k <= E && z-k >= -E {
			return z
		}
		k = z
	}
}

func InvSqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)
	x = math.Float32frombits(i)

	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)

	return 1 / x
}

//  斐波那契数列
// 1.递归    时间复杂度O(2^n)
func Fibonacci(n int) int {
	//defer utils.Trace("Fibonacci1")()
	if n == 1 || n == 0 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//2.循环    时间复杂度O(n)
func Fibonacci2(n int) int {
	//defer utils.Trace("Fibonacci2")()
	f1, f2 := 1, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}

//3.矩阵求解    时间复杂度O(logn)

var f = [2][2]int{{0, 1}, {1, 1}}

func Fibonacci3(n int) int {
	//defer utils.Trace("Fibonacci3")()
	if n == 1 || n == 2 {
		return 1
	}
	return pow(n, f)[1][1]
}

func pow(n int, f [2][2]int) [2][2]int {
	if n == 1 {
		return f
	}

	if n == 2 {
		return fu(f, f)
	}

	if n%2 == 0 { //偶数
		f = pow(n/2, f)
		return fu(f, f)
	} else {
		return fu(pow(n/2, f), pow(n/2+1, f))
	}
}

func fu(f, m [2][2]int) [2][2]int {
	var temp [2][2]int
	temp[0][0] = f[0][0]*m[0][0] + f[0][1]*m[1][0]
	temp[0][1] = f[0][0]*m[0][1] + f[0][1]*m[1][1]
	temp[1][0] = f[1][0]*m[0][0] + f[1][1]*m[1][0]
	temp[1][1] = f[1][0]*m[0][1] + f[1][1]*m[1][1]
	return temp
}
