package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	begin := make(chan bool)
	ch := make(chan bool, 2000)

	for i := 0; i < 50000; i++ {
		// 负载
		go func() {
			var count int
			load := 100000
			for {
				count++
				if count >= load {
					count = 0
					runtime.Gosched()
				}
			}
		}()
	}

	go func() {
		//runtime.LockOSThread()
		<-begin
		fmt.Println("begin")
		tm := time.Now()
		for i := 0; i < 10000000; i++ {
			<-ch
		}
		fmt.Println("complete", time.Now().Sub(tm))
		os.Exit(0)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for {
				ch <- true
			}
		}()
	}

	fmt.Println("all start")
	begin <- true

	select {}
}
