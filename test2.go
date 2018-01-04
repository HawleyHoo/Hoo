package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 26

func main() {
	fmt.Println("-------")
	go func() {
		//fmt.Println("-------")
	}()
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		//fmt.Println("i:", i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("defer1 i:%d  \n", i)
			// A
			//fmt.Printf("%c", 'a'+i)
		}(i)

		go func(i int) {
			defer wg.Done()
			fmt.Printf("defer2 i:%d  \n", i)
			//fmt.Printf("%c", 'A'+i)
		}(i)
	}
	wg.Wait()


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


func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}