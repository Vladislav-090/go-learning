package main

import (
	"errors"
	"fmt"
)

type Room struct {
	Number   string
	IsBooked bool
}

func (r *Room) Book() error {
	if r.IsBooked {
		return errors.New("Room is already booked!")
	}
	r.IsBooked = true
	fmt.Println("Room has been booked!")
	return nil
}

func (r *Room) CancelBooking() error {
	if !r.IsBooked {
		return errors.New("Room is already available to book!")
	}
	r.IsBooked = false
	fmt.Println("Room has been available again!")
	return nil
}

func (r *Room) PrintInfo() {
	fmt.Println("Room:",r.Number, "Status:", r.IsBooked)
}

func PrintInfo(r *Room) {
	r.PrintInfo()

	err := r.Book()
	if err != nil {
		fmt.Println("Book error", err)
	}

	err = r.Book()
	if err != nil {
		fmt.Println("Book error", err)
	}

	err = r.CancelBooking()
	if err != nil {
		fmt.Println("CancleBooking error",err)
	}

	err = r.CancelBooking()
	if err != nil {
		fmt.Println("CancleBooking error",err)
	}

	r.PrintInfo()
}

func main() {
	room := Room{
		Number: "Room №1",
		IsBooked: false,
	}

	PrintInfo(&room)
}