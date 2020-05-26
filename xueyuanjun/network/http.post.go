package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://www.vipidea.dd/api/v1/login", url.Values{"username": {"18611439826"}, "password": {"111111"}})

	if err != nil {
		log.Fatal(err)
	}

	var buf []byte

	defer resp.Body.Close()
	n, err := resp.Body.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	log.Println(string(buf))
	fmt.Println(resp.StatusCode)
}
