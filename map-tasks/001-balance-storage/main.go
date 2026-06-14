package main

import "fmt"

type BalanceStorageService interface {
	AddBalance(address string, amount float64)
	DeleteAddress(address string)
	GetBalanceByAddress(address string) float64
	PrintBalances()
	AdressCount() int
}

type BalanceStorage struct {
	Balances map[string]float64
}

func (b *BalanceStorage) AddBalance(address string, amount float64) {
	b.Balances[address] = amount
	fmt.Printf("Balance %f added for %s address \n", amount, address)
}

func (b *BalanceStorage) DeleteAddress(address string) {
	delete(b.Balances, address)
	fmt.Printf("%s address has been deleted \n", address)
}

func (b *BalanceStorage) GetBalanceByAddress(address string) float64 {
	return b.Balances[address]
}

func (b *BalanceStorage) PrintBalances() {
	fmt.Println("Available adress with balance:", b.Balances)
}

func (b *BalanceStorage) AdressCount() int {
	return len(b.Balances)
}

func PrintInfo(b BalanceStorageService, address string, amount float64) {
	b.PrintBalances()
	b.AddBalance(address, amount)
	fmt.Printf("Balance for %s is : %.2f\n", address, b.GetBalanceByAddress(address))
	fmt.Println("Address count is:", b.AdressCount())
	b.PrintBalances()
	b.DeleteAddress(address)
	fmt.Println("Address count is:", b.AdressCount())
	b.PrintBalances()
}

func main() {
	balances := BalanceStorage{
		Balances: map[string]float64{
			"0x_0asdasd90":   1000.09,
			"9x_afsdw09x":    968.13,
			"samds_0x9s0xad": 100.00,
		},
	}
	PrintInfo(&balances, "lolkek9x", 123412.92)
}
