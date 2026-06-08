package main

import "fmt"

type Student struct{
	Name string
	Grade int
}

func(s *Student) ImproveGrade(points int) {
	s.Grade += points

	if s.Grade > 100 {
		s.Grade = 100
		fmt.Println("You have a max grade! ", s.Grade)
		
	} else {
		
		fmt.Println("Now your grade is: ", s.Grade)
	}
}

func main() {
	student := Student{
		Name: "Vladislav",
		Grade: 85,
	}
	student.ImproveGrade(20)
}