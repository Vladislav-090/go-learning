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

func main() {
	user := User{
		Name: "Vladislav",
		Age:  29,
		Wallets: []Wallet{
			{Address: "a0wsdk2", Balance: 15000},
			{Address: "09x0asdjw", Balance: 199000},
		},
	}

	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}

	fmt.Println(string(jsonData))
}
