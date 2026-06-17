package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Owner   string
	Balance float64
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}

	w.Balance += amount
	fmt.Println("The wallet was replenished successfully.")
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}

	if amount > w.Balance {
		return errors.New("Dont have enough money!")
	}

	w.Balance -= amount
	fmt.Println("Withdraw successfull")
	return nil
}

func (w *Wallet) Transfer(to *Wallet, amount float64) error {
	if to == nil {
		return errors.New("Wallet not found!")
	}

	if err := w.Withdraw(amount); err != nil {
		return err
	}

	if err := to.Deposit(amount); err != nil {
		return err
	}

	fmt.Println("Transaction succsessfull!")
	return nil
}

func (w *Wallet) PrintInfo() {
	fmt.Printf("Owner: %s | Balance: %.2f\n", w.Owner, w.Balance)
}

func PrintInfo(w *Wallet, to *Wallet) {
	w.PrintInfo()
	to.PrintInfo()

	if err := w.Transfer(to, 250); err != nil {
		fmt.Println("Transfer error:", err)
	}

	if err := w.Transfer(to, 0); err != nil {
		fmt.Println("Transfer error:", err)
	}

	if err := w.Transfer(to, 2000); err != nil {
		fmt.Println("Transfer error:", err)
	}

	if err := w.Transfer(nil, 100); err != nil {
		fmt.Println("Transfer error:", err)
	}

	w.PrintInfo()
	to.PrintInfo()
}

func main() {
	wallets1 := Wallet{
		Owner:   "Vladislav",
		Balance: 1000.00,
	}
	wallets2 := Wallet{
		Owner:   "Viola",
		Balance: 0,
	}
	PrintInfo(&wallets1, &wallets2)
}
