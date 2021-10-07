package main

type osmcAdapter struct {
	//
	mc *OsMc
}

// Adapt XBMC turnOff() bool -> OSMC setPower bool
func (osmcAdapt *osmcAdapter) turnOff() {
	osmcAdapt.mc.setPower(false)
}

// Adapt XBMC turnOn() bool -> OSMC setPower bool
func (osmcAdapt *osmcAdapter) turnOn() {
	osmcAdapt.mc.setPower(true)
}

// Adapt XBMC volumeDown() int -> OSMC getVolume/setVolume int
func (osmcAdapt *osmcAdapter) volumeDown() int {
	vol := osmcAdapt.mc.getVolume() - 1
	osmcAdapt.mc.setVolume(vol)
	return vol
}

// Adapt XBMC volumeUp() int -> OSMC getVolume/setVolume int
func (osmcAdapt *osmcAdapter) volumeUp() int {
	vol := osmcAdapt.mc.getVolume() + 1
	osmcAdapt.mc.setVolume(vol)
	return vol
}

// Adapt XBMC channelDown() int -> OSMC getChannel/setChannel int
func (osmcAdapt *osmcAdapter) channelDown() int {
	chn := osmcAdapt.mc.getChannel() - 1
	osmcAdapt.mc.setChannel(chn)
	return chn
}

// Adapt XBMC channelUp() int -> OSMC getChannel/setChannel int
func (osmcAdapt *osmcAdapter) channelUp() int {
	chn := osmcAdapt.mc.getChannel() + 1
	osmcAdapt.mc.setChannel(chn)
	return chn
}

// Adapt XBMC goToChannel() int -> OSMC getChannel/setChannel int
func (osmcAdapt *osmcAdapter) goToChannel(chn int) {
	osmcAdapt.mc.setChannel(chn)
}