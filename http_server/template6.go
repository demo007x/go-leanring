package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse(`
			{{if .yIsZero}}
			除数不能为0
			{{else}}
			{{.result}}
			{{end}}
		`)
		if err != nil {
			fmt.Fprintf(w, "Parse error: %v", err)
			return
		}
		x, _ := strconv.ParseInt(r.URL.Query().Get("x"), 10, 64)
		y, _ := strconv.ParseInt(r.URL.Query().Get("y"), 10, 64)

		yIsZero := y == 0
		result := 0.0
		if !yIsZero {
			result = float64(x) / float64(y)
		}
		err = temp.Execute(w, map[string]interface{}{
			"yIsZero": yIsZero,
			"result":  result,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute error: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
