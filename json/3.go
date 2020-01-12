package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{
		ServerName: "Shanghai_vpn",
		ServerIP:   "127.0.0.1",
	})
	s.Servers = append(s.Servers, Server{
		ServerName: "Beijing_vpn",
		ServerIP:   "127.0.0.2",
	})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	// 格式化输出
	//c, err := json.MarshalIndent(s, "", "    ")
	//if err != nil {
	//	fmt.Println("json err:", err)
	//}
	//fmt.Println(string(c))
}
