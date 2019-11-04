package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		name := request.PostForm.Get("name")
		fmt.Println(name)
		res := `{"name":"`+name+`"}`
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(res))
	})
	http.ListenAndServe(":8080", nil)
}

