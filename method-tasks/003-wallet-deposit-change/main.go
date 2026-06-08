package main

import "fmt"

type Wallet struct {
	Owner   string
	Balance int
}

func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
}

func main() {
	wallet := Wallet{
		Owner:   "Mr.Pastukhov",
		Balance: 10000,
	}
	(&wallet).Deposit(2000)
	fmt.Println("now yor balance is", wallet.Balance)
}