package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type helloHandler struct{}

func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello demo!"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		log.Println("goruntine")
		// 直到收到了中断信号， 才走下面的代码
		<-quit
		log.Println("goruntine channel is received")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()
	log.Println("server is starting ...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}
