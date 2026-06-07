//

package main

import "fmt"


type Product struct {
	Name string
	Price int
}

func ApplyDiscaunt(product *Product, discount int){
	product.Price = product.Price -discount

}

func main() {
	product := Product {
		Name : "Phone",
		Price: 1000,
	}

	ApplyDiscaunt(&product, 200)

	fmt.Println(product.Name,"price is:", product.Price)


	
}