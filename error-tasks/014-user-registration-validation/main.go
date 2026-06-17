package main

import (
	"errors"
	"fmt"
	"strings"
)

type UserRegistration struct {
	UserName string
	Email    string
	Password string
}

func (u *UserRegistration) Registrer() error {
	if u.UserName == "" {
		return errors.New("User Name is empty")
	} else if len(u.UserName) < 3 {
		return errors.New("User name cannot be shorter then 3 symbols")
	}

	if u.Email == "" {
		return errors.New("Email is empty")
	} else if !strings.Contains(u.Email, "@") {
		return errors.New("Email must contain @")
	}

	if u.Password == "" {
		return errors.New("Password is empty")
	} else if len(u.Password) < 6 {
		return errors.New("Passwort cannot be shorter then 6 symbols")
	}

	fmt.Println("Registration has been successfull")
	return nil
}

func (u *UserRegistration) PrintInfo() {
	fmt.Printf("UserName: %s\nEmail:%s\nPassword:%s\n", u.UserName, u.Email, u.Password)
}

func PrintInfo(u *UserRegistration) {
	u.PrintInfo()

	err := u.Registrer()
	if err != nil {
		fmt.Println("Register error", err)
	}
}

func main() {
	user1 := UserRegistration{
		UserName: "",
		Email:    "test@gmail.com",
		Password: "123456",
	}
	user2 := UserRegistration{
		UserName: "Vl",
		Email:    "test@gmail.com",
		Password: "123456",
	}
	user3 := UserRegistration{
		UserName: "Vladislav",
		Email:    "testgmail.com",
		Password: "123456",
	}
	user4 := UserRegistration{
		UserName: "Vladislav",
		Email:    "test@gmail.com",
		Password: "123",
	}
	user5 := UserRegistration{
		UserName: "Vladislav",
		Email:    "test@gmail.com",
		Password: "123456",
	}

	PrintInfo(&user1)
	PrintInfo(&user2)
	PrintInfo(&user3)
	PrintInfo(&user4)
	PrintInfo(&user5)

}
