package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
)

const (
	WIDTHBLOCKS  = 32
	HEIGHTBLOCKS = 24
)

var display st7789.Device
var btnA, btnB, btnUp, btnLeft, btnDown, btnRight machine.Pin

func main() {
	// Setup the SPI connection of the GopherBadge
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	// Create a new display device
	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	// Setup the buttons
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

	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// fill the whole screen with black
	display.FillScreen(black)

	x := int16(WIDTHBLOCKS / 2)
	y := int16(HEIGHTBLOCKS / 2)
	for {

		// "clear" our previous snake position
		display.FillRectangle(x*10, y*10, 10, 10, green)

		// Check if any direction button is pressed
		if !btnLeft.Get() {
			x--
		}
		if !btnUp.Get() {
			y--
		}
		if !btnDown.Get() {
			y++
		}
		if !btnRight.Get() {
			x++
		}

		// control it doesn't get out of bounds
		if x >= WIDTHBLOCKS {
			x = 0
		}
		if x < 0 {
			x = WIDTHBLOCKS - 1
		}
		if y >= HEIGHTBLOCKS {
			y = 0
		}
		if y < 0 {
			y = HEIGHTBLOCKS - 1
		}

		// draw our little snake-rectangle in their new position
		display.FillRectangle(x*10, y*10, 10, 10, green)

		time.Sleep(100 * time.Millisecond)
	}
}
