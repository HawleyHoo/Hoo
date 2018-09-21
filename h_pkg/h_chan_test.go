package h_pkg

import (
	"testing"
	"fmt"
)

func TestSay(t *testing.T) {
	//go Say("world")
	//Say("hello")

	a := []int{7,2,6,8,-9,4,0}

	c := make(chan int)
	go Sum(a[:len(a)/2], c)
	go Sum(a[len(a)/2:], c)
	x,y := <-c, <-c
	fmt.Println(x,y,x + y)

	d := make(chan int, 10)
	go fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}
}

func TestSum(t *testing.T) {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func TestChClose(t *testing.T) {
	ChClose()
}

/*
  goroutine test
*/
func TestGoRecover(t *testing.T) {
	//GoRecover()

	//GoChan()

	ChPC()
}