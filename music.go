package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func Music() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "MUSIC")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (320-int16(w32))/2, 110, "MUSIC", color.RGBA{0, 100, 250, 255})
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (320-int16(w32))/2, 160, "Press any key", color.RGBA{200, 0, 0, 255})

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
