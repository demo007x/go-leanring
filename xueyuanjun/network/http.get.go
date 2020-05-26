package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://xueyuanjun.com")

	if err != nil {
		fmt.Printf("请求失败",err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}