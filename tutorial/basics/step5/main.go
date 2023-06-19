package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"

	"tinygo.org/x/tinydraw"
)

func main() {

	// Setup the screen pins
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

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

	display.FillScreen(color.RGBA{255, 255, 255, 255})

	circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	ring := color.RGBA{200, 0, 0, 255}

	// Clear the display to white
	display.FillScreen(white)

	// Draw blue circles to represent each of the buttons
	tinydraw.FilledCircle(&display, 25, 120, 14, circle) // LEFT
	tinydraw.FilledCircle(&display, 95, 120, 14, circle) // RIGHT
	tinydraw.FilledCircle(&display, 60, 85, 14, circle)  // UP
	tinydraw.FilledCircle(&display, 60, 155, 14, circle) // DOWN

	tinydraw.FilledCircle(&display, 260, 120, 14, circle) // B
	tinydraw.FilledCircle(&display, 295, 85, 14, circle)  // A

	for {
		if !btnA.Get() {
			tinydraw.Circle(&display, 295, 85, 16, ring)
		} else {
			tinydraw.Circle(&display, 295, 85, 16, white)
		}
		if !btnB.Get() {
			tinydraw.Circle(&display, 260, 120, 16, ring)
		} else {
			tinydraw.Circle(&display, 260, 120, 16, white)
		}
		if !btnLeft.Get() {
			tinydraw.Circle(&display, 25, 120, 16, ring)
		} else {
			tinydraw.Circle(&display, 25, 120, 16, white)
		}
		if !btnRight.Get() {
			tinydraw.Circle(&display, 95, 120, 16, ring)
		} else {
			tinydraw.Circle(&display, 95, 120, 16, white)
		}
		if !btnUp.Get() {
			tinydraw.Circle(&display, 60, 85, 16, ring)
		} else {
			tinydraw.Circle(&display, 60, 85, 16, white)
		}
		if !btnDown.Get() {
			tinydraw.Circle(&display, 60, 155, 16, ring)
		} else {
			tinydraw.Circle(&display, 60, 155, 16, white)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
