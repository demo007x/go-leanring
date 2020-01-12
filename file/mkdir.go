package main

import (
	"fmt"
	"os"
)

func main() {
	baseDir := "/home/test/go/src/go-leanring/file/"
	os.Mkdir(baseDir+"axtaxie", 0777)
	os.MkdirAll(baseDir+"astaxie/test1/test2", 0777)
	err := os.Remove(baseDir + "axtaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll(baseDir + "astaxie/test1/test2")
}
