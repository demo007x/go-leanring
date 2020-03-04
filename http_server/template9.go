package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("test").Parse(`Inventory
//SKU: {{.Inventory.SKU}}
//Name: {{.Inventory.Name}}
//UnitPrice: {{.Inventory.UnitPrice}}
//Quantity: {{.Inventory.Quantity}}

{{- with .Inventory }}
SKU: {{- .SKU }}
	Name: {{- .Name }}
	UnitPrice: {{- .UnitPrice }}
	Quantity: {{- .Quantity }}
{{- end }}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse: %v", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Inventory": Inventory{
				SKU:       "11000",
				Name:      "Phone",
				UnitPrice: 699.99,
				Quantity:  666,
			},
		})

		if err != nil {
			fmt.Fprintf(w, "Excute error: %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
