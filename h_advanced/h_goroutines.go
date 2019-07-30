/*
@Time : 2019-07-14 20:23
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const limit int64 = 10000000000

func main() {

	t0 := time.Now()
	fmt.Println(SerialSum())
	fmt.Println("dur1 :", time.Now().Sub(t0).Nanoseconds())

	t0 = time.Now()
	fmt.Println(ConcurrentSum())
	fmt.Println("dur2 :", time.Now().Sub(t0).Nanoseconds())

	t0 = time.Now()
	fmt.Println(ChannelSum())
	fmt.Println("dur3 :", time.Now().Sub(t0).Nanoseconds())

}

func SerialSum() int64 {
	sum := int64(0)
	for i := int64(0); i < limit; i++ {
		sum += i
	}
	return sum
}

func ConcurrentSum() int64 {
	n := int64(runtime.GOMAXPROCS(0))
	sums := make([]int64, n)
	wg := sync.WaitGroup{}
	wg.Add(int(n))
	for i := int64(0); i < n; i++ {

		go func(i int64) {
			// 将输入分割到每个块
			start := (limit / n) * i
			end := start + (limit / n)
			// 在每个块中运行各自的 loop
			for j := start; j < end; j++ {
				sums[i] += j
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	// 从各个块中收集
	sum := int64(0)
	for _, s := range sums {
		sum += s
	}
	return sum
}

func ChannelSum() int64 {
	n := int64(runtime.GOMAXPROCS(0))

	res := make(chan int64, n)

	for i := int64(0); i < n; i++ {
		go func(i int64, r chan<- int64) {
			// 本地变量取代了全局变量
			sum := int64(0)
			// 采用了分块处理
			start := (limit / n) * i
			end := start + (limit / n)
			// 计算中间值
			for j := start; j < end; j++ {
				sum += j
			}
			// 传递结果
			r <- sum

		}(i, res)
	}

	sum := int64(0)
	// This loop reads n values from the channel. We know exactly how many elements we will receive through the channel ,  hence we need no
	// 读取 n 个值  , n 事先确定
	for i := int64(0); i < n; i++ {
		// 读取值并相加
		// 无值时通道被阻塞，完美的的同步机制  ,
		// 本通道无值等待，直到 读取到所有的 n 个值后才关闭 .
		sum += <-res
	}
	return sum

}
