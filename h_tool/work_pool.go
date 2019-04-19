package main

import (
	"fmt"
	"strconv"
	"time"
)

type Payload struct {
	name string
}

func (p *Payload) Play() {
	fmt.Println("4 ", p.name, "打LOL游戏... 当前任务已完成。")
	time.Sleep(time.Second)
}

type Worker struct {
	WorkerPool chan (chan Job)
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,     // 任务池
		JobChannel: make(chan Job), // 工作任务
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		// 任务注册到任务池中
		w.WorkerPool <- w.JobChannel

		select {
		case job := <-w.JobChannel:
			fmt.Println("3 GET job:", job.Payload)
			job.Payload.Play()

		case <-w.quit:
			return

		}

	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("2 调度者,接收到一个工作任务", job)

			// 调度者接收到一个工作任务
			go func(job Job) {
				//从现有的对象池中拿出一个
				//jobChannel := <-d.WorkerPool
				//
				//jobChannel <- job
				fmt.Println("6 job:", job, time.Now())
			}(job)
			time.Sleep(time.Duration(1000) * time.Millisecond)
		default:

			fmt.Println("ok!!")
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}

	}
}

func initialize() {
	maxWorkers := 2
	maxQueue := 4
	//初始化一个调试者,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue) //指定任务的队列长度
	//并让它一直接运行
	dispatch.Run()
}

type Job struct {
	Payload Payload
}

var JobQueue chan Job

func main() {
	//初始化对象池
	initialize()

	now0 := time.Now()
	for i := 0; i < 6; i++ {
		p := Payload{
			fmt.Sprintf("玩家-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Payload: p,
		}
		fmt.Println("1 CEO发布一个任务。")
		//time.Sleep(time.Second)
	}
	//close(JobQueue)

	fmt.Println("duration:", time.Now().Sub(now0))
	time.Sleep(time.Second * 10)
}
