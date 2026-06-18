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

func main() {
	users := []User{
		{Name: "Vladislav", Age: 29, Email: "test1@gmail.com"},
		{Name: "Viola", Age: 25, Email: "test2@gmail.com"},
		{Name: "Afina", Age: 3, Email: "test3@gmail.com"},
	}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("JSON error", err)
		return
	}
	fmt.Println(string(data))
}
