package main

import (
	"image/color"
	"time"
)

func Leds() {
	display.EnableBacklight(false)
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	ledColors := make([]color.RGBA, 2)
	var i uint8
	for {
		ledColors[0] = getRainbowRGB(i)
		ledColors[1] = getRainbowRGB(i + 30)
		leds.WriteColors(ledColors)

		if !btnB.Get() {
			break
		}
		i += 2

		time.Sleep(50 * time.Millisecond)
	}

	ledColors[0] = color.RGBA{0, 0, 0, 255}
	ledColors[1] = color.RGBA{0, 0, 0, 255}
	leds.WriteColors(ledColors)
	time.Sleep(50 * time.Millisecond)
	ledColors[0] = color.RGBA{0, 0, 0, 255}
	ledColors[1] = color.RGBA{0, 0, 0, 255}
	leds.WriteColors(ledColors)
	time.Sleep(50 * time.Millisecond)

	display.EnableBacklight(true)
}
