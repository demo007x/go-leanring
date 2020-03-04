package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse(`
		<ul>
		{{range $name, $val := .}}
		<li>{{$name}} - {{$val}}</li>
		{{end}}
		</ul>
`)
		if err != nil {
			fmt.Fprintf(w, "Parse error : %v", err)
			return
		}

		err = temp.Execute(w, map[string]interface{}{
			"Names": []string{
				"Alice",
				"Bob",
				"Carol",
				"David",
			},
			"Numbers": []int{1, 3, 5, 7},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute error: %v", err)
			return
		}
	})
	log.Println("Starting http server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
