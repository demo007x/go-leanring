package controller

import (
	"encoding/json"
	"fmt"
	"im/app/service"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type Node struct {
	Conn *websocket.Conn
	// 并行转串行
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwlocker sync.RWMutex

// 实现聊天功能
func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	// 校验token 是否合法
	isLegal := checkToken(userId, token)
	conn, err := (&websocket.Upgrader{
		HandshakeTimeout: 0,
		ReadBufferSize:   0,
		WriteBufferSize:  0,
		WriteBufferPool:  nil,
		Subprotocols:     nil,
		Error:            nil,
		CheckOrigin: func(r *http.Request) bool {
			return isLegal
		},
		EnableCompression: false,
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 获取socket链接的conn, 设置用户链接信息
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	comIds := concatService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	// 开启协程处理发送逻辑
	go sendproc(node)
	// 开启协程， 处理完成接收逻辑
	go revcproc(node)

	sendMsg(userId, []byte("welcome!"))
}

func checkToken(userId int64, token string) {
	user := service.UserService.Find(userId)
	return user.Token == token
}

// 发送信息逻辑
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// 接收信息逻辑
func revcproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		// toto 对数据进一步处理
		dispatch(data)
		fmt.Printf("recv<=%s", data)
	}
}

func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}

	switch msg.Cmd {
	case CmdSingleMsg:
		sendMsg(msg.Dstid, data)
	case CmdRoomMsg:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case CmdHeart:
		// 校验客户端的心跳
	}
}

func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
