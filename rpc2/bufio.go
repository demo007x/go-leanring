package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main () {
	// 加入心跳发生器
	ticker := time.NewTicker(time.Second * 2) // 每 2 秒跳一次
	i := 0
	go func () {
		for {
			i++
			fmt.Printf("heart beat %d \n", i)
			<-ticker.C
		}
	}() // 启动一个立即执行携程
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}