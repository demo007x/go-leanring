package main

import (
	"fmt"
	"sync"
)

func test_put(pool *sync.Pool, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	value := "Hello, demo007"
	pool.Put(value)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello, world!"
		},
	}
	go test_put(pool, wg.Done)
	wg.Wait()
	fmt.Println(pool.Get())
}
