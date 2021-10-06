package main

import "fmt"

// Define magazine type using the publication interface from ./publication
type magazine struct {
	publication
}

// Return string representation of magazine name type
func (m magazine) String() string {
	return fmt.Sprintf("Magazine name: %s", m.name)
}

// Return the magazine object
func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			name: name,
			pages: pages,
			publisher: publisher,
		},
	}
}