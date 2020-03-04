package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Inventory struct {
	SKU       string
	Name      string
	UnitPrice float64
	Quantity  int64
}

// Subtotal 根据单价和数量计算出总价值
func (i *Inventory) Subtotal() float64 {
	total := i.UnitPrice * float64(i.Quantity)
	fmt.Println(total)
	return total
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.New("test").Parse(`Inventory
SKU: {{.SKU}}
Name: {{.Name}}
UnitPrice: {{.UnitPrice}}
Quantity: {{.Quantity}}
Subtotal: {{.Subtotal}}
`)
		if err != nil {
			fmt.Fprintf(w, "Parse error : %v", err)
			return
		}
		inventory := &Inventory{
			SKU:       r.URL.Query().Get("sku"),
			Name:      r.URL.Query().Get("name"),
			UnitPrice: 0,
			Quantity:  0,
		}
		inventory.UnitPrice, _ = strconv.ParseFloat(r.URL.Query().Get("unitPrice"), 64)
		inventory.Quantity, _ = strconv.ParseInt(r.URL.Query().Get("quantity"), 10, 64)
		err = temp.Execute(w, inventory)
		if err != nil {
			fmt.Fprintf(w, "Execute error %v", err)
			return
		}
	})

	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
