package main

import (
	"fmt"
	"sync/atomic"
)

type Entity struct {
	Key string
	Val interface{}
}

type Noaway struct {
	Movies atomic.Value
	Total  atomic.Value
}

func NewNoaway() *Noaway {
	n := new(Noaway)
	n.Movies.Store(&Entity{"movie", "Wolf Warrior 2"})
	n.Total.Store("$2539306")
	return n
}

func main04062() {
	n := NewNoaway()
	val := n.Movies.Load().(*Entity)
	total := n.Total.Load().(string)

	fmt.Printf("Movies %v domestic total as of 2017: %v \n", val.Val, total)
}
