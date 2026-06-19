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

func (u *User) PrintInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Email:", u.Email)
	fmt.Println("Password:", u.Password)
}

func main() {
	var user1 User
	var user2 User
	var user3 User
	jsonString1 := `{
  "name":"Vladislav",
  "age":29,
  "email":"test@gmail.com"
}`

	err := json.Unmarshal([]byte(jsonString1), &user1)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}

	jsonString2 := `{
  "name":"Afina"
}`
	err = json.Unmarshal([]byte(jsonString2), &user2)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}

	jsonString3 := `{
  "name":"Viola",
  "age":30
}`

	err = json.Unmarshal([]byte(jsonString3), &user3)
	if err != nil {
		fmt.Println("Unmarshal error", err)
		return
	}

	user1.PrintInfo()
	user2.PrintInfo()
	user3.PrintInfo()
}
