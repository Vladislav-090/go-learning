package main

import "fmt"

type UserAgeStorageService interface {
	AddUser(name string, age int)
	DeleteUser(name string)
	GetAgeByName(name string) int
	PrintUsers()
	UsersCount() int
}

type UserAgeStorage struct {
	User map[string]int
}

func (u *UserAgeStorage) AddUser(name string, age int) {
	u.User[name] = age
	fmt.Printf("User %s : %d has been added! \n", name, age)
}

func (u *UserAgeStorage) DeleteUser(name string) {
	delete(u.User,name)
	fmt.Printf("User %s deleted \n", name)
}

func (u *UserAgeStorage) GetAgeByName(name string) int {
	return u.User[name]
}

func (u *UserAgeStorage) UsersCount()int {
	return len(u.User)
}

func (u *UserAgeStorage) PrintUsers() {
	for name, age := range u.User {
		fmt.Printf("User: %s, age: %d\n", name, age)
	}
}


func PrintInfo( u UserAgeStorageService, name string, age int) {
	u.PrintUsers()

	u.AddUser(name, age)

	fmt.Printf("Age of %s is: %d\n", name, u.GetAgeByName(name))

	fmt.Println("Users count:", u.UsersCount())

	u.PrintUsers()

	u.DeleteUser(name)

	fmt.Println("Users count:", u.UsersCount())

	u.PrintUsers()
}

func main() {
	userStorage := UserAgeStorage{
		User: map[string]int{
			"Vladislav" : 29,
			"Viola" : 30,
			"Samuel" : 45,
		},
	}
	PrintInfo(&userStorage, "Afina", 3)
}