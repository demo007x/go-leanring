package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse(`
			{{/*打印参数的值*/}}
			Inventory:
			SKU: {{.SKU}}
			Name: {{.Name}}
			UnitPrice: {{.UnitPrice}}
			Quantity: {{.Quantity}}
		`)

		if err != nil {
			fmt.Fprintf(w, "Parse error: %v", err)
			return
		}
		sku := r.URL.Query().Get("sku")
		name := r.URL.Query().Get("name")
		unitPrice, _ := strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		quantity, _ := strconv.ParseInt(r.URL.Query().Get("quantity"), 10, 64)

		err = temp.Execute(w, map[string]interface{}{
			"SKU":       sku,
			"Name":      name,
			"UnitPrice": unitPrice,
			"Quantity":  quantity,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute error: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
