package main

import "fmt"

type StudentRegistryService interface {
	AddStudent(name string)
	DeleteStudent(name string)
	UpdateStudent(oldName string, newName string)
	PrintStudents()
	StudentsCount() int
}

type StudentRegistry struct {
	Students []string
}

func (s *StudentRegistry) AddStudent(name string) {
	for _, students := range s.Students {
		if students == name {
			fmt.Println("Student already exist!", name)
			return
		}
	}
	s.Students = append(s.Students, name)
	fmt.Println("Student has been added!", name)
}

func (s *StudentRegistry) DeleteStudent(name string) {
	for i,students := range s.Students {
		if students == name {
			s.Students = append(s.Students[:i], s.Students[i+1:]...)
			fmt.Println("Student deleted!", name)
			return
		}
	}
	fmt.Println("Student not found!", name)
}

func (s *StudentRegistry) UpdateStudent(oldName string, newName string) {
	for _,students := range s.Students{
		if students == newName {
			fmt.Println("student already exist!",newName)
			return
		}
	}

	for i, students := range s.Students {
		if students == oldName {
			s.Students[i] = newName
			fmt.Println("Old student updated:", oldName, "->", newName)
			return
		}
	}
	fmt.Println("Student not found!")
	}

	func (s *StudentRegistry) PrintStudents() {
		for _, students := range s.Students {
			fmt.Println("Student:", students)
		}
	}


	func (s *StudentRegistry) StudentsCount() int {
		return len(s.Students)
	}


	func PrintInfo(s StudentRegistryService, name string, oldName string, newName string){
s.PrintStudents()
	fmt.Println("Students count:", s.StudentsCount())

	s.AddStudent(name)
	s.PrintStudents()
	fmt.Println("Students count:", s.StudentsCount())

	s.UpdateStudent(oldName, newName)
	s.PrintStudents()
	fmt.Println("Students count:", s.StudentsCount())

	s.DeleteStudent(name)
	s.PrintStudents()
	fmt.Println("Students count:", s.StudentsCount())
	}


	func main() {
		studentsRegistry := StudentRegistry{
	Students: []string{
		"Vlad",
		"John",
		"Viola",
	},
}

PrintInfo(&studentsRegistry, "Sveta", "Vlad", "John")
	}