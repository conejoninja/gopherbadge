package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont/freemono"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

func menu() int16 {
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	options := []string{
		"Badge",
		"GopherCon Schedule",
		"GopherCon ADventure",
		"Snake",
		"Rainbow LEDs",
		"Accelerometer",
		"Music!",
		"Game of Life",
		"Color Game",
		"Info",
	}

	bgColor := color.RGBA{109, 0, 140, 255}
	display.FillScreen(bgColor)

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 28, 35+20*i, 6, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 39, 39+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 39, 40+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 39, 41+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 40, 41+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 41, 41+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 41, 40+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 41, 39+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 40, 39+20*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &freemono.Regular12pt7b, 40, 40+20*i, options[i], color.RGBA{250, 250, 0, 255})
	}

	tinydraw.FilledCircle(&display, 28, 35, 4, color.RGBA{200, 200, 0, 255})

	released := true
	for {
		if released && !btnUp.Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 28, 35+20*selected, 4, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 28, 35+20*(selected+1), 4, bgColor)
		}
		if released && !btnDown.Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 28, 35+20*selected, 4, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 28, 35+20*(selected-1), 4, bgColor)
		}
		if released && !btnA.Get() {
			break
		}
		if btnA.Get() && btnUp.Get() && btnDown.Get() {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
