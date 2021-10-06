package main

import "fmt"

// Create specified publication in factory function
func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	// Create publication based on specified type
	if pubType == "newspaper" {
		return createNewspaper(name, pg, pub), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}
	
	return nil, fmt.Errorf("No such publication")
}