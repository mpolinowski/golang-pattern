package main

// Create an interface for the iterator function
// that will create the iterator object
type IterableCollection interface {
	createIterator() iterator
}

// The interface for the next and hasnext functions
// that will return the collection items
type iterator interface {
	hasNext() bool
	next() *Book
}

// Concrete iterator for the book collection
type BookIterator struct {
	current		int
	books		[]Book
}

// Check if there are still more books in the collection
func (b* BookIterator) hasNext() bool {
	if b.current < len(b.books) {
		return true
	}
	return false
}

// If there is a next book I want to pull it
func (b *BookIterator) next() *Book {
	if b.hasNext() {
		bk := b.books[b.current]
		b.current++
		return &bk
	}
	return nil
}