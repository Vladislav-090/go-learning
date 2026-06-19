package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func main() {
	user1 := User{
		Name:     "Vladislav",
		Age:      29,
		Email:    "test@gmail.com",
		Password: "123456",
	}
	user2 := User{
		Name:     "Viola",
		Age:      30,
		Email:    "",
		Password: "123456",
	}

	jsonData, err := json.MarshalIndent(user1, "", "  ")
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}

	fmt.Println(string(jsonData))

	jsonData, err = json.MarshalIndent(user2, "", "  ")
	if err != nil {
		fmt.Println("Marshal error", err)
		return
	}

	fmt.Println(string(jsonData))
}
