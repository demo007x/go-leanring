package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go SlowOperation(ctx)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			println("goroutine:", runtime.NumGoroutine())
		}
	}()

	time.Sleep(4 * time.Second)
}

func SlowOperation(ctx context.Context) {
	done := make(chan int, 1)
	go func() {
		dur := time.Duration(rand.Intn(5)+1) * time.Second
		fmt.Println(dur)
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-done:
		println("Complete work")
	case <-ctx.Done():
		println("SlowOperation timeout")
	}
}
