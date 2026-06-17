package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID     int
	Amount float64
	Payed  bool
}

func (o *Order) Pay(balance *float64) error {
	if o.Payed == true {
		return errors.New("Already payed!")
	}

	if o.Amount <= 0 {
		return errors.New("Order cant be for free")
	}

	if balance == nil {
		return errors.New("Balance is nil!")
	}

	if *balance < o.Amount {
		return errors.New("Not enough money!")
	}

	*balance -= o.Amount
	o.Payed = true
	fmt.Println("Order payed! Balance now is:", *balance)
	return nil
}

func (o *Order) Cancel() error {
	if o.Payed == false {
		return errors.New("Order not pyaed yet")
	}

	o.Payed = false
	fmt.Println("Order canceled")
	return nil
}

func (o *Order) PrintInfo() {
	fmt.Println("Order info: ID:", o.ID, "Amount:", o.Amount, "Status:", o.Payed)
}

func PrintInfo(o *Order, balance *float64) {
	o.PrintInfo()

	err := o.Pay(balance)
	if err != nil{
		fmt.Println("Pay error : ", err)
	}

	err = o.Pay(balance)
	if err != nil{
		fmt.Println("Pay error : ", err)
	}

	err = o.Cancel()
	if err != nil{
		fmt.Println("Cancel error : ", err)
	}

	err = o.Cancel()
	if err != nil{
		fmt.Println("Cancel error : ", err)
	}

	o.PrintInfo()
}

func main() {
	var balance float64 = 1000.00

	order1 := Order{
		ID: 1,
		Amount: 600,
		Payed: false,
	}
	order2 := Order{
		ID: 2,
		Amount: 1200,
		Payed: false,
	}
	PrintInfo(&order1, &balance )
	PrintInfo(&order2, &balance )
}