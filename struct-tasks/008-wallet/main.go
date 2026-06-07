package main 

import "fmt"

type Wallet struct {
	Owner string
	Balance int
}

func SpendMoney(w *Wallet, amount int)  {
	if w.Balance < amount {
		 fmt.Println("u dont have enough money, your balance stay", w.Balance)
		 return 
		
	}
	w.Balance -= amount
	fmt.Println("Now your balance is:", w.Balance)
	
}
	
func main() {
	wallet := Wallet{
		Owner: "Bob",
		Balance: 100,
	}

	SpendMoney(&wallet, 120)
	

	


}