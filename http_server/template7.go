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
			{{$name1:="alice"}}	
		{{$name2:="bob"}}
		{{$age1:=18}}
		{{$age2:=32}}
		{{if eq $age1 $age2}}
			年龄相同
		{{else}}
			年龄不相同
		{{end}}
		{{if ne $name1 $name2}}
			名字不相同
		{{else}}
			名字相同
		{{end}}
		{{if gt $age1 $age2}}
		alice 年龄比较大
		{{else}}
			bob 年量比较大
		{{end}}
		`)
		err = temp.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, "Execute error: %v", err)
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
