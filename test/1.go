package main

import "fmt"

// 引用类型的demo
func main() {
	a := []int{0, 0, 0}
	a[0] = 10
	fmt.Println(a, len(a), cap(a))
	b := make([]int, 3)
	b[0] = 20
	fmt.Println(b, len(b), cap(b))
	//c := new([]int)
	//c[0] = 30
}
