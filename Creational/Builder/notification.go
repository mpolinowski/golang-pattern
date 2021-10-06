package main

// Object to be created by the builder
// not capitalizing the names means they are not exported making them immutable
type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}