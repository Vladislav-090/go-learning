package main

import "fmt"

type PaymentProcessor interface {
	Pay(ammount int)
}

type KaspiPayment struct {
	Account string
}

type HalykPayment struct {
	Account string
}

type CryptoPayment struct {
	Account string
}

func (k *KaspiPayment) Pay(ammount int) {
	fmt.Println("You paid using Kaspi:", ammount)
}

func (h *HalykPayment) Pay(ammount int) {
	fmt.Println("You paid using Halyk:", ammount)
}

func (c *CryptoPayment) Pay(ammount int) {
	fmt.Println("You paid using Crypto:", ammount)
}

func ProcessorPayment(p PaymentProcessor, ammount int) {
	p.Pay(ammount)
}

func main() {
	kaspi := KaspiPayment{Account: "Vladislav"}
	halyk := HalykPayment{Account: "Vladislav"}
	crypto := CryptoPayment{Account: "Vladislav"}
	
	ProcessorPayment(&kaspi, 4000)
	ProcessorPayment(&halyk, 6000)
	ProcessorPayment(&crypto, 7000)
}