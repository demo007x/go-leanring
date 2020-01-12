package main

import (
	"html/template"
	"os"
)

type Persion struct {
	UserName string
	Email    string
}

func main() {
	t := template.New("filename example")
	//t, _ = t.ParseFiles("/home/test/go/src/go-leanring/tmp/tmp/1.html")
	t, _ = t.Parse("hello {{.UserName}}! email: {{.Email}}")
	p := Persion{UserName: "anziguoer", Email: "test@admin.com"}
	t.Execute(os.Stdout, p)
}
