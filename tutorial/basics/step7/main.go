package main

import (
	"machine"
	"time"
)

var bzrPin machine.Pin

func main() {

	// Enable the speaker
	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	// buzzer setup
	bzrPin = machine.SPEAKER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnUp = machine.BUTTON_UP
	btnLeft = machine.BUTTON_LEFT
	btnDown = machine.BUTTON_DOWN
	btnRight = machine.BUTTON_RIGHT
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !btnA.Get() {
			tone(1046)
		}
		if !btnB.Get() {
			break
		}

		if !btnLeft.Get() {
			tone(329)
		}
		if !btnRight.Get() {
			tone(739)
		}
		if !btnUp.Get() {
			tone(369)
		}
		if !btnDown.Get() {
			tone(523)
		}
	}
}

func tone(tone int) {
	for i := 0; i < 10; i++ {
		bzrPin.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)

		bzrPin.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
