package main

import (
	_ "embed"
	"image/color"
	"strconv"
	"time"
	"unsafe"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"

	qrcode "github.com/skip2/go-qrcode"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

const (
	logoDisplayTime = 10 * time.Second
)

var rainbow []color.RGBA
var quit bool

//go:embed logo.bin
var badgeLogo string

func Badge() {
	quit = false
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	for {
		showLogoBin()
		if quit {
			break
		}
		logoPurpleHardware()
		if quit {
			break
		}
		myNameIsRainbow(YourName)
		if quit {
			break
		}
		blinkyRainbow(YourTitleA1, YourTitleA2)
		if quit {
			break
		}
		scroll(YourMarqueeTop, YourMarqueeMiddle, YourMarqueeBottom)
		if quit {
			break
		}
		QR(YourQRText)
		if quit {
			break
		}
		blinkyRainbow(YourTitleB1, YourTitleB2)
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
	w32, size := getFontWidthSize(name)
	if size == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	} else if size == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	} else if size == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, colors[BLACK])
	}

	// gophers
	tinyfont.WriteLineColors(&display, &gophers.Regular58pt, WIDTH-84, 208, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, size := getFontWidthSize(name)
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
	w32top, sizetop := getFontWidthSize(topline)
	w32bottom, sizebottom := getFontWidthSize(bottomline)
	for i := int16(0); i < 10; i++ {
		// show black text
		if sizetop == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[BLACK])
		} else if sizetop == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[BLACK])
		} else if sizetop == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[BLACK])
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[BLACK])
		}
		if sizebottom == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[BLACK])
		} else if sizebottom == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[BLACK])
		} else if sizebottom == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[BLACK])
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[BLACK])
		}

		// repeat the other way around
		if sizetop == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[WHITE])
		} else if sizetop == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[WHITE])
		} else if sizetop == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[WHITE])
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 100, topline, colors[WHITE])
		}
		if sizebottom == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[WHITE])
		} else if sizebottom == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[WHITE])
		} else if sizebottom == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[WHITE])
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, colors[WHITE])
		}

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, sizetop := getFontWidthSize(topline)
	w32bottom, sizebottom := getFontWidthSize(bottomline)
	for i := int16(0); i < 20; i++ {
		// show black text
		if sizetop == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 100, topline, getRainbowRGB(uint8(i*12)))
		}
		if sizebottom == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(uint8(i*12)))
		}

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func scroll(topline, middleline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text, so we could center them later
	w32top, sizetop := getFontWidthSize(topline)
	w32middle, sizemiddle := getFontWidthSize(middleline)
	w32bottom, sizebottom := getFontWidthSize(bottomline)

	if sizetop == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else if sizetop == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else if sizetop == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 70, topline, getRainbowRGB(200))
	}
	if sizemiddle == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else if sizemiddle == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else if sizemiddle == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32middle))/2, 120, middleline, getRainbowRGB(80))
	}
	if sizebottom == 24 {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else if sizebottom == 18 {
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else if sizebottom == 12 {
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 200, bottomline, getRainbowRGB(120))
	}

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

func logoPurpleHardware() {
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

	time.Sleep(logoDisplayTime / 3)
	if !btnB.Get() {
		quit = true
	}
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

func showLogoBin() {
	var row = []color.RGBA{}
	row = make([]color.RGBA, WIDTH)
	unsafeBadgeLogo := unsafe.Slice(unsafe.StringData(badgeLogo), len(badgeLogo))
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			values, err := strconv.ParseUint(string(unsafeBadgeLogo[6*(WIDTH*i+j):6*(WIDTH*i+j+1)]), 16, 32)

			if err != nil {
				println(err)
				//return RGB{}, err
			}

			row[j] = color.RGBA{
				R: uint8(values >> 16),
				G: uint8((values >> 8) & 0xFF),
				B: uint8(values & 0xFF),
				A: 255,
			}
		}
		display.FillRectangleWithBuffer(0, int16(i), WIDTH, 1, row)
	}
	time.Sleep(logoDisplayTime)
	if !btnB.Get() {
		quit = true
	}
}

func QR(msg string) {
	qr, err := qrcode.New(msg, qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor) / 2
	by := (HEIGHT - size*factor) / 2
	display.FillScreen(color.RGBA{109, 0, 140, 255})
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[0])
			} else {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[1])
			}
		}
	}

	time.Sleep(logoDisplayTime)
	if !btnB.Get() {
		quit = true
	}
}

func getFontWidthSize(text string) (w32 uint32, size byte) {
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, text)
	size = 24
	if w32 < 300 {
		size = 24
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, text)
		if w32 < 300 {
			size = 18
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, text)
			if w32 < 300 {
				size = 12
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, text)
				size = 9
			}
		}
	}
	return
}
