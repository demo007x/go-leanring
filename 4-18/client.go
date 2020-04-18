package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// rpc 的客户端
func main() {
	// 连接主机
	tcpAddr := ":1234"

	// 客户端拨号
	client, err := net.Dial("tcp", tcpAddr)
	checkError(err)
	defer client.Close()

	// 调用服务端的注册用户
	client.Call("Users.Registe")

	// 监听用户的输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("you say: ", scanner.Text())
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
