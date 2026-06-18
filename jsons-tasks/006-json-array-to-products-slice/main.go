package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

func (p *Product) PrintInfo() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Count:", p.Count)
	fmt.Println("Price:", p.Price)
}

func main() {
	var products []Product

	jsonString := `[{"name":"Phone","count":4,"price":999.99},
	{"name":"Cloth","count":100,"price":87.99},
	{"name":"Car","count":20,"price":20000}
	]`

	err := json.Unmarshal([]byte(jsonString), &products)
	if err != nil {
		fmt.Println("Unmarshal error!", err)
		return
	}

	for _, product := range products {
		product.PrintInfo()
	}

}
