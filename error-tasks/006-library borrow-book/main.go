package main

import (
	"errors"
	"fmt"
)

type Book struct {
	Title     string
	Available bool
}

func (b *Book) Borrow() error {
	if b.Available == false {
		return errors.New("Book is already borrowed!")
	}
	b.Available = false
	fmt.Println("Book has been taken")
	return nil
}

func (b *Book) ReturnBook() error {
	if b.Available == true {
		return errors.New("Book is already available!")
	}
	b.Available = true
	fmt.Println("Book has been returned")
	return nil
}

func (b *Book) PrintInfo() {
	fmt.Println("Title:", b.Title, "Is Available:", b.Available)
}

func PrintInfo(b *Book) {
	b.PrintInfo()

	err := b.Borrow()

	if err != nil {
		fmt.Println("Borrow error", err)
	}

	err = b.Borrow()
	if err != nil {
		fmt.Println("Borrow error", err)
	}

	err = b.ReturnBook()
	if err != nil {
		fmt.Println("ReturnBook error", err)
	}

	err = b.ReturnBook()
	if err != nil {
		fmt.Println("ReturnBook error", err)
	}

	b.PrintInfo()
}

func main() {
	book := Book{
		Title:     "Lion King",
		Available: true,
	}
	PrintInfo(&book)
}
