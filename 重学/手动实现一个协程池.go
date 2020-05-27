package main

import (
	"fmt"
	"time"
)

// Pool struct
type Pool struct {
	work chan func()   //  任务
	sem  chan struct{} // 结构体
}

//New 然后定义一个new函数，用于创建协程池对象
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size), // size大小的缓冲池通道
	}
}

// NewTask func
func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

// worker
func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
	}()

	for {
		task()
		task = <-p.work
	}
}

func main() {
	pool := New(2)
	for i := 0; i < 5; i++ {
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}

	// 保证所有的任务都执行完毕
	time.Sleep(5 * time.Second)
}
