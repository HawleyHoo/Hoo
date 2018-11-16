package h_pkg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
https://www.jianshu.com/p/679041bdaa39

Mutex（互斥锁）
Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
在 Lock() 之前使用 Unlock() 会导致 panic 异常
已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
适用于读写不确定，并且只有一个读或者写的场景
*/
func MutexTs() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var ops int64 = 0
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				fmt.Println("hehe:", state)
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

var waitGroup = new(sync.WaitGroup)

func download(i int) {
	url := fmt.Sprintf("http://pic2016.ytqmx.com:82/2016/0919/41/%d.jpg", i)
	fmt.Printf("开始下载:%s\n", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		fmt.Printf("下载失败:%s", res.Request.URL)
	} else {
		fmt.Printf("开始读取文件内容,url=%s\n", url)
		data, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			fmt.Printf("读取数据失败")
		}
		err3 := ioutil.WriteFile(fmt.Sprintf("pic2018/1_%d.jpg", i), data, 0644)
		if err3 != nil {
			fmt.Println("文件写入失败:", i)
		}
	}
	//计数器-1
	waitGroup.Done()
}

func SyncWaitGroup() {
	os.MkdirAll("/pic2019", 0666)
	now := time.Now()
	for i := 1; i < 24; i++ {
		waitGroup.Add(1)
		go download(i)
	}
	waitGroup.Wait()
	fmt.Printf("下载总时间:%v\n", time.Now().Sub(now))
}
