package main

import "fmt"

const size = 300
func main() {
	a := make([]int, 3)
	a[0] = 0
	a[1] = 1
	a[2] = 2

	str := "hello world"
	fmt.Println(str[6:])
	fmt.Printf("type of a is %T \n", a)
	fmt.Printf("type of str is %T \n", str)
}
