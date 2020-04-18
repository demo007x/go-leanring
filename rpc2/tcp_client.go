package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/micro/go-micro/v2/util/log"
)

// type Arga of a struct
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "serve")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	client, err := rpc.Dial("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	// 调用乘积的方法
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d*%d = %d\n", args.A, args.B, reply)

	// 调用Divide
	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("artith error:", err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
