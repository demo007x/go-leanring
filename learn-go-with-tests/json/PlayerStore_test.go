package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// 需要实现server
type StubPlayerStore struct {
	score map[string]int
	winCalls []string
	league []Player
}

// 根据玩家的名字或者得分
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return  s.score[name]
}
// 设置玩家的值+1
func (s *StubPlayerStore) RecordWin(name string) {
	s.score[name]++
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestLeague(t *testing.T)  {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		request,_ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response server '%s', into slice of player ", err)
		}
		fmt.Println(got)
		assertStatus(t, response.Code, http.StatusOK)
		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v, want %v", got, wantedLeague)
		}
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest( )
		response := httptest.NewRecorder()

		if response.Header().Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
		}

		server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t *testing.T, got, want []Player)  {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
// 断言返回的状态是否与预期的一致
func assertStatus(t *testing.T, got, want int)  {
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}