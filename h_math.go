package main

import "fmt"

func main()  {
	fmt.Println(fmt.Sprintf("15的阶乘末尾有%d个零", find(15)))
	fmt.Println(fmt.Sprintf("25的阶乘末尾有%d个零", find(25)))
	fmt.Println(fmt.Sprintf("50的阶乘末尾有%d个零", find(50)))
}

func find(n int) (res int)  {
	for n > 0  {
		res += n / 5
		n = n / 5
	}
	return
}

// 质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数。
func findPrimeNumber(n int)  {
	
}
