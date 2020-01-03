package main

import "fmt"

func producer(ch chan int)  {
	for n := 0; n<10; n++ {
		ch<-n
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received", v)
	}
}
