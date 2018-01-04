package h_base

import (
	"sync"
	"fmt"
)

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

const N = 10

func ClosureTest() {
	m := make(map[int]int)
	//它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
	/*
	WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。
	Add:添加或者减少等待goroutine的数量
	Done:相当于Add(-1)
	Wait:执行阻塞，直到所有的WaitGroup数量变成0
	*/
	wg := &sync.WaitGroup{}
	//golang中sync包实现了两种锁Mutex （互斥锁）和RWMutex（读写锁）
	/*
	其中Mutex为互斥锁，Lock()加锁，Unlock()解锁，使用Lock()加锁后，便不能再次对其进行加锁，直到利用Unlock()解锁对其解锁后，才能再次加锁．适用于读写不确定场景，即读写次数没有明显的区别，并且只允许只有一个读或者写的场景，所以该锁叶叫做全局锁．
	func (m *Mutex) Unlock()用于解锁m，如果在使用Unlock()前未加锁，就会引起一个运行错误．
	已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁．
	*/
	mu := &sync.Mutex{}

	wg.Add(N)
	for i := 0; i < N; i++ {
		fmt.Println("i:", i)
		go func(a int) {
			defer wg.Done()
			mu.Lock()
			m[a] = a
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	fmt.Println("len:", len(m), m)
}