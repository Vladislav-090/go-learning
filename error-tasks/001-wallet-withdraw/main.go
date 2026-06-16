package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
}

func (w *Wallet) Deposit(amount float64) error {

	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}
	w.Balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	} else if amount > w.Balance {
		return errors.New("Not have enough money!")
	}

	w.Balance -= amount
	return nil
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func main() {
	wallet1 := Wallet{
		Balance: 1000.25,
	}

	err := wallet1.Deposit(500)
	if err != nil {
		fmt.Println("Deposit error:", err)
	} else {
		fmt.Println("Deposit sucsess!")
	}

	err = wallet1.Withdraw(1800)
	if err != nil {
		fmt.Println("Withdraw Error:", err)
	} else {
		fmt.Println("Withdraw sucsess!")
	}

	err = wallet1.Deposit(550)
	if err != nil {
		fmt.Println("Deposit error:", err)
	} else {
		fmt.Println("Deposit sucsess!")
	}

	err = wallet1.Withdraw(1700)
	if err != nil {
		fmt.Println("Withdraw Error:", err)
	} else {
		fmt.Println("Withdraw sucsess!")
	}

	fmt.Println("Final Balance is:", wallet1.GetBalance())

}
