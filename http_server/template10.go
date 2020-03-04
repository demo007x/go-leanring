package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		teml, err := template.New("test").Parse(`
	{{$name1 := "alice"}}
	name1:{{$name1}}
	{{with true}}
		{{$name1 := "alice2"}}
		{{$name2 := "bob"}}
		name1:{{$name1}}
		name2:{{$name2}}
	{{end}}
	name1 after with: {{$name1}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse err: %v", err)
			return
		}

		err = teml.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Excute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
