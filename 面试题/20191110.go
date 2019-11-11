package main

import "fmt"

func main() {
	// 1、以下代码正确的输出是什么？
	var fn1 = func() {

	}

	var fn2 = func() {

	}

	// if fn1 != fn2 {
	// 	println("fn1 not equal fn2")
	// }
	// 答案： 编译错误， 函数不能比较， 函数只能与nil比较

	if fn1 != nil {
		println("fn1 not equal nil")
	}

	if fn2 != nil {
		println("fn2 not equal nil")
	}
	// 以上两个可以正常输出

	// 2 、下面的代码输出什么？
	type T struct {
		n int
	}

	// m := make(map[int]T)
	// m[0].n = 1
	// fmt.Println(m[0].n)
	// 编译错误：cannot assign to struct field m[0].n in map
	// 修改后的代码：
	m := make(map[int]T)
	t := T{1}
	m[0] = t
	fmt.Println(m[0].n)
}
