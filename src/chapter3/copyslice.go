package main

import "fmt"

func main() {
	const elementCount  = 1000

	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copyData := make([]int, elementCount)

	copy(refData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[elementCount - 1])

	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	// 切片删除元素
	seq := []string{"a", "b", "c", "d", "e"}
	index := 2
	fmt.Println(seq[:index], seq[index+1:])
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
