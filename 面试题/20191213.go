package main

import "fmt"

//1.下面这段代码输出什么？求相同的个数
//func main() {
//	count := 0
//	for i := range [256]struct{}{} {
//		m, n := byte(i), int8(i)
//		if n == -n {
//			fmt.Println(byte(i), int8(i))
//			count++
//		}
//		if m == -m {
//			fmt.Println(byte(i), int8(i))
//			count++
//		}
//	}
//	fmt.Println(count) // 4个
//}

const (
	azero = iota
	aone = iota
)

const (
	info = "msg"
	bzero = iota
	bone = iota
)

func main() {
	fmt.Println(azero, aone) //0, 1
	fmt.Println(bzero, bone) // 1, 2
}