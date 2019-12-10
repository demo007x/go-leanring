package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 玩家的得分
type PlayerStore interface {
	// 根据玩家的名字或者得分
	GetPlayerScore(name string) int
	// 设置玩家的值+1
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	//route *http.ServeMux
	http.Handler
}
// 定义 Player 的结构 [切片]
type Player struct {
	Name string
	Wins int
}
func (p *PlayerServer) leagueHandle(w http.ResponseWriter, r *http.Request) {
	// json 编码
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	//返回一个json结构
	return []Player{
		{"Chris", 20,},
	}
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