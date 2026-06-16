package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Name    string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}

	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive!")
	}

	if b.Balance < amount {
		return errors.New("Dont have enough money")
	}

	b.Balance -= amount
	return nil
}

func (b *BankAccount) PrintInfo() {
	fmt.Println("Account:", b.Name, "Balance:", b.Balance)
}

func (b *BankAccount) Transfer(to *BankAccount, amount float64) error {
	if err := b.Withdraw(amount); err != nil {
		return err
	}
	return to.Deposit(amount)
}

func PrintInfo(from *BankAccount, to *BankAccount, amount float64) {
	from.PrintInfo()
	to.PrintInfo()

	err := from.Transfer(to, amount)
	if err != nil {
		fmt.Println("Transfer error:", err)
	} else {
		fmt.Println("Transfer success")
	}

	from.PrintInfo()
	to.PrintInfo()
}

func main() {
	vladislav := BankAccount{
		Name:    "Vladislav",
		Balance: 1000,
	}

	viola := BankAccount{
		Name:    "Viola",
		Balance: 500,
	}

	PrintInfo(&vladislav, &viola, 300)

	fmt.Println("----------------")

	PrintInfo(&vladislav, &viola, 2000)
}
