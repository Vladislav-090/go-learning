package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func IsAdult(user User) bool {
	if user.Age >= 18 {
		return true
	} else {
		return false
	}
}

func main() {
	user := User{
		Name: "Vladislav",
		Age:  17,
	}
	fmt.Println("Is adult:", IsAdult(user))

}