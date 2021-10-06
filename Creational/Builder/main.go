package main

import (
	"fmt"
)

func main() {
	var bldr = newNotificationBuilder()

	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(3)
	bldr.SetMessage("This is a level 3 alert!")
	bldr.SetNotType("alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}
}