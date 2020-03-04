package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	request.ParseForm()
	//	name := request.PostForm.Get("name")
	//	fmt.Println(name)
	//	res := `{"name":"`+name+`"}`
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.WriteHeader(http.StatusOK)
	//	writer.Write([]byte(res))
	//})
	//http.ListenAndServe(":8080", nil)

	n := 32
	letterByte := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letterByte[rand.Intn(len(letterByte))]
	}
	fmt.Println(string(b))
}
