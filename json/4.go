package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ID          int    `json:"-"`
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2, string"`
	ServerIP    string `json:"serverIP, omitempty"`
}

//type Serverslice struct {
//	Servers []Server `json:"servers"`
//}

func main() {
	s := Server{
		ID:          1,
		ServerName:  `go "1.0"`,
		ServerName2: `go "1.0"`,
		ServerIP:    ``,
	}

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)
}
