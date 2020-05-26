package main

import (
	"fmt"
	"sync"
)

// func addNum
func addNum(a, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	c := a + b
	fmt.Printf("%d + %d = 5d", a, b, c)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go addNum(i, 1, wg.Done)
	}

	wg.Wait()
}
