package main

import (
	"fmt"
	"sync"
)

var x int
// 用通道模拟锁的问题， 多个协程进来，通道中的长度为1，其他的协程只能等待， 通道工的数据输出了，通道的容量变为了1的时候其他的协程才能使用通道， 这样就避免了竞争的情况。
// 与 sync.Mutext  的原理是一样的，
// 如果通道的容量修改为2， 是什么情况: 最后x的结果同样是1000，
func increment(wg *sync.WaitGroup, ch chan int)  {
	ch <- x
	x++
	fmt.Println("the x's value is ", x)
	go func() {
		fmt.Println(<-ch)
	}()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan int, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
