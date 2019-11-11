package main

import (
	"log"
	"net/http"
	"time"
)

type helloHandler struct{}

func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

// func main() {
// 	http.Handle("/", &helloHandler{})
// 	log.Println("Staring HTTP server ...")
// 	log.Fatal(http.ListenAndServe(":4000", nil))
// }

// 服务复用器
// func main() {
// 	log.Println("Stasrting HTTP server ...")
// 	log.Fatal(http.ListenAndServe(":4000", &helloHandler{}))
// }

// htt.Server // 改写

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	mux.Handle("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2)
		w.Write([]byte("test test"))
	})
	server := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Starting HTTP server ...")
	log.Fatal(server.ListenAndServe())
}
