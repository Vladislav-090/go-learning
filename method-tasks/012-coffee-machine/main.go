package main

import "fmt"

type CoffeeMachine struct {
	Brand string
	IsOn  bool
	Water int
}

func (c *CoffeeMachine) TurnOn() {
	if c.Water == 0 {
		fmt.Println("Add water first")
		return
	} else {
		c.IsOn = true
		fmt.Println("CoffeeMachine ON")
	}
}

func (c *CoffeeMachine) AddWater(amount int) {
	c.Water += amount

	if c.Water > 100 {
		c.Water = 100
	}

	fmt.Println("Water added")
}
func (c *CoffeeMachine) MakeCoffee(amount int) {
	if !c.IsOn {
		fmt.Println("Turn on cofeemachine first")
		return
	}
	
	if c.Water < amount {
		fmt.Println("Not enough water")
		return
	}

	c.Water -= amount
	fmt.Println("Making cofeee")
}

func (c *CoffeeMachine) TurnOff() {
	c.IsOn = false
	fmt.Println("Coffee machine Off")
}

func (c *CoffeeMachine) PrintInfo() {
	fmt.Println("Brand is:", c.Brand)
	fmt.Println("Coffe machine status:", c.IsOn)
	fmt.Println("Amount of water", c.Water)
}

 
func main() {
	coffeemachine := CoffeeMachine{
		Brand: "McCofee",
		IsOn: false,
		Water: 0,
	}
	coffeemachine.PrintInfo()
	coffeemachine.TurnOn()
	coffeemachine.AddWater(50)
	coffeemachine.TurnOn()
	coffeemachine.MakeCoffee(20)
	coffeemachine.MakeCoffee(40)
	coffeemachine.TurnOff()
	coffeemachine.PrintInfo()
}