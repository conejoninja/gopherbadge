package main

import (
	"machine"
	"machine/usb/adc/midi"
	"time"
)

func main() {

	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnUp := machine.BUTTON_UP
	btnLeft := machine.BUTTON_LEFT
	btnDown := machine.BUTTON_DOWN
	btnRight := machine.BUTTON_RIGHT
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInput})

	notes := []midi.Note{midi.C4, midi.D4, midi.E4, midi.F4, midi.G4, midi.B4}
	midichannel := uint8(1) // MIDI channels are 0-15 e.g. 1-16

	note := -1
	oldNote := -1
	for {
		note = -1
		if !btnA.Get() {
			note = 0
		}
		if !btnB.Get() {
			note = 1
		}

		if !btnLeft.Get() {
			note = 2
		}
		if !btnRight.Get() {
			note = 3
		}
		if !btnUp.Get() {
			note = 4
		}
		if !btnDown.Get() {
			note = 5
		}
		println(note)
		if note != oldNote {
			if oldNote != -1 {
				midi.Midi.NoteOff(0, midichannel, notes[oldNote], 50)
			}
			if note != -1 {
				midi.Midi.NoteOn(0, midichannel, notes[note], 50)
			}
			oldNote = note
		}
		time.Sleep(100 * time.Millisecond)
	}
}
