package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// 是否具有某个前缀
func HasPrefix(s, prefix string) bool {
	return len(prefix) < len(s) && s[0:len(prefix)] == prefix
}

// 是否以xx为结尾
func HasSuffix(s, suffix string) bool {
	return len(s) > len(suffix) && s[len(s) - len(suffix):] == suffix
}

func Join(str []string, sep string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return str[0]
	}

	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}

	return  buffer.String()
}

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("",""))

	// 统计子串出现的位置
	fmt.Println(strings.Index("anziguoer", "guo"))
	// 统计子串出现的次数
	fmt.Println(strings.Count("aaaa22dd", "a"))
	// 字符串分割
	fmt.Printf("Fields are: %q \n", strings.Fields(" foo bar baz  "))
	fmt.Println(strings.FieldsFunc(" foo bar baz ", unicode.IsSpace))

	// repeat
	fmt.Println("ba", strings.Repeat("na", 2))
	// replace
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
