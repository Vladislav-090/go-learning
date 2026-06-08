package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
	fmt.Println("Your balance increased to:", a.Balance)
}

func (a *Account) Withdraw(amount int) {
	if amount > a.Balance {
		fmt.Println("You dont have enough money!")
	} else {
		a.Balance -= amount
		fmt.Println("yor balance decreased to:", a.Balance)
	}
}

func main() {
	account:= Account{
		Owner: "Mr.Pastukhov",
		Balance: 10000,
	}
	account.Deposit(5000)
	account.Withdraw(3000)
}