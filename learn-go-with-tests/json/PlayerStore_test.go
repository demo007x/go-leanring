package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 需要实现server
type StubPlayerStore struct {
	score map[string]int
}

// 根据玩家的名字或者得分
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return  s.score[name]
}
// 设置玩家的值+1
func (s *StubPlayerStore) RecordWin(name string) {
	s.score[name]++
}

func TestLeague(t *testing.T)  {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request,_ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
	})
}

// 断言返回的状态是否与预期的一致
func assertStatus(t *testing.T, got, want int)  {
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}