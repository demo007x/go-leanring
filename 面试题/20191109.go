package main

func main() {
	// 1、以下代码输出什么
	// var ch chan int

	// select {
	// case v, ok := <-ch:
	// 	println(v, ok)
	// default:
	// 	println("default")
	// }
	// 解答：参考答案及解析：default。ch 为 nil，读写都会阻塞。
}
