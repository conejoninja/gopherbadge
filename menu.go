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
		"Reflex Tester",
		"GopherCon Adventure",
		"Snake",
		"Rainbow LEDs",
		"Pacman",
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

	btnUpPressed := false
	btnDownPressed := false
	btnAPressed := false

	for {
		// Detect button press on transition from released to pressed
		if !btnUp.Get() {
			if !btnUpPressed {
				prevSelected := selected
				if selected > 0 {
					selected--
				} else {
					selected = numOpts - 1 // Wrap to bottom
				}
				tinydraw.FilledCircle(&display, 28, 35+20*selected, 4, color.RGBA{200, 200, 0, 255})
				tinydraw.FilledCircle(&display, 28, 35+20*prevSelected, 4, bgColor)
			}
			btnUpPressed = true
		} else {
			btnUpPressed = false
		}

		if !btnDown.Get() {
			if !btnDownPressed {
				prevSelected := selected
				if selected < (numOpts - 1) {
					selected++
				} else {
					selected = 0 // Wrap to top
				}
				tinydraw.FilledCircle(&display, 28, 35+20*selected, 4, color.RGBA{200, 200, 0, 255})
				tinydraw.FilledCircle(&display, 28, 35+20*prevSelected, 4, bgColor)
			}
			btnDownPressed = true
		} else {
			btnDownPressed = false
		}

		if !btnA.Get() {
			if !btnAPressed {
				break
			}
			btnAPressed = true
		} else {
			btnAPressed = false
		}

		time.Sleep(100 * time.Millisecond)
	}
	return selected
}
