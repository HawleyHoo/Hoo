package h_pkg

import (
	"runtime"
	"fmt"
)

/*
goroutine是Go并行设计的核心。
goroutine说到底其实就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，
Go语言内部帮你实现了这些goroutine之间的内存共享。
执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
也正因为如此，可同时运行成千上万个并发任务。
goroutine比thread更易用、更高效、更轻便。
*/

func Say(s string)  {
	for i := 0; i < 6; i++  {
		// runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		/*
		默认情况下，调度器仅使用单线程，也就是说只实现了并发。
		想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。
		GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。
		如果n < 1，不会改变当前设置。*/
		runtime.Gosched()
		fmt.Println(s)
	}
}

func Sum(a []int, c chan int)  {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func fibonacci(n int, c chan int)  {
	x, y := 1,1
	for i := 0; i < n; i++  {
		c <- x
		x,y = y, x + y
	}
	close(c)
}