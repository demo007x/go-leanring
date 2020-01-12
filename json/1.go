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
	jsonData := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	var s Serverslice
	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		println(err)
		return
	}
	//fmt.Printf("%#v", s)

	for _, v := range s.Servers {
		fmt.Println(v.ServerName, v.ServerIP)
	}
}
