package main

import "fmt"

type User struct {
	Name string
	Age int
}

func (u User) PrintInfo(){
	fmt.Println("my name is", u.Name, "and iam", u.Age)
}

func main() {
	user := User {
		Name: "Vladislav",
		Age: 30,
	}

	user.PrintInfo()
}