package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.New("test").Funcs(template.FuncMap{
			"add": func(a, b int) int {
				return a + b
			},
		})
		_, err := temp.Parse(`
			result:{{add 1 2 | add 3 | add 4 }}
		`)

		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		err = temp.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Excute: %v", err)
			return
		}
	})

	log.Println("Starting HTTP Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
