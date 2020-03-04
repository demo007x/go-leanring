package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	basePath, _ := os.Getwd()
	templatesDir := basePath + "/tmpl/"
	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	widgets, err := filepath.Glob(templatesDir + "widgets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, layout := range layouts {
		files := append(widgets, layout)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

// func render
func render(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		panic("template is not found...")
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}

func handle(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"Title": "My title", "Body": "Hi this is my body"}
	err := render(w, "local.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
