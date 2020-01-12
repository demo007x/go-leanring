package main

import (
	"fmt"
	"os"
)

func main() {
	baseDir := "/home/test/go/src/go-leanring/file/"
	userFile := baseDir + "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}
