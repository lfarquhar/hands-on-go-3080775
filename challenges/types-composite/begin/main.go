// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library struct {
	name  string
	books map[string]book
}

func newLibrary(name string) library {
	lib := library{
		name:  name,
		books: make(map[string]book),
	}

	return lib
}

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l.books[b.title] = b
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(name string) []book {

	result := make([]book, 0)

	for _, v := range l.books {
		
		if (v.author.name == name) {
			result = append(result, v)
		}

	}

	return result
}

func main() {
	// create a new library
	lib := newLibrary("Merrimack Town Library")

	// add 2 books to the library by the same author
	b1 := book{
		title:  "foo",
		author: author{name: "Jimmy"},
	}
	b2 := book{
		title:  "bar",
		author: author{name: "Luie"},
	}

	b3 := book{
		title:  "X",
		author: author{name: "Jimmy"},
	}

	lib.addBook(b1)
	lib.addBook(b2)
	lib.addBook(b3)

	// dump the library with spew
	// spew.Dump(lib)

	if(lib.name != "Merrimack Town Library"){
		fmt.Println("it didn't work")
	}
	// lookup books by known author in the library
	var books = lib.lookupByAuthor("Jimmy")

	// print out the first book's title and its author's name
	// spew.Dump(lib)
	spew.Dump(books)
}
