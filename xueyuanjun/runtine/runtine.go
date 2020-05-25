package main

import (
	"fmt"
	"time"
)

func add(a, b int) {
	var c = a + b

	fmt.Printf("%d + %d = %d \n", a, b, c)
}

func main() {
	count := 10
	for i := 0; i < count; i++ {
		go add(i, i)
	}
	time.Sleep(1e9)
}
