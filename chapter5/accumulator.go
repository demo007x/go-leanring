package main

import "fmt"

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	accumulate := Accumulate(1)

	fmt.Println(accumulate())
	fmt.Println(accumulate())
	fmt.Printf("%p\n", accumulate)
	// 创建一个累加器， 初始值为1
	accumulate2 := Accumulate(10)
	// 累加1并打印地址
	fmt.Println(accumulate2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulate2)
}
