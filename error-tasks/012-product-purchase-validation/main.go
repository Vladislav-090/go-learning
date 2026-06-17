package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ProductName string
	Price       float64
	Stock       int
}

func (p *Product) Buy(quantity int) error {
	if quantity <= 0 {
		return errors.New("Quantity to buy must be positive!")
	}
	if p.Stock == 0 {
		return errors.New("Dont have product in stock")
	}

	if quantity > p.Stock {
		return errors.New("Dont have enough products")
	}

	p.Stock -= quantity
	fmt.Println("Products bought!Now in stok", p.Stock, "products!")
	return nil
}

func (p *Product) Restock(quantity int) error {
	if quantity <= 0 {
		return errors.New("Quantity to restock must be positive")
	}

	p.Stock += quantity
	fmt.Println("The number of items in stock has been updated")
	return nil
}

func (p *Product) PrintInfo() {
	fmt.Printf("Product:%s\nPrice:%.2f\nIn stock:%d\n", p.ProductName, p.Price, p.Stock)
}

func PrintInfo(p *Product) {
	p.PrintInfo()

	err := p.Buy(10)
	if err != nil {
		fmt.Println("Buy error:", err)
	}

	p.PrintInfo()

	err = p.Restock(25)
	if err != nil {
		fmt.Println("Restock err:", err)
	}

	p.PrintInfo()

	err = p.Buy(10)
	if err != nil {
		fmt.Println("Buy error:", err)
	}

	err = p.Buy(0)
	if err != nil {
		fmt.Println("Buy error:", err)
	}

	p.PrintInfo()
}

func main() {
	product := Product{
		ProductName: "Iphone",
		Price:       1000.00,
		Stock:       5,
	}

	PrintInfo(&product)
}
