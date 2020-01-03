package main

import (
	"fmt"
	"time"
)

func hello(done chan bool)  {
	fmt.Println("hello go runtine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello world goroutine")
	done<-true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go roruntine")
	go hello(done)
	<-done
	fmt.Println("main function")
}
