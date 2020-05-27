package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/utils"
)

type MathService struct {

}

func (m *MathService) Multiply(args utils.Args, reply *int) error {
	*reply = args.A * args.B

	return nil
}

func (m *MathService) Divide(args utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}

	*reply = args.A / args.B

	return nil
}

func main() {
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()
	// 启动服务
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动http服务失败：", err)
	}
}
