package main

import (
	"errors"
	"fmt"
)

type Car struct {
	Model    string
	IsRented bool
	Fuel     int
}

func (c *Car) RentCar() error {
	if c.IsRented {
		return errors.New("Car is already rented")
	}
	if c.Fuel < 10 {
		return errors.New("Need refuel the car first!")
	}

	c.IsRented = true
	fmt.Println("You Rented car!")
	return nil
}

func (c *Car) ReturnCar(fuel int) error {
	if !c.IsRented {
		return errors.New("Car already returned!")
	}
	if fuel < 0 {
		return errors.New("You cannot drive without fuel!")
	}

	c.IsRented = false
	c.Fuel = fuel
	fmt.Println("Car has been returned")
	return nil
}

func (c *Car) AddFuel(amount int) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}
	if c.Fuel+amount > 100 {
		return errors.New("The gas tank does not hold more than 100 gasoline")
	}
	c.Fuel += amount
	fmt.Println("Fuel has been added!")
	return nil
}

func (c *Car) PrintInfo() {
	fmt.Printf("Model of car: %s , Is rented: %t , Fuel level: %d \n", c.Model, c.IsRented, c.Fuel)
}

func PrintInfo(c *Car) {
	c.PrintInfo()

	if err := c.RentCar(); err != nil {
		fmt.Println("RentCar error!", err)
	}
	if err := c.AddFuel(20); err != nil {
		fmt.Println("AddFuel error:", err)
	}

	if err := c.RentCar(); err != nil {
		fmt.Println("RentCar error:", err)
	}

	if err := c.RentCar(); err != nil {
		fmt.Println("RentCar error:", err)
	}

	if err := c.ReturnCar(8); err != nil {
		fmt.Println("ReturnCar error:", err)
	}

	c.PrintInfo()
}

func main() {
	car := Car{
		Model:    "Toyota Camry",
		IsRented: false,
		Fuel:     5,
	}
	PrintInfo(&car)
}
