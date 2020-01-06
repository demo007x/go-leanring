package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 写文件
	f, err := os.Create("./writefile/test.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	l, err := f.WriteString("Hello World...")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l, "bytes written successfully")
}
