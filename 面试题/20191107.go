package main
import "fmt"

func main() {
	// 1、一下代码有什么问题
	// data := []int{1,2,3}
	// i := 0
	// ++i // 这里不能不能使用前置的++ 操作
	// // ++i
	// fmt.Println(data[i++]) // 切片的下标不能使用变量的计算

	// 2、 一下代码的最后一行输出什么
	x := 1
	fmt.Println(x)

	{
		fmt.Println(x)
		i,x := 2,2
		fmt.Println(i, x)
	}

	fmt.Println(x) // 输出的结果是1.
	//  是不是可以理解到 {} 里面的x是局部的变量？
}