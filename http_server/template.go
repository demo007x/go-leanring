package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse("Hello demo...")
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		err = temp.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Excute: %v", err)
		}
	})
	log.Println("Start http server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
