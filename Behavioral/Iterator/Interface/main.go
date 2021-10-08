package main

import "fmt"

// Iterate using the callback function
func main() {

	// Create iterator
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}

}