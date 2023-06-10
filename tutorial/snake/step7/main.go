package main

import (
	"image/color"
	"machine"
	"math/rand"
	"strconv"
	"time"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"

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
	body      [768][2]int16
	length    int16
	direction int16
}

var snake = Snake{
	body: [768][2]int16{
		{0, 3},
		{0, 2},
		{0, 1},
	},
	length:    3,
	direction: SnakeRight,
}

var appleX = int16(-1)
var appleY = int16(-1)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var black = color.RGBA{0, 0, 0, 255}
var white = color.RGBA{255, 255, 255, 255}

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
	createApple()
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

	if collisionWithSnake(x, y) {
		gameOver()
	}

	// draw head
	drawSnakePartial(x, y, green)

	if x == appleX && y == appleY {
		// grow our snake if we eat the apple
		snake.length++
		// create a new apple
		createApple()
	} else {
		// remove tail in case we do not eat the apple
		drawSnakePartial(snake.body[snake.length-1][0], snake.body[snake.length-1][1], black)
	}
	// move each segment coords to the next one
	for i := snake.length - 1; i > 0; i-- {
		snake.body[i][0] = snake.body[i-1][0]
		snake.body[i][1] = snake.body[i-1][1]
	}
	snake.body[0][0] = x
	snake.body[0][1] = y
}

func drawSnake() {
	for i := int16(0); i < snake.length; i++ {
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

func createApple() {
	appleX = int16(rand.Int31n(WIDTHBLOCKS))
	appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	drawSnakePartial(appleX, appleY, red)
}

func collisionWithSnake(x, y int16) bool {
	for i := int16(0); i < snake.length; i++ {
		if x == snake.body[i][0] && y == snake.body[i][1] {
			return true
		}
	}
	return false
}

func gameOver() {
	display.FillScreen(black)

	scoreStr := strconv.Itoa(int(snake.length - 3))

	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 8, 50, "GAME OVER", white)
	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 8, 100, "Press START", white)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 120, "SCORE: "+scoreStr, white)

	time.Sleep(2 * time.Second)

	for {
		if !btnA.Get() || !btnB.Get() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	// reset our game status
	snake = Snake{
		body: [768][2]int16{
			{0, 3},
			{0, 2},
			{0, 1},
		},
		length:    3,
		direction: SnakeRight,
	}
	display.FillScreen(black)
	drawSnake()
	createApple()
}
