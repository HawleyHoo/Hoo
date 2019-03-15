package h_pkg

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func MakeGarbage() {
	pool := make([][]byte, 20)

	var m runtime.MemStats
	makes := 0
	for {
		b := makeBuffer()
		makes += 1
		i := rand.Intn(len(pool))
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		//HeapSys：程序向应用程序申请的内存
		//
		//HeapAlloc：堆上目前分配的内存
		//
		//HeapIdle：堆上目前没有使用的内存
		//
		//HeapReleased：回收到操作系统的内存
		fmt.Printf("HeapSys:%d, HeapAlloc:%d, HeapIdle:%d, HeapReleased:%d, bytes:%d, makes:%d\n",
			m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, bytes, makes)
	}
}

func ManualGC01() {
	pool := make([][]byte, 20)

	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0
	for {
		var b []byte
		select {
		case b = <-buffer:
		default:
			makes += 1
			b = makeBuffer()
		}

		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}

		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}

var makes int
var frees int

type queued struct {
	when  time.Time
	slice []byte
}

func makeRecycler() (get, give chan []byte) {
	get = make(chan []byte)
	give = make(chan []byte)

	go func() {
		q := new(list.List)
		for {
			if q.Len() == 0 {
				q.PushFront(queued{when: time.Now(), slice: makeBuffer()})
			}

			e := q.Front()

			timeout := time.NewTimer(time.Minute)
			select {
			case b := <-give:
				timeout.Stop()
				q.PushFront(queued{when: time.Now(), slice: b})

			case get <- e.Value.(queued).slice:
				timeout.Stop()
				q.Remove(e)

			case <-timeout.C:
				e := q.Front()
				for e != nil {
					n := e.Next()
					if time.Since(e.Value.(queued).when) > time.Minute {
						q.Remove(e)
						e.Value = nil
					}
					e = n
				}
			}
		}

	}()

	return
}

// golang手动管理内存
// https://www.cnblogs.com/luckcs/articles/4107647.html
func ManualGC02() {
	pool := make([][]byte, 20)

	get, give := makeRecycler()

	var m runtime.MemStats
	for {
		b := <-get
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			give <- pool[i]
		}

		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes, frees)
	}
}
