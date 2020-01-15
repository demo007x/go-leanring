package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	hello := "你好, 世界! hello world"
	debyte := base64.StdEncoding.EncodeToString([]byte(hello))
	fmt.Println("base64 encode", debyte)

	str, err := base64.StdEncoding.DecodeString(string(debyte))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(str))
}
