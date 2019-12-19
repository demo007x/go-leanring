package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	//fmt.Println(basename("b.c.go"))
	//fmt.Println(basename("abc"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
//func basename(s string) string {
//	for i := len(s) -1; i >= 0; i-- {
//		if s[i] == '/' {
//			s = s[i+1:]
//			break
//		}
//	}
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '.' {
//			s = s[:i]
//			break
//		}
//	}
//	return s
//}
