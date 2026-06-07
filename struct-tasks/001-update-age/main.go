package main	

import "fmt"

type User struct {
	Name string
	Age int
}

func PrintUser(user User){
	fmt.Println("hello, iam", user.Name,"and i am", user.Age,"years old")
}

func UpdateAge(user *User, newAge int){
	user.Age = newAge
}

func main() {
	user := User {
		Name: "Vladislav",
		Age: 30,
	}
	UpdateAge(&user, 31)

	PrintUser(user)
}

