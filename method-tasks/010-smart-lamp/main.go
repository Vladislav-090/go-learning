package main

import "fmt"

type Lamp struct {
	Brand      string
	IsOn       bool
	Brightness int
}

func (l *Lamp) TurnOn() {
	if  l.IsOn {
		fmt.Println("Lamp is already on!")
	} else {
		l.IsOn = true
		fmt.Println("Lamp turned on")
	}
}

func (l *Lamp) TurnOff() {
	if !l.IsOn {
		fmt.Println("Lamp is already off!")
	} else {
		l.IsOn = false
		fmt.Println("Lamp turned off")
	}
}

func (l *Lamp) IncreaseBrightness(amount int) {
	if !l.IsOn {
		fmt.Println("Turn on the lamp first!")
		return
	}

	l.Brightness += amount

	if l.Brightness > 100 {
		l.Brightness = 100
		} 
		fmt.Println("Brightness increased to:", l.Brightness)
}

func (l Lamp) PrintInfo() {
	fmt.Println(l.Brand)
	fmt.Println(l.IsOn)
	fmt.Println(l.Brightness)
}

func main() {
	lamp := Lamp{
		Brand: "LovelyLight",
		IsOn: true,
		Brightness: 85,
	}
	
lamp.PrintInfo()
lamp.TurnOff()
lamp.IncreaseBrightness(14)
lamp.TurnOn()
lamp.IncreaseBrightness(14)
lamp.PrintInfo()

}