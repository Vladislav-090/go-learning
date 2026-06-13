package main

import "fmt"

type ProductCatalogService interface {
	UpdateProduct(oldName string, newName string)
	AddProduct(name string)
	DeleteProduct(name string)
	PrintCatalog()
	QuantityOfProduct() int
}

type ProductCatalog struct {
	Products []string
}

func (p *ProductCatalog) AddProduct(name string) {
	for _, products := range p.Products {
		if products == name {
			fmt.Println("Product already exist!!!")
			return
		}
	}
	p.Products = append(p.Products, name)
	fmt.Println("Product has been added!")
}

func (p *ProductCatalog) DeleteProduct(name string) {
	for i, products := range p.Products {
		if products == name {
			p.Products = append(p.Products[:i], p.Products[i+1:]...)
			fmt.Println("Product has been deleted!", name)
			return
		}
	}
	fmt.Println("Product not found!", name)
}

func (p *ProductCatalog) PrintCatalog() {
	for _, products := range p.Products {
		fmt.Println("Product in catalog:", products)
	}
}

func (p *ProductCatalog) UpdateProduct(oldName string, newName string) {
	for i, products := range p.Products {
		if products == oldName {
			p.Products[i] = newName
			fmt.Println("Old updated:", oldName, "->", newName)
			return
		}
	}
	fmt.Println("Product not found!")
}

func (p *ProductCatalog) QuantityOfProduct() int {
	return len(p.Products)
}

func PrintInfo(p ProductCatalogService, name string, oldName string, newName string) {
	p.PrintCatalog()

	p.AddProduct(name)
	p.PrintCatalog()

	p.UpdateProduct(oldName, newName)
	p.PrintCatalog()
	p.DeleteProduct(newName)
	p.PrintCatalog()
	fmt.Println("Products count:", p.QuantityOfProduct())
}

func main() {
	productCatalog := ProductCatalog{
		Products: []string{"apple", "peach", "strawberry", "Blueberry"},
	}
	PrintInfo(&productCatalog, "cherry", "peach", "apricot")
}
