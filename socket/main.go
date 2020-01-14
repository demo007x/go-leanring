package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var rely string
		if err = websocket.Message.Receive(ws, &rely); err != nil {
			fmt.Println("can't receive")
		}

		fmt.Println("Received back from client: " + rely)
		msg := "Received back from client: " + rely
		fmt.Println("Sending to client " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send message..")
			break
		}
	}
}
