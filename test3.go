package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup

	wg.Add(5 * 2)
	for i := 0; i < 5; i++ {
		//wg.Add(2)
		go func(n int) {
			 defer wg.Done()
			//defer wg.Add(-1)
			//EchoNumber(n)
			fmt.Printf(" (2: %d) \n", n)
		}(i)

		go func(n int) {
			defer wg.Done()
			//defer wg.Add(-1)
			fmt.Printf(" (1: %d) \n", n)
		}(i)

		if i == 4 {
			//wg.Add(- 10)
		}
	}

	wg.Wait()
	fmt.Println("------------")
}

func EchoNumber(i int) {
	//time.Sleep(3e9)
	fmt.Println(i)
}