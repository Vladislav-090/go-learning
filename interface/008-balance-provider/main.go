package main

import "fmt"

type BalanceProvider interface {
	GetBalance() float64
	GetName() string
}

type BinanceWallet struct {
	Name string
	Balance float64
}

type EthereumWallet struct {
	Name string
	Balance float64
}

type ArbitrumWallet struct {
	Name string
	Balance float64
}

func (b *BinanceWallet) GetBalance() float64 {
	
	return b.Balance
}
func (b *BinanceWallet) GetName() string {
	return b.Name
}

func (e *EthereumWallet) GetBalance() float64 {
	return e.Balance
}
func (e *EthereumWallet) GetName() string {
	return e.Name
}

func (a *ArbitrumWallet) GetBalance() float64 {
	return a.Balance
}

func (a *ArbitrumWallet) GetName() string {
	return a.Name
}

func PrintBalance(b BalanceProvider) {
	fmt.Println(b.GetName(),"Balance:", b.GetBalance())
}

func main() {
binance := BinanceWallet{
	Name: "BinanceWallet",
	Balance: 1500.90,
}
ethereum := EthereumWallet{
	Name: "EthereumWallet",
	Balance: 100.90,
}
arbitrum := ArbitrumWallet{
	Name: "ArbitrumWallet",
	Balance: 16700.50,
}

PrintBalance(&binance)
PrintBalance(&ethereum)
PrintBalance(&arbitrum)
}