package main

import "fmt"

// Define the type of book
type BookType int

// Predefined book types
const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

// Types of data that describe a book 
type Book struct {
	Name		string
	Year		int
	PageCount	int
	ISBN		string
	Type		BookType
}

// Create a type that holds all the books
type Library struct {
	Collection []Book
}

// Call callBack function to iterate over all books
// and pushes the result back to the caller
func (l *Library) IterateBooks(f func(Book) error) {
	// Iterate over all books and trigger callback function
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// Create an iterator that can access the book collection on demand



// Create library structure

var lib *Library = &Library {
	Collection: []Book {
		{
			Name:		"Leviathan Wakes",
			Year:		2011,
			PageCount:	592,
			ISBN:		"978-0-316-12908-4",
			Type:		HardCover,
		},
		{
			Name:		"Caliban's War",
			Year:		2012,
			PageCount:	595,
			ISBN:		"978-1-841-49990-1",
			Type:		SoftCover,
		},
		{
			Name:		"Abaddon's Gate",
			Year:		2013,
			PageCount:	539,
			ISBN:		"978-0-316-12907-7",
			Type:		HardCover,
		},
		{
			Name:		"Cibola Burn",
			Year:		2014,
			PageCount:	583,
			ISBN:		"978-0-316-21762-0",
			Type:		PaperBack,
		},
		{
			Name:		"Nemesis Games",
			Year:		2015,
			PageCount:	544,
			ISBN:		"978-0-316-21758-3",
			Type:		HardCover,
		},
		{
			Name:		"Babylon's Ashes",
			Year:		2016,
			PageCount:	608,
			ISBN:		"978-0-316-33474-7",
			Type:		HardCover,
		},
		{
			Name:		"Persepolis Rising",
			Year:		2017,
			PageCount:	560,
			ISBN:		"978-0-316-33283-5",
			Type:		HardCover,
		},
		{
			Name:		"Tiamat's Wrath",
			Year:		2019,
			PageCount:	544,
			ISBN:		"978-0-316-33286-6",
			Type:		HardCover,
		},
		{
			Name:		"Leviathan Falls",
			Year:		2021,
			PageCount:	528,
			ISBN:		"978-0-356-51039-2",
			Type:		EBook,
		},
	},
}


// func main() {
// }