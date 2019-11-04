// 编写一个非递归的comma函数，运用bytes.Buffer 而不是简单的字符串拼接
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345678955688"))
}

func comma(s string) string {
	var newByte byte = ','
	n := len(s)

	buf := bytes.NewBuffer([]byte{})

	if n <= 3 {
		return s
	}

	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(newByte)
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}


