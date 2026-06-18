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

func main() {
	products := []Product{
		{Name: "Phone", Count: 4, Price: 999.99},
		{Name: "Cloth", Count: 100, Price: 87.99},
		{Name: "Car", Count: 20, Price: 20000.00},
	}

	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}
	fmt.Println(string(data))
}
