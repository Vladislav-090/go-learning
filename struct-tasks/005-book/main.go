package main

import "fmt"

type Book struct {
	Tittle   string
	Author string
	Pages  int
}

func AddPages(p *Book, add int) {
	p.Pages += add
}

func main() {
	book := Book{
		Tittle:   "Golang for begginers",
		Author: "Pastukhov Vladislav",
		Pages:  100,
	}
	AddPages(&book, 50)
	fmt.Println("Name:", book.Tittle)
	fmt.Println("Author:", book.Author)
	fmt.Println("Pages:", book.Pages)


}