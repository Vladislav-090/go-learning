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

func (u *User) PrintInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Email", u.Email)
}

func main() {
	var users []User

	jsonString := `[
  {"name":"Vladislav","age":29,"email":"test1@gmail.com"},
  {"name":"Viola","age":25,"email":"test2@gmail.com"},
  {"name":"Afina","age":3,"email":"test3@gmail.com"}
]`

	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		fmt.Println("Unmarshal error!", err)
		return
	}

	for _, user := range users {
		user.PrintInfo()
		fmt.Println("------------------------")
	}
}
