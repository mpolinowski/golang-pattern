package main

import "fmt"

// Interface for the observer type
type observer interface {
	onUpdate(data string)
}

// Type for the observer id
// that is used to register as DataListner
type DataListener struct {
	Name string
}

// This function is called when the observer is notified
func (dl *DataListener) onUpdate(data string) {
	fmt.Println(dl.Name, "received update:", data)
}