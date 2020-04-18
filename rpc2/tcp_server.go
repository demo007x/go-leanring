package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err)
		os.Exit(1)
	}
}

func main() {
	arith := new(Arith)
	rpc.Register(arith) // 注册rpc服务

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Printf("accept at %v\n", time.Now())
		conn, err := listener.Accept() // 接受连接
		if err != nil {
			continue
		}
		fmt.Printf("get at %v\n", time.Now())
		rpc.ServeConn(conn)
	}
}
