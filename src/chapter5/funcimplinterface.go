package main

import "fmt"

type Invoker interface {
	// 需要实现Call方法
	Call(interface{})
}

type FuncCaller func(interface{})

// 实现 Invoker
func (f FuncCaller) Call(p interface{}) {
	f(p)
}
var invoker Invoker
func main() {
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	invoker.Call("Hello")
}
