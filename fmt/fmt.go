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

}
