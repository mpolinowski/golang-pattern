package main

import "fmt"

// API for the XBMC Kodi Media Center
type XbMc struct {
	vol		int
	channel	int
	isOn	bool
}

func (mc *XbMc) turnOn() bool {
	fmt.Println("XBMC is On")
	mc.isOn = true
	return mc.isOn
}

func (mc *XbMc) turnOff() bool {
	fmt.Println("XBMC is Off")
	mc.isOn = false
	return mc.isOn
}

func (mc *XbMc) volumeUp() int {
	mc.vol++
	fmt.Println("XBMC Volume:", mc.vol)
	return mc.vol
}

func (mc *XbMc) volumeDown() int {
	mc.vol--
	fmt.Println("XBMC Volume:", mc.vol)
	return mc.vol
}

func (mc *XbMc) channelUp() int {
	mc.channel++
	fmt.Println("XBMC Channel:", mc.channel)
	return mc.channel
}

func (mc *XbMc) channelDown() int {
	mc.channel--
	fmt.Println("XBMC Channel:", mc.channel)
	return mc.channel
}


func (mc *XbMc) goToChannel(chn int) {
	mc.channel = chn
	fmt.Println("XBMC Channel:", mc.channel)
}