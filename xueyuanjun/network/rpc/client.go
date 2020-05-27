package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/utils"
)

func main() {
	var serverAddress = "localhost:8080"
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("与服务建立失败：", err)
	}
	args := &utils.Args{
		10, 10,
	}
	var reply int

	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)

	divideCall := client.Go("MathService.Divide", args, &reply, nil)

	for {
		select {
		case <-divideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}
}
