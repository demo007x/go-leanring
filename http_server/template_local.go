package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	basePath, _ := os.Getwd()
	// 注册模版
	temp := template.New("demo")
	// 设置模板的解析路径
	temp.ParseGlob(basePath + "/*.html")
	// 注册函数
	temp.Funcs(template.FuncMap{
		"join": strings.Join,
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})

	temp, err := temp.ParseFiles(basePath+"/header.html", basePath+"/footer.html", basePath+"/local.html")
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = temp.ExecuteTemplate(w, "local", map[string]interface{}{
			"names":   []string{"Alice", "Bob", "Cindy", "David"},
			"content": `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}

	})

	log.Println("Starting HTTP server ...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
