package main

import (
	"machine"
	"time"
)

func main() {

	// get the pin that represents that LED on the back of the device D13
	led := machine.LED

	// configuring the LED pin mode for output
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// configure the A button on the device
	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		// was the start button pressed?
		// if so, turn on the LED, otherwise turn it off
		if btnA.Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
