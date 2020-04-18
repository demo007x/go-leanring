package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

// User type of struct
type User struct {
	ID      int    // 系统分配的用户的 id
	Name    string // 用户的名称
	Message string // 用户发送的信息
}

// Say func of user
func (u *User) Say(msg string) {
	u.Message = msg
	// @todo 用户发信的时候, 把用户的信息广播出去
	fmt.Println(msg)
}

// Users 所有的用户 map
var Users map[int]User

func main() {
	// users := make(map[int]User)
	user := new(User)
	rpc.Register(user) // rpc 注册
	// rpc.Register(users) // rpc 注册

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	// 监听客户端, 如果有客户端连接进来, 则在 users 中就添加一个user 的信息
	for {
		conn, err := listener.Accept() // 接收连接
		if err != nil {
			continue
		}
		fmt.Println("新用户进来了:", time.Now())
		defer conn.Close()
		rpc.ServeConn(conn)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
