package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch1<-1
	ch2 := make(chan int, 1)
	ch2<-2

	select {
	case k1 := <-ch1:
		fmt.Println(k1)
	case k2 := <-ch2:
		fmt.Println(k2)
	default:
		fmt.Println("chan")
	}
}
