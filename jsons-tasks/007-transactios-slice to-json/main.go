package main

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Fee    float64 `json:"fee"`
	Status bool    `json:"status"`
}

func main() {
	transactions := []Transaction{
		{From: "0xsadwdwa", To: "sdjjwusda02", Amount: 2000, Fee: 20, Status: true},
		{From: "0xsdlwlxl", To: "wlsdjwda02", Amount: 10000, Fee: 100, Status: false},
		{From: "9xmdjw,sda", To: "ldasjh890x", Amount: 15000, Fee: 1500, Status: true},
	}

	data, err := json.Marshal(transactions)
	if err != nil {
		fmt.Println("Marshal error!", err)
		return
	}

	fmt.Println(string(data))
}
