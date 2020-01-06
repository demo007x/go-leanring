package main

import "fmt"

func main() {
	a := 5
	func(){
		// 闭包函数
		fmt.Println("a=", a)
	}()
}