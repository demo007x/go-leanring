package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//s := "hello, world"
	//
	//fmt.Println(len(s))
	//fmt.Println(s[0], s[len(s) - 1])
	//fmt.Println(s[0:5])
	//fmt.Println(s[5:])
	//fmt.Println(s[:])
	//fmt.Println(s[:6] + "anzi,"+s[6:])

	//s := "left foot"
	//t := s
	//s += ", right foot"
	//
	//fmt.Println(s)
	//fmt.Println(t)

	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

/**
 字符串是否 prefix 开始
**/
func HasPrefix(s, prefix string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

// 字符串是否已 suffer 结尾
func HasSuffix(s, suffix string) bool {
	return len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 是否包含子串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}
