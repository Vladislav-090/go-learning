package main

import "fmt"

type UserManager interface {
	CreateUser(name string)
	DeleteUser(name string) 
	GetUserCount()int
	GetServiceName()string
}

type AdminService struct {
	Name string
	UserCount int
}

type CRMService struct {
	Name string
	UserCount int
}


func (a *AdminService) CreateUser(name string) {
	a.UserCount ++
	fmt.Println("User created!", name)
}

func (a *AdminService) DeleteUser(name string) {
	if a.UserCount == 0 {
		fmt.Println("dont have any users!")
		return
	}
	a.UserCount --
}

func (a *AdminService) GetUserCount() int {
	return a.UserCount
}
func (a *AdminService) GetServiceName()string {
	return a.Name
}

/////////////////////////

func (c *CRMService) CreateUser(name string) {
	c.UserCount ++
	fmt.Println("User created!", name)
}

func (c *CRMService) DeleteUser(name string) {
	if c.UserCount == 0 {
		fmt.Println("dont have any users!")
		return
	}
	c.UserCount --
	fmt.Println("User Deleted!")
}

func (c *CRMService) GetUserCount() int {
	return c.UserCount
}
func (c *CRMService) GetServiceName()string {
	return c.Name
}

func PrintInfo(u UserManager, name string) {
	fmt.Println("Service name:", u.GetServiceName())

	u.CreateUser(name)
	fmt.Println("User count:", u.GetUserCount())

	u.DeleteUser(name)
	fmt.Println("User count:", u.GetUserCount())

	

}

func main () {
	admin := AdminService{
		Name: "Admin",
		UserCount: 4,
	}

	crm := CRMService{
		Name: "CRM",
		UserCount: 10,
	}

	PrintInfo(&admin, "Vladislav")
	PrintInfo(&crm, "Viola")
}