package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u *User) UpdateAge(age int) error {
	if age < 0 {
		return errors.New("Age must be positive!")
	}
	if age > 130 {
		return errors.New("Age must be less then 130!")
	}

	u.Age = age
	fmt.Println("Years has been updated!")
	return nil
}

func (u *User) IsAdult() (bool, error) {
	if u.Age < 0 {
		return false, errors.New("Age must be positive!")
	}
	if u.Age >= 18 {
		return true, nil
	}

	return false, nil
}

func (u *User) PrintInfo() {
	fmt.Println("User name:", u.Name, "Age:", u.Age)
}

func PrintInfo(u *User) {
	u.PrintInfo()

	IsAdult, err := u.IsAdult()
	if err != nil {
		fmt.Println("IsAdult error", err)
	} else {
		fmt.Println("User is adult", IsAdult)
	}

	err = u.UpdateAge(30)
	if err != nil {
		fmt.Println("UpdateAge error", err)
	}
	u.PrintInfo()

	IsAdult, err = u.IsAdult()
	if err != nil {
		fmt.Println("IsAdult error", err)
	} else {
		fmt.Println("User is adult", IsAdult)
	}

}

func main() {
	user := User{
		Name: "Vladislav",
		Age:  17,
	}
	PrintInfo(&user)
}
