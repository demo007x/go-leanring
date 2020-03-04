package main

import (
	"fmt"
	"html/template"
	"im/app/controller"
	"log"
	"net/http"
)

// 注册模板
func registerView() {
	//basePath, _ := os.Getwd()
	tpl, err := template.ParseGlob("./app/view/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		fmt.Println(tplName)
		http.HandleFunc(tplName, func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, tplName, nil)
		})
	}
}

func main() {
	// 静态文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 用户上传的资源文件
	http.Handle("/resource/", http.FileServer(http.Dir(".")))
	registerView()
	// 注册
	http.HandleFunc("/register", controller.UserRegister)
	// 登录
	http.HandleFunc("/login", controller.UserLogin)
	// chat router
	http.HandleFunc("/chat", controller.Chat)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
