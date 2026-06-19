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
			{Address: "awdksadjwdk", Balance: 12000.00},
			{Address: "0x09xwdskawd", Balance: 20000.00},
		},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error!", err)
		return
	}

	fmt.Println(string(jsonData))

}
