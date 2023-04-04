package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyfont/freesans"

	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/tinydraw"
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.I2C0_SCL_PIN, SDA: machine.I2C0_SDA_PIN})
	accel := lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()

	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   320,
	})

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	display.FillScreen(white)
	tinydraw.Rectangle(&display, 50, 16, 260, 16, black)
	tinydraw.Rectangle(&display, 50, 56, 260, 16, black)
	tinydraw.Rectangle(&display, 50, 96, 260, 16, black)

	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 20, 180, "MOVE the Gopher to see", black)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 32, 200, "the accelerometer in", black)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 90, 220, "action.", black)

	tinyfont.WriteLine(&display, &freesans.Regular24pt7b, 4, 40, "X:", black)
	tinyfont.WriteLine(&display, &freesans.Regular24pt7b, 4, 80, "Y:", black)
	tinyfont.WriteLine(&display, &freesans.Regular24pt7b, 4, 120, "Z:", black)

	x, y, z := accel.ReadRawAcceleration()
	for {
		x, y, z = accel.ReadRawAcceleration()
		x = x / 250
		y = y / 250
		z = z / 250
		if x > 128 {
			x = 128
		}
		if y > 128 {
			y = 128
		}
		if z > 128 {
			z = 128
		}
		if x < -128 {
			x = -128
		}
		if y < -128 {
			y = -128
		}
		if z < -128 {
			z = -128
		}
		display.FillRectangle(51, 22, 258, 6, white)
		display.FillRectangle(51, 62, 258, 6, white)
		display.FillRectangle(51, 102, 258, 6, white)
		if x < 0 {
			display.FillRectangle(179+x, 22, -x, 6, red)
		} else {
			display.FillRectangle(179, 22, x, 6, red)
		}
		if y < 0 {
			display.FillRectangle(179+y, 62, -y, 6, green)
		} else {
			display.FillRectangle(179, 62, y, 6, green)
		}
		if z < 0 {
			display.FillRectangle(179+z, 102, -z, 6, blue)
		} else {
			display.FillRectangle(179, 102, z, 6, blue)
		}

		println("X:", x, "Y:", y, "Z:", z)
		time.Sleep(time.Millisecond * 100)
		time.Sleep(50 * time.Millisecond)
	}
}
