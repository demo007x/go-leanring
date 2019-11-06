
package main

import "fmt"

func main()  {
	// 1、未使用的常亮可以编译通过
	// const x = 123
	// const y = 1.23

	// fmt.Println(x)
	// 结论：为使用的常量可以编译通过，常量是一个简单值的表示符号， 在程序运行的时候不会被修改。

	// 2、未初始化的常量的情况
	// const (
	// 	x uint16 = 120
	// 	y 
	// 	c
	// 	s = "abc" 
	// 	z
	// )

	// fmt.Printf("%T %v \n", y, y)
	// fmt.Printf("%T %v \n", z, z)
	// fmt.Printf("%T, %v \n", c, c)
	// 常量组中如果不指定类型和初始化的值，则与上一行的非空常量的右值相同

	// 3、将 nil 分配给 string 类型的变量
	// var x string = nil
	// if x == nil {
	// 	x = "default"
	// }
	// fmt.Println(x)
	// // 解析： 1、nil不能分配给string的变量。 若果只是想想赋值默认值， 可以修改代码如下：
	// var x string 
	// if x == "" {
	// 	x = "default"
	// }

	// fmt.Println(x)
}
