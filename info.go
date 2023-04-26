package main

import (
	_ "embed"
	"image/color"
	"time"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"

	qrcode "github.com/skip2/go-qrcode"
)

const QR_INFO_SIZE = 180

func Info() {

	qr, err := qrcode.New("https://gopherbadge.com/", qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(QR_INFO_SIZE / len(qrbytes))

	bx := (QR_INFO_SIZE - size*factor) / 2
	by := (QR_INFO_SIZE - size*factor) / 2
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

	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "SCAN")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, QR_INFO_SIZE+((WIDTH-QR_INFO_SIZE)-int16(w32))/2, 45, "SCAN", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, "ME")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, QR_INFO_SIZE+((WIDTH-QR_INFO_SIZE)-int16(w32))/2, 80, "ME", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, "Press any")
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, QR_INFO_SIZE+((WIDTH-QR_INFO_SIZE)-int16(w32))/2, 120, "Press any", colors[WHITE])
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, "button")
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, QR_INFO_SIZE+((WIDTH-QR_INFO_SIZE)-int16(w32))/2, 140, "button", colors[WHITE])
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, "to continue")
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, QR_INFO_SIZE+((WIDTH-QR_INFO_SIZE)-int16(w32))/2, 160, "to continue", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, "Visit https://gopherbadge.com/")
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, QR_INFO_SIZE+25, "Visit https://gopherbadge.com/", colors[WHITE])
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, "for more information")
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, QR_INFO_SIZE+45, "for more information", colors[WHITE])

	for {
		time.Sleep(100 * time.Millisecond)
		if !btnA.Get() || !btnB.Get() || !btnUp.Get() || !btnLeft.Get() || !btnRight.Get() || !btnDown.Get() {
			break
		}
	}

}
