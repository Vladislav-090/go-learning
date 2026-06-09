package main

import "fmt"

type Door struct {
	Name     string
	IsOpen   bool
	IsLocked bool
}

func (d *Door) Open() {
	if d.IsLocked {
		fmt.Println("Unlock the door first")
		return
	} 
	if d.IsOpen {
		fmt.Println("Door is already open")
		return
	}	 
	
	d.IsOpen = true
	fmt.Println("Door is open now")
	}

	func (d *Door) Close() {
	if !d.IsOpen {
		fmt.Println("door is already close")
		return
	} 
	d.IsOpen = false
	fmt.Println("Door closed now")
	}

func (d *Door) Lock() {
	if d.IsOpen {
		fmt.Println("Close the door first")
		return
	}
	if d.IsLocked {
		fmt.Println("Door is already locked")
		return
	}
	d.IsLocked = true
	fmt.Println("Door is locked now")
}

func (d *Door) Unlock() {
	if !d.IsLocked {
		fmt.Println("Door is already unlocked")
		return
	}
	d.IsLocked = false
	fmt.Println("Door is unlocked now")
}

func (d *Door) PrintInfo() {
	fmt.Println("Name:", d.Name)
	fmt.Println("Door is open:", d.IsOpen)
	fmt.Println("Door is locked:", d.IsLocked)
	
}


func main() {
	door := Door{
		Name: "Wooden",
		IsOpen: true,
		IsLocked: false,
	}

	door.PrintInfo()
	door.Close()
	door.Lock()
	door.PrintInfo()
	door.Close()
	door.Unlock()
	door.Open()
}
