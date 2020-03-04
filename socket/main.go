package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)

	// 服务端对客户端的http请求，（升级为websocket协议）进行应答，协议升级为websocket，http建立链接的tcp三次握手将保持
	if conn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	// 启动一个协程, 每秒向客户端发送一次心跳检测
	go func() {
		var err error
		for {
			if err = conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	// 对客户端的数据惊进行处理
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println("receive message", string(data))
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	// 出错之后关闭socket链接
	conn.Close()
}

func main() {
	http.HandleFunc("/", wsHandler)
	http.ListenAndServe(":7777", nil)
}
