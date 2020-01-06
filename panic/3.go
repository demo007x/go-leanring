package main

import (
	"fmt"
	"runtime/debug"
)

func r()  {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
		debug.PrintStack() // 恢复堆栈信息
	}
}

func a()  {
	defer r()
	n := []int{1,2,3}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}