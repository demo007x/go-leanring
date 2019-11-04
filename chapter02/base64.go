package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away fro keyboard. https://golang.org/"
	// 编码信息
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出编码完成的信息
	fmt.Println(encodeMessage)
	// 解码信息
	data, err := base64.StdEncoding.DecodeString(encodeMessage)
	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
