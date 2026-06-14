package main

import "fmt"

type ProductPriceService interface {
	AddProduct(name string, price int)
	DeleteProduct(name string)
	GetPriceOfProduct(name string) int
	UpdatePrice(name string, newPrice int)
	PrintStorage()
	GetCountOfProducts() int
}

type ProductPriceStorage struct {
	Products map[string]int
}

func (p *ProductPriceStorage) AddProduct(name string, price int) {
	if _, exist := p.Products[name]; exist {
		fmt.Println("This product already exist!", name)
		return
	}
	p.Products[name] = price
	fmt.Println("Product added :", name, price)
}

func (p *ProductPriceStorage) DeleteProduct(name string) {
	if _, exist := p.Products[name]; exist {
		delete(p.Products, name)
		fmt.Println("The product deleted!", name)
		return
	}
	fmt.Println("Product not found!", name)
}

func (p *ProductPriceStorage) GetPriceOfProduct(name string) int {
	return p.Products[name]
}

func (p *ProductPriceStorage) UpdatePrice(name string, newPrice int) {
	if _, exists := p.Products[name]; exists {
		p.Products[name] = newPrice
		fmt.Printf("New price for %s is: %d \n", name, newPrice)
		return
	}
	fmt.Println("Product not found!")
}

func (p *ProductPriceStorage) PrintStorage() {
	for name, price := range p.Products {
		fmt.Printf("Product name: %s , price : %d \n", name, price)
	}
}

func (p *ProductPriceStorage) GetCountOfProducts() int {
	return len(p.Products)
}

func PrintInfo(p ProductPriceService, name string, price int, newPrice int) {
	fmt.Println("Products count:", p.GetCountOfProducts())

	p.AddProduct(name, price)

	fmt.Printf("Price for %s is: %d\n", name, p.GetPriceOfProduct(name))

	p.PrintStorage()

	p.UpdatePrice(name, newPrice)

	p.PrintStorage()

	p.DeleteProduct(name)

	fmt.Println("Products count:", p.GetCountOfProducts())

	p.PrintStorage()
}

func main() {
	products := ProductPriceStorage{
		Products: map[string]int{
			"apple": 100,
			"phone": 325000,
			"water": 300,
		},
	}
	PrintInfo(&products, "table", 20000, 19000)
}
