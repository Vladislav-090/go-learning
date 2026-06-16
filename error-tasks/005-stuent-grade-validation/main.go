package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) UpdateGrade(grade int) error {
	if grade > 100 {
		return errors.New("Grade must be less then 100")
	}
	if grade < 0 {
		return errors.New("Grade cant be less then 0")
	}

	s.Grade = grade
	fmt.Println("Grade has been updated")
	return nil
}

func (s *Student) IncreaseGrade(point int) error {
	if s.Grade+point > 100 {
		return errors.New("Grade cant be more then 100")
	}
	if point <= 0 {
		return errors.New("Point must be positive!")
	}

	s.Grade += point
	fmt.Println("The grade has been increased")
	return nil
}

func (s *Student) IsExcellent() (bool, error) {
	if s.Grade < 0 {
		return false, errors.New("Grade cant be less then 0!")
	}

	if s.Grade >= 90 {
		return true, nil
	}

	return false, nil
}

func (s *Student) PrintInfo() {
	fmt.Println("Student name:", s.Name, "Grade:", s.Grade)
}

func PrintInfo(s *Student) {
	s.PrintInfo()

	IsExelente, err := s.IsExcellent()
	if err != nil {
		fmt.Println("IsExelente error!", err)
	} else {
		fmt.Println("Grade Is exelente", IsExelente)
	}

	err = s.UpdateGrade(92)
	if err != nil {
		fmt.Println("UpgradeGrade error", err)
	}

	err = s.IncreaseGrade(100)
	if err != nil {
		fmt.Println("IncreaseGrade error", err)
	}

	err = s.IncreaseGrade(0)
	if err != nil {
		fmt.Println("IncreaseGrade error", err)
	}

	err = s.IncreaseGrade(5)
	if err != nil {
		fmt.Println("IncreaseGrade error", err)
	}

	IsExelente, err = s.IsExcellent()
	if err != nil {
		fmt.Println("IsExelente error!", err)
	} else {
		fmt.Println("Grade Is exelente", IsExelente)
	}

	s.PrintInfo()
}

func main() {
	student := Student{
		Name:  "Vladislav",
		Grade: 50,
	}
	PrintInfo(&student)
}
