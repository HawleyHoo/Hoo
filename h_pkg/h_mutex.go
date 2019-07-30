package main

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

/*
RWMutex的使用主要事项
1、读锁的时候无需等待读锁的结束
2、读锁的时候要等待写锁的结束
3、写锁的时候要等待读锁的结束
4、写锁的时候要等待写锁的结束
RWMutex的四种操作方法
RLock（） //读锁定
RUnlock（） //读解锁

Lock（） //写锁定
Unlock（） //写解锁
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

var m *sync.RWMutex

func main04063() {
	m = new(sync.RWMutex)
	go write(1)
	go read(21)
	go write(3)
	go read(22)
	go write(4)
	go read(23)
	go write(5)
	go read(24)
	go write(6)
	go read(25)
	go write(7)

	time.Sleep(20 * time.Second)
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	var p = 0
	var pr = "read"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)

	}
	m.RUnlock()
	println(i, "read end")
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	var p = 0
	var pr = "write"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)

	}
	m.Unlock()
	println(i, "write end")
}

//---------------------
//作者：番薯粉
//来源：CSDN
//原文：https://blog.csdn.net/u010230794/article/details/78554370
//版权声明：本文为博主原创文章，转载请附上博文链接！

type Model struct {
	//sync.Mutex  // 匿名和不匿名效果一样
	mu sync.Mutex
	A  string
}

func (m *Model) Read() {
	m.mu.Lock()
	defer m.mu.Unlock()
	time.Sleep(time.Second)
	fmt.Println("read:", m.A, time.Now())
}

func (m *Model) Write(v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	time.Sleep(time.Second * 2)
	m.A = v
	fmt.Println("write:", m.A, time.Now())
}

func main() {
	m := Model{A: "hehe"}
	go m.Read()
	go m.Write("haha")
	go m.Read()

	time.Sleep(time.Second * 6)

}
