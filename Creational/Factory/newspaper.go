package main

import "fmt"

// Define newspaper type using the publication interface from ./publication
type newspaper struct {
	publication
}

// Return string representation of newspaper name type
func (n newspaper) String() string {
	return fmt.Sprintf("Newspaper name: %s", n.name)
}

// Return the newspaper object
func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name: name,
			pages: pages,
			publisher: publisher,
		},
	}
 }