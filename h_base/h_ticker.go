package main

import (
	"fmt"
	"time"
)

func main()  {
	//test1()
	done := startTimer(test1)
	//time.Sleep(6 * time.Second)
	//wg.Wait()
	close(done)
	//fmt.Println("close :", time.Now().Format("2006-01-02 15:04:05"))
}

func test1()  {
	//ticker:=time.NewTicker(time.Second*1)
	//ticker.Stop()
	//go func() {
	i := 0
	//	for _ = range ticker.C {
	i++
	fmt.Println("test", i, time.Now().Format("2006-01-02 15:04:05"))
	//}
	//}()

	//time.Sleep(time.Second * 10)
}

func startTimer(f func()) chan struct{} {
	done := make(chan struct{}, 1)
	//wg.Add(1)
	go func() {
		i := 0
		timer := time.NewTicker(1 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				//f()
				i++
				fmt.Println("hehe", i, time.Now().Format("2006-01-02 15:04:05"))
			case <-done:
				fmt.Println("done", time.Now().Format("2006-01-02 15:04:05"))
				//wg.Done()
				return
			}
		}
	}()
	return done
}