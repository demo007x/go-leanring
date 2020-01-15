package main

import (
	"html/template"
	"net/http"
	"regexp"
)

type website struct {
	Title string
	Names []string
}

func renderTmpl(w http.ResponseWriter, r *http.Request) {
	wb := website{
		Title: "Hello World",
		Names: []string{"astaxie", "herry", "marry"},
	}
	t, _ := template.ParseFiles("/home/test/go/src/go-leanring/form/form.html")
	t.Execute(w, wb)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	CleanMap := make(map[string]interface{}, 0)
	if name == "astaxie" || name == "herry" || name == "marry" {
		CleanMap["name"] = name
	}

	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]+$", name); ok {
		CleanMap["name"] = name
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("输入的名称不合法"))
	}
}

func handleWhoami(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTmpl(w, r)
	} else if r.Method == "POST" {
		handleForm(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/whoami", handleWhoami)
	http.ListenAndServe(":1234", nil)
}
