package main

import "fmt"

type WalletRegistryService interface {
	ViewWallets()
	AddWallet(name string)
	DeleteWallet(name string)
	GetQuantity() int
}

type WalletRegistry struct {
	Wallets []string
}

func (w *WalletRegistry) ViewWallets() {
	for _, wallets := range w.Wallets {
		fmt.Println("Wallet name is:", wallets)
	}
}

func (w *WalletRegistry) AddWallet(name string) {
	for _, wallets := range w.Wallets {
		if wallets == name {
			fmt.Println("Wallet already exist!!!")
			return
		}
	} 
	w.Wallets = append(w.Wallets, name)
	fmt.Println("New Wallet has been added",name )
}

func (w * WalletRegistry) DeleteWallet(name string) {
	for i, wallets := range w.Wallets {
		if wallets == name {
		w.Wallets = append(w.Wallets[:i], w.Wallets[i+1:]... )
		fmt.Println("Wallet has been deleted!", name)
		return
		}
	}
	fmt.Println("Wallet not found", name)
}

func (w * WalletRegistry) GetQuantity()int {
	return len(w.Wallets)
}


func PrintInfo( w WalletRegistryService, name string) {
	w.ViewWallets()
	fmt.Println("Quantity of Wallets:", w.GetQuantity())
	w.AddWallet(name)
	w.ViewWallets()
	fmt.Println("Quantity of Wallets:", w.GetQuantity())
	w.DeleteWallet(name)
	w.ViewWallets()
	fmt.Println("Quantity of Wallets:", w.GetQuantity())
} 

func main() {
	walletRegistry := WalletRegistry{
		Wallets: []string {"Woopay", "Paypal", "Qiwi", "KaspiPay"},
	}
	PrintInfo(&walletRegistry, "Qiwi")
}