package main

import "fmt"

func main() {
	// Create instances of both media centers with default values
	xbmc := &XbMc {
		vol: 35,
		channel: 99,
		isOn: true,
	}
	osmc := &OsMc {
		currentChan: 21,
		currentVolume: 11,
		mcOn: true,
	}

	// The interface was written for XBMC
	// so it can be used directly
	xbmc.turnOn()
	xbmc.volumeDown()
	xbmc.volumeUp()
	xbmc.channelDown()
	xbmc.channelDown()
	xbmc.channelUp()
	xbmc.goToChannel(1)
	xbmc.turnOff()

	fmt.Println("-----------------------")

	// XBMC needs to use an adapter to be
	// able to use the OSMC interface
	mcAdapter := &osmcAdapter {
		mc: osmc, 
	}
	mcAdapter.turnOn()
	mcAdapter.volumeDown()
	mcAdapter.volumeUp()
	mcAdapter.channelDown()
	mcAdapter.channelDown()
	mcAdapter.channelUp()
	mcAdapter.goToChannel(21)
	mcAdapter.turnOff()
}