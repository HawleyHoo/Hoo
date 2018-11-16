package h_math

import "math"

const MIN = 0.000001

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
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
