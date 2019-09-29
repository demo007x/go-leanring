package main

import "fmt"

type Data struct {
	complax []int // 测试切片在参数传递中的效果
	instance InnerData // 实例分配的  InnerData
	ptr *InnerData     // 将ptr声明为 InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	// 输出参数成员格式
	fmt.Printf("InFunc value: %+v \n", inFunc)
	// 打印inFunc 的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}

func main() {
	in := Data{
		complax:  []int{1,2,3},
		instance: InnerData{
			5,
		},
		ptr:      &InnerData{1},
	}

	// 输入结构体成员的情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构体指针地址
	fmt.Printf("in prt: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	//输出结构体成员的情况
	fmt.Printf("out value: %+v \n", out)
	// 输出结构体指针地址
	fmt.Printf("out ptr: %p\n", &out)
}

