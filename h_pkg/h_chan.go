package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
goroutine是Go并行设计的核心。
goroutine说到底其实就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，
Go语言内部帮你实现了这些goroutine之间的内存共享。
执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
也正因为如此，可同时运行成千上万个并发任务。
goroutine比thread更易用、更高效、更轻便。
*/

func Say(s string) {
	for i := 0; i < 6; i++ {
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

func Sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/* channel  栗子*/
func sendData(ch chan<- string) {
	ch <- "go"
	ch <- "java"
	ch <- "c"
	ch <- "c++"
	ch <- "python"
	ch <- "PHP"
	ch <- "OC"
	ch <- "swift"
	close(ch)
}

func getData(ch <-chan string, chClose chan bool) {
	for {
		str, ok := <-ch
		if !ok {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(str)
	}
	chClose <- true
}

func ChClose() {
	ch := make(chan string, 10)
	chclose := make(chan bool, 1)

	go sendData(ch)
	go getData(ch, chclose)

	<-chclose
	close(chclose)

}

/* eg */
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

func ChPC() {
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)

	time.Sleep(time.Second)
}

func main() {
	ch := make(chan int, 1)
	SafeSend(ch, 2)
	SafeReceive(ch)
	SafeClose(ch)
	SafeClose(ch)
	SafeReceive(ch)
	fmt.Println("--------")
}

// 安全的关闭/接受/发送 channel
func SafeClose(ch chan int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()

	close(ch)
	return true

}

func SafeSend(ch chan int, value int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

func SafeReceive(ch chan int) {
	i, ok := <-ch
	if ok {
		fmt.Println("read:", i)
	} else {
		fmt.Println("channel closed")
	}
}

// 一个发送者，N个接受者
func OneToN() {
	notify := make(chan int)

	datach := make(chan int, 100)
	go func() {
		<-notify

		for i := 0; i < 100; i++ {
			datach <- i
		}
		fmt.Println("close datach")
		close(datach)
	}()

	//time.Sleep(time.Second * 2)
	fmt.Println("开始发送信息")
	notify <- 1
	time.Sleep(time.Second * 1)
	fmt.Println("3秒后接收到数据管道数据 此时datach 在接收端已经关闭")

	for i := 0; i < 5; i++ {
		go func(n int) {
			for {
				if v, ok := <-datach; ok {
					fmt.Println(n, "read datach", v)
				} else {
					fmt.Println("break")
					break
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 5)

}

// N个发送者，1个接受者
func NToOne() {
	datach := make(chan int, 1)
	stopch := make(chan int)

	for i := 0; i < 10000; i++ {

		go func(i int) {
			for {
				//value := rand.Intn(10000)
				select {
				case <-stopch:
					fmt.Println("接受到停止发送的信号")
					return
				case datach <- i:
					fmt.Println("write")

				}
			}

		}(i)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("1秒后开始接收数据")
	for {
		if v, ok := <-datach; ok {
			fmt.Println("read datach", v)
			if v == 9999 {
				fmt.Println("接收端接收到9999时 告诉发送端不要再发了")
				//close(stopch)
				stopch <- 1
				return
			}
		} else {
			fmt.Println("break")
			break
		}
	}

}

// N个发送者，M个接受者
func NToM() {
	datach := make(chan int, 10)
	toStop := make(chan string)
	stopch := make(chan int)

	// 简约版调度器
	go func() {
		if t, ok := <-toStop; ok {
			fmt.Println("to stop", t)
			close(stopch)
		}
	}()

	// 生产者
	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				val := rand.Intn(100)
				if val == 99 {
					select {
					case toStop <- "sender id:" + strconv.Itoa(i) + "to close":
					default:

					}
				}

				select {
				case <-stopch:
					return
				default:

				}

				select {
				case <-stopch:
					fmt.Println("生产者 stopch return")
					return
				case datach <- val:
					fmt.Println("p", val)

				}

			}
		}(i)
	}

	// 消费者
	for i := 0; i < 20; i++ {
		go func(i int) {
			select {
			case <-stopch:
				fmt.Println("消费者 1 stopch return")
				return
			default:

			}

			select {
			case <-stopch:
				fmt.Println("消费者 stopch return")
				return
			case val := <-datach:
				if val == 99 {
					select {
					case toStop <- "receiver id:" + strconv.Itoa(i) + "to close":
					default:

					}
				}
				fmt.Println("c", val)
			}

		}(i)
	}

	time.Sleep(time.Second * 10)
}
