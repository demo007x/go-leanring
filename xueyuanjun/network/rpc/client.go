package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/utils"
)

func main() {
	var serverAddress  = "127.0.0.1"
	client, err := rpc.Dial("tcp", serverAddress + ":8080")
	if err != nil {
		log.Fatal("与服务端建立失败:", err)
	}
	args := &utils.Args{
		A: 10, B: 10,
	}

	var reply int
	// 调用远程服务的方法
	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用远程的方式 MathService.Multiply 失败", err)
	}

	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	// 使用go异步方式 调用
	divideCall := client.Go("MathService.Divide", args, &reply, nil)
	for {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}
}