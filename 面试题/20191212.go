package main

import "fmt"

//1.下面的代码输出什么？
func main() {
	var a []int = nil
	//a = []int{1, 2}
	a, a[0] = []int{1,2}, 9
	fmt.Println(a)
}
//参考答案即解析：运行时错误。知识点：多重赋值。
//
//多重赋值分为两个步骤，有先后顺序：

//2.下面代码中的指针 p 为野指针，因为返回的栈内存在函数结束时会被释放？

//type TimesMatcher struct {
//	base int
//}
//
//func NewTimesMatcher(base int) *TimesMatcher  {
//	return &TimesMatcher{base:base}
//}
//
//func main() {
//	p := NewTimesMatcher(3)
//	fmt.Println(p)
//}
//A. false
//
//B. true

//参考答案及解析：
//A。Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），因此在 Go 语言中返回函数参数或临时变量是安全的。