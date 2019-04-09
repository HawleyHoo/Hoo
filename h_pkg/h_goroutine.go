package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func NumOfCPU() {
	num := runtime.NumCPU() // 本地机器的逻辑CPU数
	runtime.GOMAXPROCS(num) // 设置可同时执行的最大CPU数
	fmt.Println(num)
}

var domainSyncChan = make(chan int, 10)

/* 如果某个goroutine panic了，而且这个goroutine里面没有捕获(recover)，那么整个进程就会挂掉。
所以，好的习惯是每当go产生一个goroutine，就需要写下recover。 */
func goroutineOfRecover(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error to chan put")
		}
		domainSyncChan <- num
	}()
	panic("error ... ")
}

func GoRecover() {
	for i := 0; i < 10; i++ {
		domainName := i
		go goroutineOfRecover(domainName)
	}
	time.Sleep(time.Second * 2)

	userChan := make(chan interface{})
	go func() {
		for {
			userChan <- "nick"
			err := recover()
			if err != nil {
				fmt.Println("err:", err)
			}
		}
	}()
	for {
		name := <-userChan
		fmt.Println("name:", name)
		time.Sleep(time.Second)
	}
}

/*
 goroutine + channel 栗子
 多个goroutine处理任务，等待一组channel的返回结果
*/
func calc(taskChan, resChan chan int, exitChan chan bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}()

	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}

	exitChan <- true
}

func GoChan() {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 1000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	// 启动8个goroutine做任务
	for i := 0; i < 8; i++ {
		go calc(intChan, resChan, exitChan)
	}

	go func() {
		// 等所有goroutine结束
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
		close(exitChan)
	}()

	for v := range resChan {
		fmt.Println(v)
	}

}

/**
等待一组channel的返回结果 sync.WaitGroup的解决方案
waitgroup 用于等待一组线程的结束，父线程调用add方法来设定等地啊的线程数量，
每个被等待的线程在结束时应调用Done方法，同时主线程里可调用wait方法阻塞至所有线程结束
*/

func ChanMerge(cs <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	//for c := range cs {
	//	go output(c)
	//}
	go output(cs)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
