package main

import "fmt"

type LibraryRegistryService interface {
	AddBook(name string)
	DeleteBook(name string)
	UpdateTitleBook(oldName string, newName string)
	ViewAllBooks()
	GetBooksCount() int
	GetLibraryName() string
}

type LibraryRegistry struct {
	Name  string
	Books []string
}

func (l *LibraryRegistry) AddBook(name string) {
	for _, book := range l.Books {
		if book == name {
			fmt.Println("Book already exist!", name)
			return
		}
	}
	l.Books = append(l.Books, name)
	fmt.Println("Book added:", name)
}

func (l *LibraryRegistry) DeleteBook(name string) {
	for i, book := range l.Books {
		if book == name {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Println("Book deleted!", name)
			return
		}
	}
	fmt.Println("Book not found!", name)
}

func (l *LibraryRegistry) UpdateTitleBook(oldName string, newName string) {
	for _, book := range l.Books {
		if book == newName {
			fmt.Println("Book already exist!", newName)
			return
		}
	}
	for i, book := range l.Books {
		if book == oldName {
			l.Books[i] = newName
			fmt.Println("Title of book has been updated! ", oldName, "to", newName)
			return
		}
	}
	fmt.Println("Book not found", oldName)
}

func (l *LibraryRegistry) ViewAllBooks() {
	for _, book := range l.Books {
		fmt.Println("Available Book:", book)
	}
}

func (l *LibraryRegistry) GetBooksCount() int {
	return len(l.Books)
}

func (l *LibraryRegistry) GetLibraryName() string {
	return l.Name
}

func PrintInfo(l LibraryRegistryService, name string, oldName string, newName string) {
	fmt.Println("Library Name is:", l.GetLibraryName())

	l.ViewAllBooks()
	fmt.Println("Quantity of Avalable books:", l.GetBooksCount())

	l.AddBook(name)
	l.UpdateTitleBook(oldName, newName)
	l.ViewAllBooks()
	fmt.Println("Quantity of Avalable books:", l.GetBooksCount())

	l.DeleteBook(name)

	l.ViewAllBooks()
	fmt.Println("Quantity of Avalable books:", l.GetBooksCount())

}

func main() {
	libraryRegistry := LibraryRegistry{
		Name: "Karagandy Library",
		Books: []string{
			"lion king",
			"Arlo",
			"Batman",
		},
	}

	PrintInfo(&libraryRegistry, "Autobiography", "lion king", "LION KING2")
}
