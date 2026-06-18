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
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Wallet Wallet `json:"wallet"`
}

func main() {
	users := User{
		Name: "Vladislav",
		Age:  30,
		Wallet: Wallet{
			Address: "0sadwsadm0x",
			Balance: 100000,
		},
	}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}
	fmt.Println(string(data))
}
