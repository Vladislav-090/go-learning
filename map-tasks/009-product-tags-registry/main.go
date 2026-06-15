package main

import "fmt"

type ProductTagsRegistry interface {
	AddTag(product string, tag string)
	DeleteTag(product string, tag string)
	GetTags(product string) []string
	ViewAllProducts()
	ProductsCount() int
}

type ProductTagsStruct struct {
	Products map[string][]string
}

func (p *ProductTagsStruct) AddTag(product string, tag string) {
	if _, exists := p.Products[product]; !exists {
		p.Products[product] = []string{}
		fmt.Println("New product created with empty slice of tags", product)
	}

	for _, currentTag := range p.Products[product] {
		if currentTag == tag {
			fmt.Println("Tag already exist!", tag)
			return
		}
	}

	p.Products[product] = append(p.Products[product], tag)
	fmt.Println("Tag has been added", product, tag)
}

func (p *ProductTagsStruct) DeleteTag(product string, tag string) {
	tags, exist := p.Products[product]
	if !exist {
		fmt.Println("Product not found!", product)
		return
	}
	for i, currentTag := range tags {
		if currentTag == tag {
			tags = append(tags[:i], tags[i+1:]...)
			p.Products[product] = tags
			fmt.Println("Tag deleted!", tag)
			return
		}
	}

	fmt.Println("Tag not found", tag)
}

func (p *ProductTagsStruct) GetTags(product string) []string {
	return p.Products[product]
}

func (p *ProductTagsStruct) ViewAllProducts() {
	for product, tags := range p.Products {
		fmt.Println("Product:", product, "Tags:", tags)
	}
}

func (p *ProductTagsStruct) ProductsCount() int {
	return len(p.Products)
}

func PrintInfo(p ProductTagsRegistry, product string, tag string) {
	p.ViewAllProducts()
	fmt.Println("Products count:", p.ProductsCount())

	p.AddTag(product, tag)
	fmt.Println("Tags of product:", product, p.GetTags(product))

	p.ViewAllProducts()

	p.DeleteTag(product, tag)
	fmt.Println("Tags of product:", product, p.GetTags(product))

	p.ViewAllProducts()
	fmt.Println("Products count:", p.ProductsCount())
}

func main() {
	productTagsStruct := ProductTagsStruct{
		Products: map[string][]string{
			"Phone": {"Iphone", "Xiomi", "Samsung"},
			"Cloth": {"T-Shirt", "Shirt", "Jeans"},
		},
	}
	PrintInfo(&productTagsStruct, "Phone", "Lenovo")
}
