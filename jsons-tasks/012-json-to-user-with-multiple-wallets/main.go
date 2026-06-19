package main

import (
	"encoding/json"
	"fmt"
)

type Wallet struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Wallets []Wallet `json:"wallets"`
}

func (u *User) PrintInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	for _, wallets := range u.Wallets {
		fmt.Println("Wallet address:", wallets.Address)
		fmt.Println("Wallet balance:", wallets.Balance)

	}
}

func main() {
	var users User
	jsonString := `{"name":"Vladislav","age":29,
	"wallets":[{"address":"awdksadjwdk","balance":12000},
	{"address":"0x09xwdskawd","balance":20000}]}`

	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		fmt.Println("Unmarshal error!", err)
		return
	}
	users.PrintInfo()
}
