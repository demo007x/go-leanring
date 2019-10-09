package main

import (
	"fmt"
)

var any interface{}
var a interface{} = 100
var b interface{} = "hi"
func main() {
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false

	fmt.Println(any)
	fmt.Println(a == b)
}
