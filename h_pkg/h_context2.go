/*
@Time : 2019-07-18 17:55
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go mainTask(ctx, "1")
	//go mainTask(ctx, "2")
	//go mainTask(ctx, "3")

	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("main exit...")
	time.Sleep(3 * time.Second)
}

func PrintTask(ctx context.Context, taskName string) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("task:", taskName, " exit...")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("task:", taskName, " doing something...")
		}

	}

}

func mainTask(ctx context.Context, taskName string) {

	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()

	// create a new task
	newTaskName := taskName + " ctx"
	go PrintTask(ctx1, newTaskName)

	for {

		select {

		case <-ctx.Done():
			fmt.Println("main task:", taskName, " exit...")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("main task:", taskName, " doing something...")
		}

	}

}
