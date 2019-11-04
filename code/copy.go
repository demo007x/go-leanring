package main

import "fmt"

const elementCount  = 1000

func main() {
	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	// 应用切片类型
	refData := srcData
	copyData := make([]int, elementCount)

	// 将数据复制到新的切片中
	copy(copyData, srcData)

	// 修改原始数据的第一个元素
	srcData[0] = 999

	// 打印应用切片的第一个元素
	fmt.Println("应用切片的第一个元素", refData[0])
	fmt.Println(" // 打印复制切片的第一个和最后一个元素", copyData[0], copyData[elementCount-1])
	// 复制元素数据中4到6（不包含）
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d \n", copyData[i])
	}
}
