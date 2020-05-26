package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func addNum(a *int32, b int, deferFunc func()) {
	defer func ()  {
		deferFunc()
	}

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(a)
		newNum := curNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number 当前的值： %d[%d-%d]\n", *a, b,i)
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", b, i)
		}
	}
}

func main() {
	count := 10
	var num int32
	fmt.Printf("number 初始值为%d", num)
	fmt.Println("启动子协程...")
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < count; i++ {
		go addNum(&num, i, func(){
			if atomic.LoadInt32(&num) == int32(count) {
				cancelFunc()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("所有子协程执行完毕...")
}
