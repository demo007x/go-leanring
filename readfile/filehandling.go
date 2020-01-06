package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	f, err := os.Open(*fptr) // 打开文件
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("error reading file:", err)
			break
		}

		fmt.Println(string(b))
	}
}