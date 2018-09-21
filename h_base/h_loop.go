package h_base

import (
	"fmt"
	"runtime"
)

func LoopTest()  {
	//  放在for前面，此例会一直循环下去
	Loop:
	fmt.Println("test")
	for a:=0;a<5;a++{
		fmt.Println(a)
		if a>3{
			goto Loop
		}
	}
}

func LoopTest2()  {
	for a:=0;a<5;a++{
		fmt.Println(a)
		if a>3{
			goto Loop
		}
	}
	Loop:           //放在for后边
	fmt.Println("test")
}

func BreakTest()  {
	//在没有使用loop标签的时候break只是跳出了第一层for循环
	//使用标签后跳出到指定的标签,break只能跳出到之前，如果将Loop标签放在后边则会报错
	//break标签只能用于for循环，跳出后不再执行标签对应的for循环
	Loop:
	for j:=0;j<3;j++{
		fmt.Println("j:", j)
		for a:=0;a<6;a++{
			fmt.Println("---a:", a)
			if a>3 {
				break Loop
			}
		}
	}
}

func continueTest()  {
	for j:=0;j<3;j++{
		fmt.Println("j:", j)
		for a:=0;a<6;a++{
			fmt.Println("---a:", a)
			if a>3{
				break
			}
		}
	}

	// 等效于上面的break
	Loop:
	for j:=0;j<3;j++{
		fmt.Println("j:", j)
		for a:=0;a<6;a++{
			fmt.Println("---a:", a)
			if a>3{
				continue Loop
			}
		}
	}
}

func loop(done chan bool) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
		//runtime.Gosched()  //// 显式地让出CPU时间给其他goroutine

	}
	done <- true
}

func looptest3()  {
	runtime.GOMAXPROCS(2)
	done := make(chan bool)
	go loop(done)
	go loop(done)

	<-done
	<-done
}