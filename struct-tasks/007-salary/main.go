package main 

import "fmt"

type Employee struct {
	Name string
	Salary int
}

func RaiseSalary(e *Employee, percent int) {
	e.Salary += e.Salary * percent / 100
}


func main() {
	 employee:= Employee{
		Name : "John", 
		Salary : 2000}

	RaiseSalary(&employee, 15)

	fmt.Println("our Employee's name is:", employee.Name)
	fmt.Println("our Employee's salary is:", employee.Salary)

}
