package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse("The value is : {{.}}")
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}
		val := r.URL.Query().Get("val")
		err = temp.Execute(w, val)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
