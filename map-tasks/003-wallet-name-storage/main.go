package main

import "fmt"

type WalletNameStorageService interface {
	AddWallet(address string, name string)
	DeleteWallet(address string)
	GetNameByAddress(address string) string
	PrintWallets()
	WalletsCount() int
}

type WalletNameStorage struct {
	Wallet map[string]string
}

func (w *WalletNameStorage) AddWallet(address string, name string) {
	if _, exist := w.Wallet[address]; exist {
		fmt.Println("Address already exist!", address)
		return
	}
	w.Wallet[address] = name
	fmt.Printf("New wallet has been added: [adress:%s , name:%s] \n ", address, name)
}

func (w *WalletNameStorage) DeleteWallet(address string) {
	delete(w.Wallet, address)
	fmt.Printf("Wallet %s deleted \n", address)
}

func (w *WalletNameStorage) GetNameByAddress(address string) string {
	return w.Wallet[address]
}

func (w *WalletNameStorage) PrintWallets() {
	for address, name := range w.Wallet {
		fmt.Printf("Address : %s, Name: %s \n", address, name)
	}
}

func (w *WalletNameStorage) WalletsCount() int {
	return len(w.Wallet)
}

func PrintInfo(w WalletNameStorageService, address string, name string) {
	w.PrintWallets()
	fmt.Println("Quantity of Wallets is:", w.WalletsCount())
	w.AddWallet(address, name)
	fmt.Println("Owner: ", w.GetNameByAddress(address))
	w.PrintWallets()
	w.DeleteWallet(address)
	fmt.Println("Quantity of Wallets is:", w.WalletsCount())
}

func main() {
	wallets := WalletNameStorage{
		Wallet: map[string]string{
			"0dsawd02s": "Viola",
			"09szxwada": "Vladislav",
			"01230asd9": "Afina",
		},
	}
	PrintInfo(&wallets, "09_xsa9123", "Samuel")
}
