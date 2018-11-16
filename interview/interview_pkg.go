package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"
)

const N = 100

func main() {

	m := &map[int]int{} //A
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			mu.Lock()
			(*m)[i] = i //B
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)

	str := "HelloWord"
	l1 := len([]rune(str))
	l2 := bytes.Count([]byte(str), nil)
	l3 := strings.Count(str, "")
	l4 := utf8.RuneCountInString(str)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
}
