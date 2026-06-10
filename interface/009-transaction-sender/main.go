package main

import "fmt"

type TransactionSender interface {
	GetName() string
	SendTransaction(amount float64)
	GetLastTxHash() string
}

type Ethereum struct {
	Name        string
	Transaction float64
	LastHash    string
}
type BSC struct {
	Name        string
	Transaction float64
	LastHash    string
}
type Arbitrum struct {
	Name        string
	Transaction float64
	LastHash    string
}

func (e *Ethereum) GetName() string {
	return e.Name
}
func (e *Ethereum) SendTransaction(ammount float64) {
	e.Transaction = ammount
	e.LastHash = "eth_123_0x09123"
	fmt.Println("Transaction has been sent", ammount)
}

func (e *Ethereum) GetLastTxHash() string {
	return e.LastHash
}

func (b *BSC) GetName() string {
	return b.Name
}

func (b *BSC) SendTransaction(amount float64) {
	b.Transaction = amount
	b.LastHash = "bsc_11223_0x09123"
	fmt.Println("Transaction has been sent", amount)
}

func (b *BSC) GetLastTxHash() string {
	return b.LastHash
}
////
func (a *Arbitrum) GetName() string {
	return a.Name
}

func (a *Arbitrum) SendTransaction(ammount float64) {
	a.Transaction = ammount
	a.LastHash = "arb_092x09_fy123"
	fmt.Println("Transaction has been sent", ammount)
}

func (a *Arbitrum) GetLastTxHash() string {
	return a.LastHash
}

func Sender(t TransactionSender, amount float64) {
	fmt.Println("Network:", t.GetName())
	t.SendTransaction(amount)
	fmt.Println("Last hash:", t.GetLastTxHash())
}

func main() {
	ethereum := Ethereum{
		Name: "Ethereum",
		Transaction: 0,
		LastHash: "eth",
	}
	bsc := BSC{
		Name: "BSC",
		Transaction: 0,
		LastHash: "bsc",
	}
	arbitrum := Arbitrum{
		Name: "Arbitrum",
		Transaction: 0,
		LastHash: "arb",
	}
	Sender(&ethereum, 123.05)
	Sender(&bsc, 1000.99)
	Sender(&arbitrum, 98761.56)
}