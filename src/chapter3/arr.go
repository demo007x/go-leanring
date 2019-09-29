package main

import "fmt"

var strList []string
var numList []string
var numListEmpty = []int{} // 已被分配内存, 所以不等于nil
func main() {
	fmt.Println(strList, numList, numListEmpty)

	fmt.Println(len(strList), len(numListEmpty), len(numList))

	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	var a []string
	a[1] = "1"
	a[0] = "1"
	fmt.Println(a == nil, len(a))

	b := make([]string, 1)
	b = append(b, "9")
	b = append(b, "8")
	fmt.Println(len(b), b[0], b[1], b[2])
}



