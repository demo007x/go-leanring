package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("demo")

	fmt.Printf("%v", string(md5.Sum(data)))
}
