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

func (t *Transaction) PrintInfo() {
	fmt.Println("From:", t.From)
	fmt.Println("To:", t.To)
	fmt.Println("Amount:", t.Amount)
	fmt.Println("Fee:", t.Fee)
	fmt.Println("Status:", t.Status)
}

func main() {
	var transactions []Transaction
	jsonString := `[{"from":"0xsadwdwa","to":"sdjjwusda02","amount":2000,"fee":20,"status":true},
	{"from":"0xsdlwlxl","to":"wlsdjwda02","amount":10000,"fee":100,"status":false},
	{"from":"9xmdjw,sda","to":"ldasjh890x","amount":15000,"fee":1500,"status":true}
	]`

	err := json.Unmarshal([]byte(jsonString), &transactions)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	for _, transaction := range transactions {
		transaction.PrintInfo()
	}
}
