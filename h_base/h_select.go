package h_base

import "fmt"
/*
select
每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，他就执行：其他被忽略
如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
	否则：1.如果有default子句，则执行该句；
		 2.如果没有default子句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值
*/
func fibonacci(c, quit chan int)  {
	x, y := 1, 1
	for  {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Fib()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit<-0
	}()
	fmt.Println(c)
	fmt.Println(quit)
	fibonacci(c, quit)


}



