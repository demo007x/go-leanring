package main

import (
	"fmt"
	"log"
	"net/http"
)
func PlayerServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL.Path)
	player := r.URL.Path[len("/players/"):]
	switch player {
	case "Pepper":
		fmt.Fprint(w, "20")
	case "Floyd":
		fmt.Fprint(w, "10")
	}
}
func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
