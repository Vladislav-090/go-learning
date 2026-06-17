package main

import (
	"errors"
	"fmt"
)

type Course struct {
	Title       string
	Students    []string
	MaxStudents int
}

func (c *Course) AddStudent(name string) error {

	if name == "" {
		return errors.New("Name cannot be empty!")
	}

	for _, student := range c.Students {
		if student == name {
			return errors.New("Already exists!")
		}
	}
	if len(c.Students) >= c.MaxStudents {
		return errors.New("There are no more places on the course")

	}
	c.Students = append(c.Students, name)
	fmt.Println("Student added", name)
	return nil
}

func (c *Course) DeleteStudent(name string) error {
	if name == "" {
		return errors.New("Name cannot be empty!")
	}

	for i, students := range c.Students {
		if students == name {
			c.Students = append(c.Students[:i], c.Students[i+1:]...)
			fmt.Println("Student deleted")
			return nil
		}
	}
	return errors.New("Student not found")
}

func (c *Course) StudentsCount() int {
	return len(c.Students)
}

func (c *Course) PrintInfo() {
	fmt.Println("Course info: \nTitle:", c.Title, "\nStudents:", c.Students, "\nMax students on course:", c.MaxStudents)
}

func PrintInfo(c *Course) {
	c.PrintInfo()

	err := c.AddStudent("Viola")
	if err != nil {
		fmt.Println("Add Student error", err)
	}

	err = c.AddStudent("Alex")
	if err != nil {
		fmt.Println("Add Student error", err)
	}

	err = c.AddStudent("Vladislav")
	if err != nil {
		fmt.Println("Add Student error", err)
	}

	err = c.DeleteStudent("Viola")
	if err != nil {
		fmt.Println("Delete Student error", err)
	}

	err = c.DeleteStudent("Bruce")
	if err != nil {
		fmt.Println("Delete Student error", err)
	}

	c.PrintInfo()
}

func main() {
	course := Course{
		Title:       "Golang",
		Students:    []string{"Vladislav"},
		MaxStudents: 2,
	}
	PrintInfo(&course)
}
