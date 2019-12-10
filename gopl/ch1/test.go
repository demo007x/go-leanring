package main

import "fmt"

func main() {
	x := 1
	p := &x

	fmt.Println(*p) // 取 p的值

	*p = 10
	fmt.Println(x)
	fmt.Println(p)
}
