package main

import (
	"image"
	"machine"

	draw2 "golang.org/x/image/draw"

	"tinygo.org/x/drivers/amg88xx"
	"tinygo.org/x/drivers/st7789"
)

const AMG88XX_SIZE = 8
const AMG88XX_PIXELS = 24

var display st7789.Device
var data [AMG88XX_SIZE * AMG88XX_SIZE]int16
var pixels [AMG88XX_PIXELS * AMG88XX_PIXELS]int16
var adj [16]int16

func main() {

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.I2C0_SCL_PIN, SDA: machine.I2C0_SDA_PIN})
	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	camera := amg88xx.New(machine.I2C0)
	camera.Configure(amg88xx.Config{})

	var value int16
	for {
		// get the values of the sensor in millicelsius
		camera.ReadPixels(&data)

		srcImage := image.NewRGBA(image.Rect(0, 0, 8, 8))
		for j := 0; j < AMG88XX_SIZE; j++ {
			for i := 0; i < AMG88XX_SIZE; i++ {
				value = data[63-(i+j*AMG88XX_SIZE)]
				// treat anything below 18°C as 18°C
				if value < 18000 {
					value = 0
				} else {
					value = (value - 18000) / 36
					// our color array only have 433 values, avoid getting a value that doesn't exist
					if value > 432 {
						value = 432
					}
				}
				srcImage.Set(i, j, colors[value])
			}
		}

		dst := image.NewRGBA(image.Rect(0, 0, 24, 24))
		draw2.BiLinear.Scale(dst, image.Rect(0, 0, 24, 24), srcImage, image.Rect(0, 0, 8, 8), draw2.Over, nil)

		for j := 0; j < AMG88XX_PIXELS; j++ {
			for i := 0; i < AMG88XX_PIXELS; i++ {
				display.FillRectangle(int16(40+(AMG88XX_PIXELS-i-1)*10), int16(j*10), 10, 10, dst.RGBAAt(i, j))
			}
		}
	}

}
