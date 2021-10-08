package main

import "fmt"

// Iterate using the callback function
func main() {
	// Iterate over library via callback function (lib defined in books.go)
	lib.IterateBooks(myBookCallback)

	// Iterate via anonymous function
	lib.IterateBooks(func (b Book) error {
		fmt.Println("Published:", b.Year)
		return nil
	})

	// Create iterator

}

// Callback function to process each book that
// is returned by the iteration
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}