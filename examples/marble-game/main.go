package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/tinydraw"

	"github.com/go-gl/mathgl/mgl32"
	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/drivers/st7789"
)

const WIDTH = 320
const HEIGHT = 240

var display st7789.Device
var accel lis3dh.Device

var ballX = float32(WIDTH / 2)
var ballY = float32(HEIGHT / 2)
var goalX = int16(-1)
var goalY = int16(-1)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var black = color.RGBA{0, 0, 0, 255}
var white = color.RGBA{255, 255, 255, 255}

var (
	// The array is split for a visual purpose too.
	ballBuf = []color.RGBA{
		black, black, white, white, white, black, black,
		black, white, white, white, white, white, black,
		white, white, white, white, white, white, white,
		white, white, white, white, white, white, white,
		white, white, white, white, white, white, white,
		black, white, white, white, white, white, black,
		black, black, white, white, white, black, black,
	}
)

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
		Rotation: st7789.ROTATION_90,
		Height:   320,
	})

	accel = lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	println("connected:", accel.Connected())

	var dx, dy int16

	goal()
	for {
		display.FillRectangle(int16(ballX-4), int16(ballY-4), 7, 7, black)
		xi, yi, zi, _ := accel.ReadAcceleration()
		vec := mgl32.Vec3{
			float32(xi) / 1000000,
			float32(yi) / 1000000,
			float32(zi) / 1000000,
		}
		vec = vec.Normalize()

		ballX -= vec[0] * 10
		ballY -= vec[1] * 10
		if ballX <= 4 {
			ballX = 4
		}
		if ballY <= 4 {
			ballY = 4
		}
		if ballX >= WIDTH-4 {
			ballX = WIDTH - 4
		}
		if ballY >= HEIGHT-4 {
			ballY = HEIGHT - 4
		}

		dx = int16(ballX) - goalX
		dy = int16(ballY) - goalY
		if dx < 7 && dx > -7 && dy < 7 && dy > -7 {
			goal()
		}

		display.FillRectangleWithBuffer(int16(ballX-4), int16(ballY-4), 7, 7, ballBuf)
		time.Sleep(time.Second / 20)
	}

}

func goal() {
	tinydraw.FilledCircle(&display, goalX, goalY, 3, black)

	goalX = 7 + int16(rand.Int31n(WIDTH-20))
	goalY = 7 + int16(rand.Int31n(HEIGHT-20))
	tinydraw.FilledCircle(&display, goalX, goalY, 3, green)
}
