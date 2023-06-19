package main

import (
	"image/color"
	"machine"
	"time"
	"tinygo.org/x/drivers/st7789"
)

var display st7789.Device

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

	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// fill the whole screen with black
	display.FillScreen(black)

	w, h := display.Size()
	// draw a green 10x10 rectangle at the middle of the screen
	display.FillRectangle((w-10)/2, (h-10)/2, 10, 10, green)
	for {
		time.Sleep(1 * time.Second)
	}
}
