package main 

import "fmt"

type Student struct{
	Name string
	Grade int
}


func ImproveGrade(std *Student, points int){
	
	std.Grade += points
	if std.Grade > 100{
		std.Grade = 100
	}
	}


func main(){
	std1 := Student{
		Name: "Vladislav",
		Grade: 95,
	}
	ImproveGrade(&std1, 3)
	fmt.Println("student's name:", std1.Name)
	fmt.Println("student's grade:", std1.Grade)
}