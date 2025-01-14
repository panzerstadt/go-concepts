package gomindset

import (
	"fmt"
	"time"
)

// a struct is a wrapper around a data
// like a group of data
// the data "title" is grouped inside something
// called Book
// this is technically a data structure
type Book struct {
	title    string
	author   string
	numPages int

	isSaved bool
	savedAt time.Time
}

// all the things we can do with this data structure
// 1. we can read data from this book
// 2. we can write data into this book

// if we wanna do stuff with book
// we can only do that by using functions
// functions are how we interact with this
// data structure

// the following two functions are equivalent,
// the upper one is just syntactic sugar

func (book *Book) saveBook() {
	book.isSaved = true
	book.savedAt = time.Now()
}

func saveBook(book *Book) {
	book.isSaved = true
	book.savedAt = time.Now()
}

func BookTest() {

	b := Book{title: "foo", author: "bar", numPages: 2, isSaved: false, savedAt: time.Now()}
	fmt.Println(b)
	time.Sleep(time.Second * 2)
	b.saveBook()
	fmt.Println(b)
}

// golang is ONLY about
// - structures (struct)
// - data (the stuff in the struct)
// - functions (the 'methods')
