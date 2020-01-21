package main

import (
	"fmt"
	"sync"
)

func main01() {
	c := make(chan bool, 100)
	iSlice := []int{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			iSlice = append(iSlice, i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
