package main

import "fmt"

type Wallet struct {
	Owner   string
	Balance int
}

func (w *Wallet) Withdraw(amount int) {
	if amount <= w.Balance {
		w.Balance -= amount
		fmt.Println("now your balance is:", w.Balance)
	} else {
		fmt.Println("not enough money")
	}
		}
func main(){
	wallet := Wallet{
		Owner: "Mr.Pastukhov",
		Balance: 10000,
	}
	wallet.Withdraw(10001)
	
}