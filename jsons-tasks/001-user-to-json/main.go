package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) PrintJson()error {
	data,err := json.Marshal(u)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	user := User {
		Name: "Vladislav",
		Age: 29,
	}

	err := user.PrintJson()
	if err != nil{
		fmt.Println("JSON error",err)
	}

}
