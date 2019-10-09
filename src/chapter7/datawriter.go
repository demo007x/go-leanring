package main

import (
	"fmt"
	"io"
)

type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

// 定义文件结构， 用于实现DataWriter
type file struct {

}

func (d *file) WriteData(data interface{}) error  {
	fmt.Println("WriteData: ", data)
	return nil
}

func (d *file) CanWrite() bool  {
	return true
}

type Socket struct {

}

func (s *Socket) Write(p []byte) (n int,err error)  {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

func usingWriter(writer io.Writer)  {
	writer.Write(nil)
}

func usingCloser(closer io.Closer)  {
	closer.Close()
}

func main() {
	f := new(file)

	var write DataWriter
	write = f
	write.WriteData("data")

	s := new(Socket)
	usingWriter(s)
	usingCloser(s)
}
