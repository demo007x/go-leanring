// new(T) 为每一个新的类型T分配一片内存，初始化为0,并返回类型T的指针地址。
//  这种方法返回一个T值为0的地址指针。它适用于值类型： 数组，结构体， 他相当于 &T{}

//  make(T)返回一个类型为T的初始值，塌只适用于内建的引用类型： 切片、map 和channel

package main

import "log"

func main() {
	p := new([]int) //初始化一个切片
	log.Println(*p) //断点1
	b := make([]int, 0)
	log.Println(b) // 断点2
}
