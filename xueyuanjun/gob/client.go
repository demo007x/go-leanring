package main

import (
	"gob/utils"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 30 * time.Second)

	if err != nil {
		log.Fatal("客户端发起连接失败: %v",err)
	}

	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item utils.Item
	client.Call("ServiceHandler.GetName", 1, &item)

	log.Printf("ServiceHandler.GetName 返回结果：%v\n", item)

	var resp utils.Response
	item = utils.Item{
		2, "学院军",
	}

	client.Call("ServiceHandler.SaveName", item, &resp)

	log.Printf("ServiceHandler.SaveName 返回结果：%v\n", resp)
}
