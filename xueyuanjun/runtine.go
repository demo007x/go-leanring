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
	go add(1, 3)
	time.Sleep(1e9)
}