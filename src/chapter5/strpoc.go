package main

import (
	"fmt"
	"strings"
)

// 字符串处理
func StringProcess(list []string, chain []func(string) string)  {
	// 遍历每一个字符串
	for index, str := range list {
		result := str
		// 遍历每一个处理函数
		for _, proc := range chain  {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func visit(list []int, f func(int))  {
	for _, v := range list {
		f(v)
	}
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	chain := []func(string) string {
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list, chain)

	for _, str := range list  {
		fmt.Println(str)
	}

	visit([]int{1,2,3,4}, func(v int) {
		fmt.Println(v)
	})
}

