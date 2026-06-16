package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Stock int
}

func (p *Product) AddStock(amount int) error {
	if amount == 0 {
		return errors.New("Amout must be positive!")
	}
	p.Stock += amount
	return nil
}

func (p *Product) RemoveStock(amount int) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}
	if p.Stock < amount {
		return errors.New("Dont have any stock!")
	}

	p.Stock -= amount
	return nil
}

func (p *Product) UpdateStock(amount int) error {
	if amount < 0 {
		return errors.New("Stock cant be negative")
	}
	p.Stock = amount
	return nil
}

func (p *Product) PrintInfo() {
	fmt.Println("Product:", p.Name, "Count:", p.Stock)
}

func PrintInfo(p *Product) {
	p.PrintInfo()
	err := p.AddStock(5)
	if err != nil {
		fmt.Println("AddStock error", err)
	}
	p.PrintInfo()

	err = p.RemoveStock(20)
	if err != nil {
		fmt.Println("RemoveStock error", err)
	}

	err = p.RemoveStock(6)
	if err != nil {
		fmt.Println("RemoveStock error", err)
	}
	p.PrintInfo()

	err = p.UpdateStock(50)
	if err != nil {
		fmt.Println("UpdateStock error", err)
	}

	p.PrintInfo()
}

func main() {
	product := Product{
		Name:  "Iphone",
		Stock: 10,
	}

	PrintInfo(&product)
}
