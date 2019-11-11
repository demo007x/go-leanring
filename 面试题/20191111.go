package main

import "fmt"

type X struct{}

func (x *X) test() {
	println(x)
}

func main() {
	// 1.下面的代码有什么问题？
	// var a *X
	// a.test()
	// X{}.test()

	// 2.下面代码有什么不规范的地方吗？
	x := map[string]string{"one": "a", "two": "", "three": "c"}
	
	ifv v := x["two"]; v == "" {
		fmt.Println("no entry")
	}
}
