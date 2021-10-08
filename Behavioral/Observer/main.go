package main

func main() {
	// Create 2 observers and assign names
	observer1 := DataListener{
		Name: "Brother Number One",
	}
	observer2 := DataListener{
		Name: "Brother Number Two",
	}
	// Add a subject object that is observed
	subject := &DataSubject{}
	// Register observers to subject
	subject.registerObserver(observer1)
	subject.registerObserver(observer2)

	// Change the data field that will trigger the notification call
	subject.ChangeItem("Update 1")
	subject.ChangeItem("Update 2")
	subject.ChangeItem("Update 3")

	// Unregister the observers
	subject.unregisterObserver(observer1)
	// After removing one the following update should
	// only be send to remaining observer
	subject.ChangeItem("Update 4")
}