package main

import (
	"fmt"
	"net/http"
)

// 玩家的得分
type PlayerStore interface {
	// 根据玩家的名字或者得分
	GetPlayerScore(name string) int
	// 设置玩家的值+1
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
	//route *http.ServeMux
	http.Handler
}

func (p *PlayerServer) leagueHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showCase(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

// 创建新的playerServer
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandle))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
	p.store.GetPlayerScore(player)
}

func (p *PlayerServer) showCase(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}