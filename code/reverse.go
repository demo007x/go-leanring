package main

import "fmt"

func main() {
	var arr [7]int = [7]int{1,2,3,6,48,299,4990}
	reverse(&arr)
	fmt.Println("In main(), arr values:", arr)
}

func reverse(arr *[7]int)  {
	// 分别去两端后赋值
	for i, j := 0, len(*arr) -1; i < j; i,j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

