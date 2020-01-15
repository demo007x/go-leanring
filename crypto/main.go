package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	h256 := sha256.New()
	io.WriteString(h256, "anziguoer")
	fmt.Printf("anziguoer => sha256: %x \n", h256.Sum(nil))

	h1 := sha1.New()
	io.WriteString(h1, "anziguoer")
	fmt.Printf("anziguoer => sha1: %x \n", h1.Sum(nil))

	m5 := md5.New()
	io.WriteString(m5, "anziguoer")
	fmt.Printf("anziguoer => md5: %x \n", m5.Sum(nil))
}
