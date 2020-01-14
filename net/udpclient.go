package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
	}
	conn, _ := net.DialUDP("udp", nil, udpAddr)
	var buf [512]byte
	n, _ := conn.Read(buf[0:])
	fmt.Println(string(n))
	os.Exit(0)
}
