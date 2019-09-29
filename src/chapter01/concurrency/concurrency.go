package main

import (
	"fmt"
	"math/rand"
	"time"
)
// 生产者
func producer(header string, channel chan <- string)  {
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <- chan string)  {
	// 不停的获取数据
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string)
	go producer("cat", channel)
	go producer("dog", channel)

	customer(channel)
}