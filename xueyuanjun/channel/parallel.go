package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	count := 10000000
	for i := 0; i < count; i++ {
		sum += i
	}
	ch <- sum // 将值写入通道
	fmt.Printf("子协程%d运行结果:%d\n", seq, sum)
}

func main() {
	// 启动时候
	start := time.Now()
	// 最大的cpu核心数
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)// 协程最大数
	chs := make([]chan int, cpus)

	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}

	sum := 0

	for _, ch := range chs {
		sum += <-ch
	}
	// 结束时间
	end := time.Now()
	fmt.Printf("最终运算结果: %d, 执行耗时(s): %f\n", sum, end.Sub(start).Seconds())
}