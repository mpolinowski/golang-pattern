package main

// Interface that works with the XBMC Kodi Media Center and has to be adapted
// to the OSMC Media Center so we can use our program with both.
type MediaCenter interface {
	volumeDown() int
	volumeUp() int
	channelDown() int
	channelUp() int
	turnOff()
	turnOn()
	goToChannel(ch int)
}