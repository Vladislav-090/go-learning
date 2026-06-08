package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Available bool
}

func (b *Book) Borrow() {
	if b.Available {
		b.Available = false
		fmt.Println("Book borrowed")
	} else {
		fmt.Println("Book is not available")
	}
}

func (b *Book) Return() {
	if !b.Available {
		b.Available = true
		fmt.Println("Book returned")
	} else {
		fmt.Println("Book is already available")
	}
}

func (b Book) PrintInfo() {
	fmt.Println("Tittle:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Available:", b.Available)
}

func main() {
	book:= Book{
		Title: "Clean code",
		Author: "Mr.Pastukhov",
		Available: false,
	}
	book.PrintInfo()
	book.Borrow()
	book.Borrow()
	book.Return()
	book.PrintInfo()
}