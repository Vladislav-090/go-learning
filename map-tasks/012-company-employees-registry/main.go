package main

import "fmt"

type CompanyEmployeesRegistry interface {
	AddEmployee(company string, employee string)
	DeleteEmployee(company string, employee string)
	GetEmployees(company string) []string
	ViewAllCompanies()
	CompaniesCount() int
}

type CompanyEmployeesStruct struct {
	Companies map[string][]string
}

func (c *CompanyEmployeesStruct) AddEmployee(company string, employee string) {
	if _, exist := c.Companies[company]; !exist {
		c.Companies[company] = []string{}
		fmt.Println("New Company without employees created!", company)
	}
	for _, currentEmployee := range c.Companies[company] {
		if currentEmployee == employee {
			fmt.Println("Employee already exist!", employee)
			return
		}
	}

	c.Companies[company] = append(c.Companies[company], employee)
	fmt.Println("Employe Added to company", company, employee)
}

func (c *CompanyEmployeesStruct) DeleteEmployee(company string, employee string) {
	employees, exist := c.Companies[company]
	if !exist {
		fmt.Println("Company not found!", company)
		return
	}
	for i, currentEmploye := range employees {
		if currentEmploye == employee {
			employees = append(employees[:i], employees[i+1:]...)
			c.Companies[company] = employees
			fmt.Println("Employee deleted!", employee)
			return
		}
	}

	fmt.Println("Employee not found", employee)

}

func (c *CompanyEmployeesStruct) GetEmployees(company string) []string {
	return c.Companies[company]
}

func (c *CompanyEmployeesStruct) ViewAllCompanies() {
	for companies, employees := range c.Companies {
		fmt.Println("Companies info:", companies, employees)
	}
}

func (c *CompanyEmployeesStruct) CompaniesCount() int {
	return len(c.Companies)
}

func PrintInfo(c CompanyEmployeesRegistry, company string, employee string) {
	c.ViewAllCompanies()
	fmt.Println("Companies count:", c.CompaniesCount())

	c.AddEmployee(company, employee)
	fmt.Println("Employees of company:", company, c.GetEmployees(company))

	c.ViewAllCompanies()

	c.DeleteEmployee(company, employee)
	fmt.Println("Employees of company:", company, c.GetEmployees(company))

	c.ViewAllCompanies()
	fmt.Println("Companies count:", c.CompaniesCount())
}

func main() {
	companies := CompanyEmployeesStruct{
		Companies: map[string][]string{
			"Backend": {"Vladislav", "Alex"},
			"DevOps":  {"Bruce", "Max"},
		},
	}

	PrintInfo(&companies, "Backend", "Viola")
}
