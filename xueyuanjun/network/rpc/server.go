package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/utils"
)

type MathService struct {
}

func (m *MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0 ")
	}
	*reply = args.A / args.B

	return nil
}

func main() {
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()
	// 服务端监听端口
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("启动服务监听失败:", err)
	}

	// 启动服务
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动http服务失败:", err)
	}

	fmt.Println("启动成功...")
}
