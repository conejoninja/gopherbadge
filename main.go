package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/lis3dh"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/drivers/ws2812"
)

var display st7789.Device
var leds ws2812.Device
var accel lis3dh.Device
var bzrPin machine.Pin
var btnA, btnB, btnUp, btnLeft, btnDown, btnRight machine.Pin

const (
	BLACK = iota
	WHITE
	RED
	SNAKE
	TEXT
	ORANGE
	PURPLE
)

var colors = []color.RGBA{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{250, 0, 0, 255},
	color.RGBA{0, 200, 0, 255},
	color.RGBA{160, 160, 160, 255},
	color.RGBA{255, 153, 51, 255},
	color.RGBA{153, 51, 255, 255},
}

var snakeGame = NewSnakeGame()

func main() {

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	machine.I2C0.Configure(machine.I2CConfig{
		SCL: machine.I2C0_SCL_PIN,
		SDA: machine.I2C0_SDA_PIN,
	})
	accel = lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()

	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)

	bzrPin = machine.SPEAKER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	btnA = machine.BUTTON_A
	btnB = machine.BUTTON_B
	btnUp = machine.BUTTON_UP
	btnLeft = machine.BUTTON_LEFT
	btnDown = machine.BUTTON_DOWN
	btnRight = machine.BUTTON_RIGHT
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInput})

	black := color.RGBA{0, 0, 0, 255}
	display.FillScreen(black)
	setCustomData()

	Info()

	for {
		switch menu() {
		case 0:
			Badge()
			break
		case 1:
			schedule(0, 0)
			break
		case 2:
			adventure()
			break
		case 3:
			snakeGame.Loop()
			break
		case 4:
			Leds()
			break
		case 5:
			Accel3D()
			break
		case 6:
			Music()
			break
		case 7:
			GameOfLife()
			break
		case 8:
			ColorGame()
			break
		case 9:
			Info()
			break
		default:
			break
		}
		println("LOOP")
		time.Sleep(1 * time.Second)
	}

}
