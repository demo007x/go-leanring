package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	// 从参数中读取主机的信息
	service := os.Args[1]
	// 建立网络连接
	// conn, err := net.Dial("tcp", service)
	// 测试网络超时
	conn, err := net.DialTimeout("tcp", service, 3*time.Second)
	checkError(err)
	// 设置读写超时 5s
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	// 调用返回的连接对象提供的write方法发起请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 通过连接对象的read方法，读取所有的响应数据
	result, err := readFully(conn)
	checkError(err)

	// 打印响应的数据
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应的数据后， 主动关闭连接
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte //  byte的切片

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
