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

func (u *User) PrintInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Wallet address:", u.Wallet.Address)
	fmt.Println("Wallet balance:", u.Wallet.Balance)
}
func main() {
	var user User
	jsonString := `{"name":"Vladislav","age":30,"wallet":{"address":"0sadwsadm0x","balance":100000}}`

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}
	user.PrintInfo()
}
