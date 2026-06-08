package main

import "fmt"

type Car struct{
	Brand string
	Model string
	Mileage int
	IsRunning bool
}

func (c *Car) Start() {
	c.IsRunning = true
	 fmt.Println("Car is started")
} 

func (c *Car) Drive( km int){
	if !c.IsRunning {
		fmt.Println("Please start the car first!")
	} else{
	c.Mileage += km
	fmt.Println("Car driving", km, "km")
	}
}

func (c *Car) Stop() {
	c.IsRunning = false
	fmt.Println("Car is stopped")
}

func main(){
	car:= Car{
		Brand: "Toyota",
		Model: "Camry",
		Mileage: 0,
	}

	car.Drive(50)
	car.Start()
	car.Drive(50)
	car.Stop()
	fmt.Print("Mileage ", car.Mileage)
}