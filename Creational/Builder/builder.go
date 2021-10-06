package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotType(notType string) {
	nb.NotType = notType
}

// Return finished notification object
func (nb *NotificationBuilder) Build() (*Notification, error) {

	// Error checking - if an Icon is given you also need an image
	if nb.Icon != "" && nb.Image == "" {
		return nil, fmt.Errorf("you need to provide an image when using an icon!")
	}

	// Error checking - max priority is 5
	if nb.Priority < 1 || nb.Priority > 5 {
		return nil, fmt.Errorf("priority has to be between 1 and 5")
	}

	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
