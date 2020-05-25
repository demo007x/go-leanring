package main

import (
	"fmt"
	"time"
)

// // 申明一个数组，元素为chan
// var chs [10]chan int

// // 申明一个切片， 值为Chan
// var chs2 []chan int

// // 申明一个字典, 值为chan
// var chs3 map[string]chan int

func test(ch chan<- int) {
	// 往通道中写入数据
	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int, 10) // 初始化一个20个容量的通道
	go test(ch)

	// 读取通道中的数据
	for i := range ch {
		fmt.Println("接收到的数据：", i)
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行时间：(s)", consume)
}
