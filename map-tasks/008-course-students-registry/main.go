package main

import "fmt"

type CourseStudentRegistry interface {
	AddStudent(course string, name string)
	DeleteStudent(course string, name string)
	GetStudents(course string) []string
	ViewAllCourses()
	CoursesCount() int
}

type CourseStudentStruct struct {
	Courses map[string][]string
}

func (c *CourseStudentStruct) AddStudent(course string, name string) {
	if _, exists := c.Courses[course]; !exists {
		c.Courses[course] = []string{}
		fmt.Println("New course created!", course)
	}

	for _, currentName := range c.Courses[course] {
		if currentName == name {
			fmt.Println("Student already exist!", name)
			return
		}
	}

	c.Courses[course] = append(c.Courses[course], name)
	fmt.Println("Student Added", course, name)
}

func (c *CourseStudentStruct) DeleteStudent(course string, name string) {
	names, exist := c.Courses[course]
	if !exist {
		fmt.Println("Course not found", course)
		return
	}

	for i, currentName := range names {
		if currentName == name {
			names = append(names[:i], names[i+1:]...)
			c.Courses[course] = names
			fmt.Println("Student deleted", name)
			return
		}
	}
	fmt.Println("Student not found!", name)
}

func (c *CourseStudentStruct) GetStudents(course string) []string {
	return c.Courses[course]
}

func (c *CourseStudentStruct) ViewAllCourses() {
	for course, name := range c.Courses {
		fmt.Println("Available Course", course, ",Students:", name)
	}
}

func (c *CourseStudentStruct) CoursesCount() int {
	return len(c.Courses)
}

func PrintInfo(c CourseStudentRegistry, course string, name string) {
	c.ViewAllCourses()
	fmt.Println("Courses count:", c.CoursesCount())

	c.AddStudent(course, name)

	fmt.Println("Students of course:", course)
	fmt.Println(c.GetStudents(course))

	c.ViewAllCourses()

	c.DeleteStudent(course, name)

	fmt.Println("Students of course:", course)
	fmt.Println(c.GetStudents(course))

	c.ViewAllCourses()
	fmt.Println("Courses count:", c.CoursesCount())
}

func main() {
	courseStudentStruct := CourseStudentStruct{
		Courses: map[string][]string{
			"Golang":     {"Vladislav", "Alex", "Bruce", "Max"},
			"PostgesSQL": {"Vladislav", "Bruce"},
		},
	}

	PrintInfo(&courseStudentStruct, "Golang", "Afina")
}
