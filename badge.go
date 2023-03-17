package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

const (
	BLACK = iota
	WHITE
	RED
)

const (
	logoDisplayTime = 10 * time.Second
)

var colors = []color.RGBA{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{255, 0, 0, 255},
}

var rainbow []color.RGBA
var pressed uint8
var quit bool

func Badge() {
	setNameAndTitle()
	quit = false
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	for {
		logo()
		if quit {
			break
		}
		scroll("This badge", "runs", "TINYGO")
		if quit {
			break
		}
		myNameIsRainbow(YourName)
		if quit {
			break
		}
		blinkyRainbow(YourTitle1, YourTitle2)
		if quit {
			break
		}
		blinkyRainbow("Printer Party", "17-19 Marzo")
		if quit {
			break
		}
	}
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 10

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[RED])
	display.FillRectangle(0, r, WIDTH, 54, colors[RED])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	display.FillRectangle(0, HEIGHT-3*r-1, WIDTH, 2*r, colors[RED])

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "HELLO")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 34, "HELLO", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique9pt7b, "my name is")
	tinyfont.WriteLine(&display, &freesans.Oblique9pt7b, (WIDTH-int16(w32))/2, 54, "my name is", colors[WHITE])

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	if w32 < 300 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < 300 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < 300 {
				tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
			}
		}
	}

	// gophers
	tinyfont.WriteLineColors(&display, &gophers.Regular58pt, WIDTH-84, 208, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	size := 24
	if w32 < 300 {
		size = 24
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < 300 {
			size = 18
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < 300 {
				size = 12
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				size = 9
			}
		}
	}
	for i := 0; i < 230; i++ {
		if size == 24 {
			tinyfont.WriteLineColors(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 18 {
			tinyfont.WriteLineColors(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 12 {
			tinyfont.WriteLineColors(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else {
			tinyfont.WriteLineColors(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		}
		i += 2
		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func blinky(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	for i := int16(0); i < 10; i++ {
		// show black text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[BLACK])
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[BLACK])

		// repeat the other way around
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[WHITE])
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[WHITE])

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	for i := int16(0); i < 20; i++ {
		// show black text
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func scroll(topline, middleline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text, so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, topline)
	w32middle, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, middleline)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, bottomline)
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))

	display.SetScrollArea(0, 0)
	for k := 0; k < 4; k++ {
		for i := int16(319); i >= 0; i-- {

			if !btnB.Get() {
				quit = true
				break
			}
			display.SetScroll(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	display.SetScroll(0)
	display.StopScroll()
}

func setNameAndTitle() {
	if YourName == "" {
		YourName = DefaultName
	}

	if YourTitle1 == "" {
		YourTitle1 = DefaultTitle1
	}

	if YourTitle2 == "" {
		YourTitle2 = DefaultTitle2
	}
}

func logo() {
	bgColor := color.RGBA{109, 0, 140, 255}
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(bgColor)

	display.FillRectangle(6, 166, 308, 21, white)

	tinydraw.FilledCircle(&display, 282, 130, 9, white)
	tinydraw.Line(&display, 259, 110, 298, 149, bgColor)
	tinydraw.Line(&display, 260, 110, 299, 149, bgColor)
	tinydraw.Line(&display, 261, 110, 300, 149, bgColor)
	tinydraw.Line(&display, 262, 110, 301, 149, bgColor)
	tinydraw.Line(&display, 263, 110, 302, 149, bgColor)
	tinydraw.Line(&display, 264, 110, 303, 149, bgColor)
	tinydraw.Line(&display, 265, 110, 304, 149, bgColor)

	display.FillRectangle(250, 98, 11, 63, white)
	display.FillRectangle(250, 98, 44, 11, white)

	display.FillRectangle(270, 150, 44, 11, white)
	display.FillRectangle(303, 98, 11, 63, white)

	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 6, 109, "Purple", white)
	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 6, 153, "Hardware by", white)

	time.Sleep(logoDisplayTime)
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
