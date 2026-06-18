package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (u *User) PrinInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Email:", u.Email)
}

func main() {
	var user User
	jsonString := `{"name":"Vladislav","age":29,"email":"test@gmail.com"}`

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("JSON error", err)
	}

	user.PrinInfo()
}
