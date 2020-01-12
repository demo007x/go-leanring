package main

import (
	"fmt"
	"os"
)

func main() {
	baseDir := "/home/test/go/src/go-leanring/file/"
	file, err := os.Open(baseDir + "astaxie.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
