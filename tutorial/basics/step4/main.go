package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
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

	// Clear the screen to black
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	// Write "Hello" 10 pixels from the right and 50 pixels from the top, in mellon green
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 50, "Hello", color.RGBA{R: 255, G: 255, B: 0, A: 255})
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 40, 80, "Gophers!", color.RGBA{R: 255, G: 0, B: 255, A: 255})
}
