package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	Red = iota
	MellonGreen
	Green
	Cyan
	Blue
	Purple
	White
	Off
)

var colors = [...]color.RGBA{
	color.RGBA{255, 0, 0, 255},     // RED
	color.RGBA{255, 255, 0, 255},   // MELLON_GREEN
	color.RGBA{0, 255, 0, 255},     // GREEN
	color.RGBA{0, 255, 255, 255},   // CYAN
	color.RGBA{0, 0, 255, 255},     // BLUE
	color.RGBA{255, 0, 255, 255},   // PURPLE
	color.RGBA{255, 255, 255, 255}, // WHITE
	color.RGBA{0, 0, 0, 255},       // OFF
}

func main() {

	// get and configure neopixels
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leds := ws2812.New(neo)
	// a color for each led on the board
	ledColors := make([]color.RGBA, 2)

	// get and configure buttons on the board
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

	c := 0
	for {
		// read buttons states
		if !btnA.Get() {
			c = Red
		}
		if !btnB.Get() {
			c = MellonGreen
		}
		if !btnUp.Get() {
			c = Green
		}
		if !btnDown.Get() {
			c = Cyan
		}
		if !btnLeft.Get() {
			c = Blue
		}
		if !btnRight.Get() {
			c = Purple
		}

		// set color for LEDs
		for i := range ledColors {
			ledColors[i] = colors[c]
		}

		leds.WriteColors(ledColors)

		time.Sleep(30 * time.Millisecond)
	}
}
