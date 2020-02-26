package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//r.ParseForm()
		var data map[string]interface{} = map[string]interface{}{}
		if r.Header.Get("Content-Type") == "application/json" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data["json_data"] = string(reqBody)
		} else {
			r.ParseMultipartForm(32 << 20)
			data["form"] = r.Form
			data["post_form"] = r.PostForm
		}
		fmt.Fprintln(w, data)
	})

	log.Fatal(http.ListenAndServe(":2020", nil))
}
