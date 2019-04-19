package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Println(ctx context.Context, a, b int) {
	for {
		fmt.Println(a + b)

		a, b = a+1, b+1
		select {
		case <-ctx.Done():
			fmt.Println("程序结束！")
			return
		default:

		}
	}
}

func main0403() {
	{
		// 超时取消
		a := 1
		b := 2
		timeout := 1 * time.Second
		ctxBg := context.Background()
		ctx, _ := context.WithTimeout(ctxBg, timeout)
		Println(ctx, a, b)

		time.Sleep(1 * time.Second) // 等待时候还会继续输出
	}

	{
		// 手动取消
		a := 1
		b := 2
		ctx, _ := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			//cancelCtx() // 在调用处主动取消
		}()
		Println(ctx, a, b)

		time.Sleep(2 * time.Second)
	}
}

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	time.Sleep(time.Second * 10)
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Println("done！")
			return
		default:
			fmt.Println("work!")
		}
	}
}

func main04061() {
	someHandler()
	fmt.Println("down")
}

// 超时控制与请求跟踪
func ExecCommand(logid string, t *int64, fn func(context.Context, chan *http.Response)) *http.Response {
	timeout := time.Millisecond * time.Duration(20)
	if t != nil {
		timeout = time.Millisecond * time.Duration(*t)
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	r := make(chan *http.Response, 1)

	go fn(context.WithValue(ctx, "logid", logid), r)

	select {
	case <-ctx.Done():
		res, err := http.Get("")
		if err != nil {
			fmt.Println("err:", err)
		}
		return res
	case rr := <-r:
		return rr

	}

	res, err := http.Get("")
	if err != nil {
		fmt.Println("err:", err)
	}
	return res
}

func IsMatch(ctx context.Context) {

}
