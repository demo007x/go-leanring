package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("test").Funcs(template.FuncMap{
			"join": strings.Join,
		})
		_, err := tmpl.Parse(`
			{{define "list"}}
				{{join . ", "}}
			{{end}}
			Names : {{template "list" .names}}
		`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		names := []string{"Alice", "Bob", "Cindy", "David"}
		fmt.Println(strings.Join(names, ", "))
		err = tmpl.Execute(w, map[string]interface{}{
			"names": names,
		})

		if err != nil {
			fmt.Fprintf(w, "Excute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
