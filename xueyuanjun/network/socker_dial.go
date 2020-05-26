package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// 几种协议的网络调用方式
	// conn, err := net.Dial("tcp", "localhost:90")
	// // udp
	// conn, err := net.Dial("upd", "localhost:90")
	// icmp链接
	conn, err := net.Dial("ip4:icmp", "www.xueyuanjun.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("链接成功...", conn)
}
