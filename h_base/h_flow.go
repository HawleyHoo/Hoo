package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println("field.name:", p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	/* val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，
	对它所做的任何修改都不会影响到集合中原有的值
	（注：如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值）。*/
	for _, v := range data {
		fmt.Printf("v: %v pointer :%p \n", v, &v)
		go v.print()
	}
	time.Sleep(3 * time.Second)
	return
	/*
		1。 func (p field) print() 则goroutines可能显示 one， two， three
			q:  接收器receiver为值类型即：（p field），接收器是v的值拷贝， 所以会输出one， two， three
			一个值类型的接收器当方法调用时会创建一份拷贝，所以外部的修改不能作用到这个接收器上。

		2。 func (p *field) print()goroutines （可能）显示: three, three, three
			接收器为指针接收器，即：(p *field)
			打印 v 和 &v 发现：
				v: {one} pointer :0xc0420461c0
				v: {two} pointer :0xc0420461c0
				v: {three} pointer :0xc0420461c0
	*/

	fmt.Println("---------------")
	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v.name)
	}
	// goroutines输出: one, two, three
	time.Sleep(3 * time.Second)

	a2 := [24]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 24; i++ {
		a2[i] = i + 1
	}

	step := 3
	for i := 0; i < len(a2); i += step {
		if i+step < len(a2) {
			s1 := a2[i : i+6]
			fmt.Println("s1:", s1)
			go func(s []int) {
				fmt.Println("s:", s)
			}(s1)
		} else {
			s2 := a2[i:]
			fmt.Println("s2:", s2)
		}
	}
	time.Sleep(time.Duration(3) * time.Second)
	/*
		s1: [1 2 3 4 5 6]
		s1: [4 5 6 7 8 9]
		s1: [7 8 9 10 11 12]
		s1: [10 11 12 13 14 15]
		s1: [13 14 15 16 17 18]
		s1: [16 17 18 19 20 21]
		s1: [19 20 21 22 23 24]
		s: [4 5 6 7 8 9]
		s2: [22 23 24]
		s: [19 20 21 22 23 24]
		s: [7 8 9 10 11 12]
		s: [10 11 12 13 14 15]
		s: [13 14 15 16 17 18]
		s: [16 17 18 19 20 21]
		s: [1 2 3 4 5 6]
	*/
}
