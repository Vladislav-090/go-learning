package main

import (
	"errors"
	"fmt"
)

type User struct {
	UserName string
	Password string
	IsActive bool
}

func (u *User) Login(username string, password string) error {
	if u.IsActive == false {
		return errors.New("User is not active!")
	}

	if username == "" {
		return errors.New("Empty login!")
	}

	if username != u.UserName {
		return errors.New("Incorrect login!")
	}

	if password == "" {
		return errors.New("Empty password")
	}

	if password != u.Password {
		return errors.New("Incorrect password")
	}

	fmt.Println("Successfull login")
	return nil
}

func (u *User) PrintInfo() {
	fmt.Printf("Login: %s\n Password: %s\n Status: %t\n", u.UserName, u.Password, u.IsActive)
}

func PrintInfo(u *User, username string, password string) {
	u.PrintInfo()

	err := u.Login("", password)
	if err != nil {
		fmt.Println("Incorrect data", err)
	}

	err = u.Login(username, "")
	if err != nil {
		fmt.Println("Incorrect data", err)
	}

	err = u.Login("wrong username", password)
	if err != nil {
		fmt.Println("Incorrect data", err)
	}

	err = u.Login(username, "wrong password")
	if err != nil {
		fmt.Println("Incorrect data", err)
	}

	err = u.Login(username, password)
	if err != nil {
		fmt.Println("Incorrect data", err)
	} else {
		fmt.Println(" Login check passed")
	}
}

func main() {
	user1 := User{
		UserName: "Vladislav",
		Password: "123",
		IsActive: true,
	}
	user2 := User{
		UserName: "Viola",
		Password: "321",
		IsActive: false,
	}

	PrintInfo(&user1, "Vladislav", "123")

	fmt.Println("--------------------------")

	PrintInfo(&user2, "Viola", "321")

}
