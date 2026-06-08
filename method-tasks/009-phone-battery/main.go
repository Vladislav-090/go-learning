package main

import "fmt"

type Phone struct {
	Brand string
	Battery int
	IsPowered bool
}

func (p *Phone) PowerOn(){
	if p.Battery > 0 {
		p.IsPowered = true
		fmt.Println("Phone has a power")
	} else {
		fmt.Println("Charge the phone first!")
	}
}

func (p *Phone) Use(minutes int){
	if !p.IsPowered {
		fmt.Println("Power on the phone first")
		return
	}
	
	p.Battery -= minutes
		if p.Battery < 0 {
			p.Battery = 0
			p.IsPowered = false
			fmt.Println("Battery is empty, phone turned off...")
			return
		}
		fmt.Println("Used", minutes, "minutes")
	
}

func (p *Phone) Charge(minutes int){
		p.Battery += minutes
	
		if p.Battery > 100 {
			p.Battery = 100
		}
		fmt.Println("Phone charged")
}

func (p Phone) PrintInfo(){
	fmt.Println("Brand:", p.Brand)
	fmt.Println("Battery is: ", p.Battery)
	fmt.Println("Is phone ON?", p.IsPowered)
}


func main() {
	phone := Phone{
		Brand: "Apple",
		Battery: 0,
		IsPowered: false,
	}
	phone.PrintInfo()

	phone.PowerOn()
	phone.Use(20)
	phone.Charge(50)
	phone.PowerOn()
	phone.Use(20)
	phone.PrintInfo()
}