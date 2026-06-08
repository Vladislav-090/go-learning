package main

import "fmt"

type Product struct {
	Name     string
	Price    int
	Quantity int
}

func (p Product) TotalPrice() int {

	return p.Price * p.Quantity
}

func main() {
	product := Product{
		Name:     "Banana",
		Price:    250,
		Quantity: 4,
	}
	fmt.Println(product.TotalPrice())
}