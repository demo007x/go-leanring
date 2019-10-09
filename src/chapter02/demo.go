package main

import "fmt"

func main() {
	//var cat int = 1
	//var str string = "banana"
	//
	//fmt.Printf("%p, %p", &cat, &str)
	//var house = "Malibu Point 10880, 90264"
	//ptr := &house
	//fmt.Printf("ptr type: %T\n", ptr) // ptr 的类型
	//// ptr 指针的地址
	//fmt.Printf("address: %p \n", ptr)
	//value := *ptr
	//fmt.Printf("value: %s\n", value)

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
