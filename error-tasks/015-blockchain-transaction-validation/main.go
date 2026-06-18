package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	From   string
	To     string
	Amount float64
	Fee    float64
	Success bool
}

func (t *Transaction) Validate(balance float64) error {
	if t.From == "" {
		return errors.New("Empty 'from' adress")
	}
	if t.To == "" {
		return errors.New("Empty 'to' adress")
	}
	if t.From == t.To {
		return errors.New("Similar adress from <-> to")
	}
	if t.Amount <= 0 {
		return errors.New("Amount must be positive")
	}

	if t.Fee < 0 {
		return errors.New("Fee must be positive")
	}
	if balance < t.Amount+ t.Fee{
		return errors.New("Dont have enough money")
	}

	return nil
}

func (t *Transaction) Execute(balance *float64) error {
	if balance == nil {
		return errors.New("Balance is nil")
	}

	err := t.Validate(*balance)
	 if err !=nil{ 
		return err
	}
	

	*balance -= t.Amount + t.Fee
	t.Success = true
	fmt.Println("Transaction Success")
	return nil
}

func (t *Transaction) PrintInfo() {
	fmt.Printf("Transaction info:\nFrom:%s\nTo:%s\nAmount:%.2f\nFee:%.2f\nSuccess:%t\n",t.From, t.To, t.Amount, t.Fee, t.Success)
}

func PrintInfo(t *Transaction, balance *float64) {
	t.PrintInfo()
	
	err := t.Execute(balance)
	if err != nil {
		fmt.Println("Execute error",err)
	}	
	fmt.Println("Transaction status now:", t.Success)
}

func main() {
	var balance float64 = 20000

	transaction1 := Transaction{
		From: "Vladislav",
		To: "Viola",
		Amount: 12000,
		Fee: 1000,
	}

	transaction2 := Transaction{
		From: "Vladislav",
		To: "Vladislav",
		Amount: 12000,
		Fee: 1000,
	}

	transaction3 := Transaction{
		From: "Vladislav",
		To: "Alex",
		Amount: 25000,
		Fee: 1000,
	}

	transaction4 := Transaction{
		From: "",
		To: "Viola",
		Amount: 12000,
		Fee: 1000,
	}

	PrintInfo(&transaction1, &balance)
	PrintInfo(&transaction2, &balance)
	PrintInfo(&transaction3, &balance)
	PrintInfo(&transaction4, &balance)
}