package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains
	str := "anzilonger"
	fmt.Println(strings.Contains(str, "longer"))
	fmt.Println(strings.Contains(str, "guoer"))
	// join
	slice := []string{"an", "zi", "guo", "er"}
	fmt.Println(strings.Join(slice, ""))
	// index
	fmt.Println(strings.Index(str, "long"))
	fmt.Println(strings.Index(str, "guo"))
	// repeat
	fmt.Println(strings.Repeat("a", 4))
	fmt.Println(strings.Replace(str, "long", "guo", 1))
	// split
	newStr := "anzi, guoer"
	fmt.Println(strings.Split(newStr, ", "))
	// trim
	fmt.Println(strings.Trim("!!!Hello world.!!!", "!"))
	fmt.Println(strings.TrimLeft("!!!Hello World.!", "!"))
	fmt.Println(strings.TrimRight("!!!Hello World.!!", "!"))
	fmt.Println("TrimSpace", strings.TrimSpace("    Hello World"))
	fmt.Println("TrimPrefix", strings.TrimPrefix("Hello World", "H"))
	fmt.Println(strings.TrimPrefix("!!Hello World!!", "!!"))
	fmt.Println(strings.TrimFunc("abcde", func(r rune) bool {
		if string(r) == "c" {
			return true
		}
		return false
	}))
	fmt.Println("HasPrefix", strings.HasPrefix("abced", "0"))
	fmt.Println("HasPrefix", strings.HasPrefix("abced", "ab"))

	fmt.Println("字符串转换")

}
