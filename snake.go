package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"

	"tinygo.org/x/tinyfont"
)

const (
	GameSplash = iota
	GameStart
	GamePlay
	GameOver
	GameQuit
)

const (
	SnakeUp = iota
	SnakeDown
	SnakeLeft
	SnakeRight
)

const (
	WIDTHBLOCKS  = 32
	HEIGHTBLOCKS = 24
)

var (
	// Those variable are there for a more easy reading of the apple shape.
	re = colors[RED]   // red
	bk = colors[BLACK] // background
	gr = colors[SNAKE] // green

	// The array is split for a visual purpose too.
	appleBuf = []color.RGBA{
		bk, bk, bk, bk, bk, gr, gr, gr, bk, bk,
		bk, bk, bk, bk, gr, gr, gr, bk, bk, bk,
		bk, bk, bk, re, gr, gr, re, bk, bk, bk,
		bk, bk, re, re, re, re, re, re, bk, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, bk, re, re, re, re, re, re, bk, bk,
		bk, bk, bk, re, re, re, re, bk, bk, bk,
		bk, bk, bk, bk, bk, bk, bk, bk, bk, bk,
	}
)

type Snake struct {
	body      [768][2]int16
	length    int16
	direction int16
}

type SnakeGame struct {
	snake          Snake
	appleX, appleY int16
	status         uint8
	score          int
	frame, delay   int
}

var splashed = false
var scoreStr string

func NewSnakeGame() *SnakeGame {
	return &SnakeGame{
		snake: Snake{
			body: [768][2]int16{
				{0, 3},
				{0, 2},
				{0, 1},
			},
			length:    3,
			direction: SnakeLeft,
		},
		appleX: 5,
		appleY: 5,
		status: GameSplash,
		delay:  120,
	}
}

func (g *SnakeGame) Splash() {
	if !splashed {
		g.splash()
		splashed = true
	}
}

func (g *SnakeGame) Start() {
	display.FillScreen(bk)

	g.initSnake()
	g.drawSnake()
	g.createApple()

	g.status = GamePlay
}

func (g *SnakeGame) Play(direction int) {
	if direction != -1 && ((g.snake.direction == SnakeUp && direction != SnakeDown) ||
		(g.snake.direction == SnakeDown && direction != SnakeUp) ||
		(g.snake.direction == SnakeLeft && direction != SnakeRight) ||
		(g.snake.direction == SnakeRight && direction != SnakeLeft)) {
		g.snake.direction = int16(direction)
	}

	g.moveSnake()
}

func (g *SnakeGame) Over() {
	display.FillScreen(bk)
	splashed = false

	g.status = GameOver
}

func (g *SnakeGame) splash() {
	display.FillScreen(bk)

	logo := `
      ___            ___            ___     
     /  /\          /  /\          /  /\    
    /  /::\        /  /::|        /  /::\   
   /__/:/\:\      /  /:|:|       /  /:/\:\  
  _\_ \:\ \:\    /  /:/|:|__    /  /::\ \:\ 
 /__/\ \:\ \:\  /__/:/ |:| /\  /__/:/\:\_\:\
 \  \:\ \:\_\/  \__\/  |:|/:/  \__\/  \:\/:/
  \  \:\_\:\        |  |:/:/        \__\::/ 
   \  \:\/:/        |__|::/         /  /:/  
    \  \::/         /__/:/         /__/:/   
     \__\/          \__\/          \__\/    
               ___            ___     
              /  /\          /  /\    
             /  /:/         /  /::\   
            /  /:/         /  /:/\:\  
           /  /::\____    /  /::\ \:\ 
          /__/:/\:::::\  /__/:/\:\ \:\
          \__\/~|:|~~~~  \  \:\ \:\_\/
             |  |:|       \  \:\ \:\  
             |  |:|        \  \:\_\/  
             |__|:|         \  \:\    
              \__\|          \__\/     
`
	for i, line := range strings.Split(strings.TrimSuffix(logo, "\n"), "\n") {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 0, int16(-6+i*11), line+"\n", gr)
	}

	tinyfont.WriteLine(&display, &freesans.Regular18pt7b, 30, 130, "Press A to start", colors[RED])

	if g.score > 0 {
		scoreStr = strconv.Itoa(g.score)
		tinyfont.WriteLineRotated(&display, &freesans.Regular12pt7b, 300, 200, "SCORE: "+scoreStr, colors[TEXT], tinyfont.ROTATION_270)
	}
}

func (g *SnakeGame) initSnake() {
	g.snake.body[0][0] = 0
	g.snake.body[0][1] = 3
	g.snake.body[1][0] = 0
	g.snake.body[1][1] = 2
	g.snake.body[2][0] = 0
	g.snake.body[2][1] = 1

	g.snake.length = 3
	g.snake.direction = SnakeRight
}

func (g *SnakeGame) collisionWithSnake(x, y int16) bool {
	for i := int16(0); i < g.snake.length; i++ {
		if x == g.snake.body[i][0] && y == g.snake.body[i][1] {
			return true
		}
	}
	return false
}

func (g *SnakeGame) createApple() {
	g.appleX = int16(rand.Int31n(WIDTHBLOCKS))
	g.appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	for g.collisionWithSnake(g.appleX, g.appleY) {
		g.appleX = int16(rand.Int31n(WIDTHBLOCKS))
		g.appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	}
	g.drawApple(g.appleX, g.appleY)
}

func (g *SnakeGame) moveSnake() {
	x := g.snake.body[0][0]
	y := g.snake.body[0][1]

	switch g.snake.direction {
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

	if g.collisionWithSnake(x, y) {
		g.score = int(g.snake.length - 3)
		g.Over()

		return
	}

	// draw head
	g.drawSnakePartial(x, y, colors[SNAKE])
	if x == g.appleX && y == g.appleY {
		g.snake.length++
		g.createApple()
	} else {
		// remove tail
		g.drawSnakePartial(g.snake.body[g.snake.length-1][0], g.snake.body[g.snake.length-1][1], colors[BLACK])
	}
	for i := g.snake.length - 1; i > 0; i-- {
		g.snake.body[i][0] = g.snake.body[i-1][0]
		g.snake.body[i][1] = g.snake.body[i-1][1]
	}
	g.snake.body[0][0] = x
	g.snake.body[0][1] = y
}

func (g *SnakeGame) drawApple(x, y int16) {
	display.FillRectangleWithBuffer(10*x, 10*y, 10, 10, appleBuf)
}

func (g *SnakeGame) drawSnake() {
	for i := int16(0); i < g.snake.length; i++ {
		g.drawSnakePartial(g.snake.body[i][0], g.snake.body[i][1], colors[SNAKE])
	}
}

func (g *SnakeGame) drawSnakePartial(x, y int16, c color.RGBA) {
	display.FillRectangle(10*x, 10*y, 9, 9, c)
}

func (g *SnakeGame) Loop() {
	g.status = GameSplash
	splashed = false
	for {
		g.update()
		if g.status == GameQuit {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (g *SnakeGame) update() {
	switch g.status {
	case GameSplash:
		g.Splash()
		if !btnA.Get() {
			g.Start()
		}
		if !btnB.Get() {
			g.status = GameOver
		}
		break

	case GamePlay:
		switch {
		case !btnB.Get():
			g.Over()
			break
		case !btnRight.Get():
			g.Play(SnakeRight)
			break

		case !btnLeft.Get():
			g.Play(SnakeLeft)
			break

		case !btnDown.Get():
			g.Play(SnakeDown)
			break

		case !btnUp.Get():
			g.Play(SnakeUp)
			break

		default:
			g.Play(-1)
			break
		}
		break
	case GameQuit:
	case GameOver:
		g.Splash()

		if !btnA.Get() {
			g.Start()
		}
		if !btnB.Get() {
			g.status = GameQuit
		}
	}
}
