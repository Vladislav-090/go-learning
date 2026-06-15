package main

import "fmt"

type UserRolesRegistry interface {
	AddRole(user string, role string)
	DeleteRole(user string, role string)
	GetRolesOfUser(user string) []string
	ViewAllUsers()
	UsersCount() int
}

type UserRolesStruct struct {
	Users map[string][]string
}

func (u *UserRolesStruct) AddRole(user string, role string) {
	if _, exist := u.Users[user]; !exist {
		u.Users[user] = []string{}
		fmt.Println("New User with empty roles created!", user)
	}

	for _, roles := range u.Users[user] {
		if roles == role {
			fmt.Println("Role already exist!", role)
			return
		}
	}

	u.Users[user] = append(u.Users[user], role)
	fmt.Println("New role added:",user, role)
}

func (u *UserRolesStruct) DeleteRole(user string, role string) {
	roles, exist := u.Users[user]
	if !exist {
		fmt.Println("User not found", user)
		return
	}

	for i, currentRole := range roles {
		if currentRole == role {
			roles = append(roles[:i], roles[i+1:]...)
			u.Users[user] = roles
			fmt.Println("Role deleted", role)
			return
		}
	}

	fmt.Println("Role not found", role)
}

func (u *UserRolesStruct) GetRolesOfUser(user string) []string {
	return u.Users[user]
}

func (u *UserRolesStruct) ViewAllUsers() {
	for user, role := range u.Users {
		fmt.Println("User info:", user, role)
	}
}

func (u *UserRolesStruct) UsersCount() int {
	return len(u.Users)
}

func PrintInfo(u UserRolesRegistry, user string, role string) {
	u.ViewAllUsers()
	fmt.Println("Users count:", u.UsersCount())

	u.AddRole(user, role)
	fmt.Println("Roles of user:", user, u.GetRolesOfUser(user))

	u.ViewAllUsers()

	u.DeleteRole(user, role)
	fmt.Println("Roles of user:", user, u.GetRolesOfUser(user))

	u.ViewAllUsers()
	fmt.Println("Users count:", u.UsersCount())
}

func main() {
	registry := UserRolesStruct{
		Users: map[string][]string{
			"Vladislav": {"admin", "developer"},
			"Viola":     {"user"},
		},
	}

	PrintInfo(&registry, "Vladislav", "manager")
}
