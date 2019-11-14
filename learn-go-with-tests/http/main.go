package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, score)
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	palyer := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(palyer)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, r)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	p.store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}

type InMemoryPlayStore struct{}

func (i *InMemoryPlayStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayStore{}}

	//handler := http.HandlerFunc(server.ServeHTTP)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
