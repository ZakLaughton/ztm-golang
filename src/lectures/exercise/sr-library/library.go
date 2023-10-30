//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

// --Summary:
//
//	Create a program to manage lending of library books.
//
// --Requirements:
// * The library must have books and members, and must include:
//   - Which books have been checked out
//   - What time the books were checked out
//   - What time the books were returned
type Book struct {
	title        string
	checkedOut   bool
	checkoutTime time.Time
	returnTime   time.Time
}

type Member struct {
	name string
}

type Library struct {
	books   []Book
	members []Member
}

func printLibraryInfo(library *Library) {
	fmt.Println("Library Info:")
	for _, book := range library.books {
		fmt.Println("\tBook:", book.title)
		fmt.Println("\t\tChecked Out:", book.checkedOut)
		fmt.Println("\t\tCheckout Time:", book.checkoutTime)
		fmt.Println("\t\tReturn Time:", book.returnTime)
	}
	fmt.Println("Members:")
	for _, member := range library.members {
		fmt.Println("\tMember:", member.name)
	}
}

func checkoutBook(library *Library, book *Book) {
	book.checkedOut = true
	book.checkoutTime = time.Now()
}

//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

func main() {
	var myLibrary Library
	myLibrary.books = append(myLibrary.books, Book{title: "The Hobbit"}, Book{title: "The Lord of the Rings"}, Book{title: "The Silmarillion"}, Book{title: "The Children of Hurin"})
	printLibraryInfo(&myLibrary)
	myLibrary.members = append(myLibrary.members, Member{name: "John"}, Member{name: "Jane"}, Member{name: "Joe"})
	printLibraryInfo(&myLibrary)
}
