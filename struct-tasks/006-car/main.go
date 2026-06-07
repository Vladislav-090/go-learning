package main

import "fmt"


type Car struct {
	Brand string
	Model string
	Year int
}

func UpdateCarsYear(c *Car, newYear int) {
	c.Year = newYear
}

func main() {
	car := Car{
		Brand: "Toyota",
		Model: "Camry",
		Year: 2022,}

	UpdateCarsYear(&car, 2025)
	fmt.Println("Brand", car.Brand)
	fmt.Println("Model", car.Model)
	fmt.Println("Year", car.Year)


}