package main

import (
	"eshop/util"
	"fmt"
	"math"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func main() {
	util.Decimal()

	fmt.Println("hehe:", Round(1.00239834234, 8))
	fmt.Println("hehe:", Round2(1.00239834234, 8))
}
