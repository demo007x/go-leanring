package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rep http.ResponseWriter, req *http.Request) {
		params := req.URL.Query() // 获取query参数
		fmt.Fprintf(rep, "你好, %s", params.Get("name"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("启动 http服务器失败： %v", err)
	}
}
