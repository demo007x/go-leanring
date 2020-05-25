package main

import (
	"fmt"
	"time"
)

func add (a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1 //将值放入通道
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10) // 创建一个len为10的切片, 切片元素为chan

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(1, i, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	consume := time.Now().Sub(start).Seconds()
	fmt.Println("程序耗时（s）：", consume)
}