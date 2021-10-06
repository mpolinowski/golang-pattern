package main

import "fmt"

func main() {
	// Generate 4 publication objects of type magazine or newspaper
	mag1, _ := newPublication("magazine", "Programmer.Inc", 47, "WTF! kubectl")
	mag2, _ := newPublication("magazine", "Rust Formation", 46, "New & Shiny")
	news1, _ := newPublication("newspaper", "Vanuatu Daily", 44, "Island Life")
	news2, _ := newPublication("newspaper", "clickBait", 42, "The Sun")
	// Pint out details
	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails (pub iPublication) {
	fmt.Println()
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("----------------------\n")
}