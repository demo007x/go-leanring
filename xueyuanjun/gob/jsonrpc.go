package main

import (
	"gob/utils"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义服务端的rpc处理
type ServiceHandler struct {

}

func (service *ServiceHandler) GetName(id int, item *utils.Item) error {
	log.Printf("receive GetName call, id: %d", id)
	item.Id = id
	item.Name = "Demo007"
	return nil
}

// 服务响应
func (service *ServiceHandler) SaveName(item utils.Item, resp *utils.Response) error {
	log.Printf("receive SaveName call, item: %v", item)
	resp.Ok = true
	resp.Id = item.Id
	resp.Message = "存储成功"
	return nil
}

func main() {
	// 初始化rpc服务
	server := rpc.NewServer()

	// 启动服务
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("注册服务处理器失败： %v", err)
	}
	defer listener.Close()

	service := new(ServiceHandler)

	err = server.Register(service)
	if err != nil {
		log.Fatalf("注册服务处理失败： %v", err)
	}
	log.Println("start listen on port localhost:8080")
	// 等待并处理客户端连接
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Fatal("接受客户端连接请求失败: %v", err)
		}
		// 定义日抛编码器， 新建一个rpc编码器，并将该编码绑定给rpc服务处理
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
