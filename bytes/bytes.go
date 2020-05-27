package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf [10]byte
	result := bytes.NewBuffer(nil)

	for i := 0; i < 10; i++ {
		buf[i] = byte(i)
	}
	result.Write(buf[0:])
	fmt.Println(result.Bytes())
}
