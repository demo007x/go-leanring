// 读写互斥锁
// rwMutex 在读锁的情况下， 有runtime过来的时候，其他的runtime可以读取， 但是写不可以
//  如果 是在写锁的情况下， 其他的runtime过来读取或者写入的情况下， 都将阻塞。

package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

func add(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock() // 添加读写锁
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter
		lock.RUnlock()
		// runtime.Gosched()
		if c >= 10 {
			break
		}
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时（s）:", consume)
}
