package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(2 * time.Second) // 停两秒
		ch <- "result 1"
	}()
	fmt.Println("Finished")
}

func main01() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}
