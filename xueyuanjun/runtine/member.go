package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.Mutex){
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d \n",counter, a, b, c)
	lock.Unlock()
}

// 通过共享内存的方式来实现协程数据的共享
func main() {
	count := 10
	start := time.Now()
	lock := &sync.Mutex{} // 实例化一个全局的互斥锁，每一个lock对应一个unlock，保证同一个时间只能有一个协程在执行
	for i := 0; i < count; i++ {
		go add(1, i, lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时：", consume)
}