package main

import "fmt"

// API for the Open-Source Media Center
type OsMc struct {
	currentChan		int
	currentVolume	int
	mcOn			bool
}

func (mc *OsMc) getVolume() int {
	fmt.Println("OSMC Volume is:", mc.currentVolume)
	return mc.currentVolume
}

func (mc *OsMc) setVolume(vol int) {
	fmt.Println("Setting OSMC Volume to:", vol)
	mc.currentVolume = vol
}

func (mc *OsMc) getChannel() int {
	fmt.Println("OSMC channel is:", mc.currentChan)
	return mc.currentChan
}

func (mc *OsMc) setChannel(chn int) {
	fmt.Println("Setting OSMC Channel to:", chn)
	mc.currentChan = chn
}


func (mc *OsMc) setPower(mcOn bool) {
	if mcOn == true {
		fmt.Println("OSMC is On")
	} else {
		fmt.Println("OSMC is Off")
	}
	mc.mcOn = mcOn
}