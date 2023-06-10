package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
)

const (
	WIDTHBLOCKS  = 32
	HEIGHTBLOCKS = 24
)

const (
	SnakeUp = iota
	SnakeDown
	SnakeLeft
	SnakeRight
)

var display st7789.Device
var btnA, btnB, btnUp, btnLeft, btnDown, btnRight machine.Pin

type Snake struct {
	body      [3][2]int16
	direction int16
}

var snake = Snake{
	body: [3][2]int16{
		{0, 3},
		{0, 2},
		{0, 1},
	},
	direction: SnakeRight,
}

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var black = color.RGBA{0, 0, 0, 255}

func main() {
	// Setup the SPI connection of the GopherBadge
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	// Create a new display device
	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   320,
	})

	// Setup the buttons
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

	// fill the whole screen with black
	display.FillScreen(black)

	drawSnake()
	for {
		switch {
		// add some checks so the snake doesn't go backwards
		case !btnLeft.Get():
			if snake.direction != SnakeRight {
				snake.direction = SnakeLeft
			}
		case !btnUp.Get():
			if snake.direction != SnakeDown {
				snake.direction = SnakeUp
			}
		case !btnDown.Get():
			if snake.direction != SnakeUp {
				snake.direction = SnakeDown
			}
		case !btnRight.Get():
			if snake.direction != SnakeLeft {
				snake.direction = SnakeRight
			}
		}

		moveSnake()
		time.Sleep(100 * time.Millisecond)
	}
}

func moveSnake() {
	// get the coords of the head
	x := snake.body[0][0]
	y := snake.body[0][1]

	switch snake.direction {
	case SnakeLeft:
		x--
		break
	case SnakeUp:
		y--
		break
	case SnakeDown:
		y++
		break
	case SnakeRight:
		x++
		break
	}
	// check the bounds
	if x >= WIDTHBLOCKS {
		x = 0
	}
	if x < 0 {
		x = WIDTHBLOCKS - 1
	}
	if y >= HEIGHTBLOCKS {
		y = 0
	}
	if y < 0 {
		y = HEIGHTBLOCKS - 1
	}

	// draw head
	drawSnakePartial(x, y, green)

	// remove tail
	drawSnakePartial(snake.body[2][0], snake.body[2][1], black)

	// move each segment coords to the next one
	for i := 2; i > 0; i-- {
		snake.body[i][0] = snake.body[i-1][0]
		snake.body[i][1] = snake.body[i-1][1]
	}
	snake.body[0][0] = x
	snake.body[0][1] = y
}

func drawSnake() {
	for i := int16(0); i < 3; i++ {
		drawSnakePartial(snake.body[i][0], snake.body[i][1], green)
	}
}

func drawSnakePartial(x, y int16, c color.RGBA) {
	modY := int16(9)
	if y == 12 {
		modY = 8
	}
	// we changed the size of 10 to 9, so a black border is shown
	// around each segment of the snake
	display.FillRectangle(10*x, 10*y, 9, modY, c)
}
