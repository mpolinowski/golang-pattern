package main

import "fmt"

// We need an interface for the observable (subject) type
// The subject needs to be able to register/deregister observers
// and to notify them in case of a state change
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// Registered observers need to stored in an array
// and a data field that is observed needs to be defined
type DataSubject struct {
	observers	[]DataListener
	field		string
}

// This function will trigger the notify function
func (ds *DataSubject) ChangeItem(data string) {
	// write data to observed data field
	ds.field = data
	// trigger notification
	ds.notifyAll()

}

// Register observers to DataListeners
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
	fmt.Println(o.Name, "was registered")
}

// Deregister observers from DataListeners
func (ds *DataSubject) unregisterObserver(o DataListener) {
	var observerList []DataListener
	// Loop through all observers in list
	for _, observer := range ds.observers {
		// Recreate the list without the observer to be removed
		if o.Name != observer.Name {
			observerList = append(observerList, observer)
			fmt.Println(o.Name, "was deregistered")
		}
	}
	ds.observers = observerList
}

// Notify all observers with updated data
func (ds *DataSubject) notifyAll() {
	// select all observers in list
	for _, observers := range ds.observers {
		// and trigger update function with updated field data
		observers.onUpdate(ds.field)
	}


}