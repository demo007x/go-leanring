package main

import (
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	data, err := ReadFrom(os.Stdin, 11)
	data, err = ReadFrom(&os.File{},9)
	data, err = ReadFrom(strings.NewReader("form string"), 12)
}
