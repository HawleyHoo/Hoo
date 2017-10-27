package h_base

import "fmt"

func test2() {

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {

		defer fmt.Println("defer i = ", i)

		defer func() {

			fmt.Println("defer closure i = ", i)

		}()

		fs[i] = func() {

			fmt.Println("closure i = ", i)

		}
	}

	for _, f := range fs {

		f()
	}
}

func test1() {

	// 普通函数
	/*普通的函数调用，
	每次调用func p时，完成 i 的值复制，然后打印，
	此时 i 值复制了3次，分别是1，2，3。
	由于defer是后进先出，所以执行变成3，2，1
	*/
	a := []int{1, 2, 3}
	for _, i := range a {
		fmt.Println(i)
		defer p(i)
	}

	fmt.Println("--------------------")
	// 闭包函数
	/*
	闭包里的非传递参数外部变量值是传引用的，
	在闭包函数里那个i就是外部非闭包函数自己的参数，
	所以是相当于引用了外部的变量， i 的值执行到第三次是3 ，
	闭包是地址引用所以打印了3次i地址指向的值，所以是3，3，3
	*/
	a1 := []int{1, 2, 3}
	for _, i := range a1 {
		fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

func p(i int) {
	fmt.Println(i)
}