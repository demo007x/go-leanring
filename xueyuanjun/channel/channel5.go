package main

import (
	"fmt"
	"time"
)

// 超时处理机制实现
func main() {
	// 初始化通道
	ch := make(chan int, 1)

	// 初始化timeout 通道
	timeout := make(chan bool, 1)
	// 戳先一个匿名等待函数
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	// 借助timeout通道结合 select 语句实现ch通道的读取超时效果
	select {
	case <-ch:
		fmt.Println("接受通道 ch 的数据")
	case <-timeout:
		fmt.Println("超时1秒，程序退出")
	}
}
