package main

import "fmt"

type BookLibraryService interface {
	GetNameOfLibrary() string
	AddBook(name string)
	DeleteBook(name string)
	GetCountsOfBooks() int
	AllBooks()
}

type BookLibrary struct {
	Name  string
	Books []string
}

func (b *BookLibrary) GetNameOfLibrary() string {
	return b.Name
}

func (b *BookLibrary) AddBook(name string) {
	b.Books = append(b.Books, name)
	fmt.Println("The new Book added:", name)
}

func (b *BookLibrary) DeleteBook(name string) {
	for index, book := range b.Books {
		if book == name {
			b.Books = append(b.Books[:index], b.Books[index+1:]...)
			fmt.Println("The Book has been Deleted!", name)
			return
		}
	}
	fmt.Println("Book not found!", name)
}

func (b *BookLibrary) GetCountsOfBooks() int {
	return len(b.Books)
}

func (b *BookLibrary) AllBooks() {
	for _, book := range b.Books {
		fmt.Println("Book wich available now:", book)
	}
}

func PrintInfo(b BookLibraryService, name string) {
	fmt.Println(b.GetNameOfLibrary())

	b.AllBooks()
	b.AddBook(name)
	fmt.Println("All books wich available:")
	b.AllBooks()
	fmt.Println("Quantity availabl books", b.GetCountsOfBooks())
	b.DeleteBook(name)
	fmt.Println("All books wich available")
	b.AllBooks()
	fmt.Println("Quantity available books", b.GetCountsOfBooks())
}

func main() {
	booklibrary := BookLibrary{
		Name:  "Pushkin library",
		Books: []string{"The Little Prince", "The Old Man and the Sea"},
	}

	PrintInfo(&booklibrary, "Animal Farm")
}
