package main

import (
	"fmt"
	//"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const N = 26

// 两个协程交替打印1-100的奇数偶数
var POOL = 100

func groutine1(p chan int) {

	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func groutine2(p chan int) {

	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func main() {
	msg := make(chan int)

	go groutine1(msg)
	go groutine2(msg)

	time.Sleep(time.Second * 1)

	return

	//test4()
	return
	//const GOMAXPROCS = 1
	//runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		//fmt.Println("i:", i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("defer1 i: %d %c  \n", i, 'a'+i)
			// A
			//runtime.LockOSThread()
			//runtime.Gosched()
			//time.Sleep(time.Millisecond * 100)
			//fmt.Printf("%c", 'a'+i)
		}(i)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("defer2 i: %d %c  \n", i, 'A'+i)
			//runtime.Gosched()
			//fmt.Printf("%c", 'A'+i)
		}(i)
	}
	wg.Wait()

	return
	fmt.Println("***")
	defer fmt.Println("------:", f())
	fmt.Println("&&&")

	/*fmt.Printf("\n  %c\n", 'A' + 0)
	defer fmt.Println("&&&&&&", f())

	for i := 0; i < 5; i++ {
		defer fmt.Println("%d ", i)
	}
	defer fmt.Println("*****", f())*/
}

func test2() {
	m := make(map[int]*int)

	for i := 0; i < 3; i++ {
		m[i] = &i //A

		//aa := i
		//m[i] = &aa
	}

	for _, v := range m {
		print(*v)

	}

	// output : 333
	// 过改为下面的注释代码则 output：012
}

func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// Golang多个goroutine顺序输出自然数序列
func test3() {
	var number uint32 = 10
	//count相当于一个接力棒
	var count uint32
	trigger := func(i uint32, fn func()) {
		//自旋锁
		for {
			fmt.Println("i:", i)
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				//一定要在执行完函数后才原子加1
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
	}

	for i := uint32(0); i < number; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
		fmt.Println("go routine trigger:", i)
	}
	fmt.Println("number", number)
	//time.Sleep(time.Second)
	trigger(number, func() {})
	//会按照自然数顺序打印（一定是这样）
}

// 多个协程顺序打印数字
func test4() {
	ch := make(chan int)

	for i := 0; i <= 100; i++ {
		go func() {
			fmt.Println("i", <-ch)
		}()
	}

	for i := 0; i <= 100; i++ {
		ch <- i
	}
	close(ch)
	//ch<-200
	time.Sleep(time.Second * 1)

}
