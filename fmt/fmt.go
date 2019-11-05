package main

import "fmt"

type Website struct {
	Name string
}

var site = Website{Name:"test"}

func main() {
	// 相应值的默认格式
	fmt.Printf("%v \n", site)
	// 相应值的go语音表示法
	fmt.Printf("%#v \n", site)
	// 相应值类型的Go语音表示法
	fmt.Printf("%T \n", site)
	// %% 字面上的百分号， 并非值的占位符号
	fmt.Printf("%% \n")

	// 布尔占位符
	fmt.Printf("%t \n", true)
	// 二进制占位符
	fmt.Printf("%b \n", 5)
	// 10精制占位符
	fmt.Printf("%d \n", 10)
	// 八精制占位符
	fmt.Printf("%o \n", 10)
	//  十六精制表示，字符形式为小写 a-f
	fmt.Printf("%x \n", 13)
	// 十六精制表示，字符形式为大写 A-F
	fmt.Printf("%X \n", 13)
	// Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("%U \n", 0x4E2D)
	// 浮点数
	fmt.Printf("%.3f \n", 10.23)
	// 字符串与字节切片
	// 输出字符串表示
	fmt.Printf("%s \n", []byte("go语音中文网"))
	// 双引号围绕的字符串， 以go语音安全的输出
	fmt.Printf("%q \n", "go语音中文网")
	// 16精制小写字面的输出
	fmt.Printf("%x \n", "go语音中文网")
	// 16精制大写字符输出
	fmt.Printf("%X \n", "go语音中文网")
	// 指针的表示
	fmt.Printf("%p", &site)
	// 错误的demo
	fmt.Printf("%d", 132132)
}
