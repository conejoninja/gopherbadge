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
	red = colors[RED]    // red
	blk = colors[BLACK]  // background
	grn = colors[SNAKE]  // green
	prl = colors[PURPLE] // purple

	// The array is split for a visual purpose too.
	appleBuf = []color.RGBA{
		blk, blk, blk, blk, blk, grn, grn, grn, blk, blk,
		blk, blk, blk, blk, grn, grn, grn, blk, blk, blk,
		blk, blk, blk, red, grn, grn, red, blk, blk, blk,
		blk, blk, red, red, red, red, red, red, blk, blk,
		blk, red, red, red, red, red, red, red, red, blk,
		blk, red, red, red, red, red, red, red, red, blk,
		blk, red, red, red, red, red, red, red, red, blk,
		blk, blk, red, red, red, red, red, red, blk, blk,
		blk, blk, blk, red, red, red, red, blk, blk, blk,
		blk, blk, blk, blk, blk, blk, blk, blk, blk, blk,
	}

	// The array is split for a visual purpose too.
	poisonAppleBuf = []color.RGBA{
		blk, blk, blk, blk, blk, grn, grn, grn, blk, blk,
		blk, blk, blk, blk, grn, grn, grn, blk, blk, blk,
		blk, blk, blk, prl, grn, grn, prl, blk, blk, blk,
		blk, blk, prl, prl, prl, prl, prl, prl, blk, blk,
		blk, prl, prl, prl, prl, prl, prl, prl, prl, blk,
		blk, prl, prl, prl, prl, prl, prl, prl, prl, blk,
		blk, prl, prl, prl, prl, prl, prl, prl, prl, blk,
		blk, blk, prl, prl, prl, prl, prl, prl, blk, blk,
		blk, blk, blk, prl, prl, prl, prl, blk, blk, blk,
		blk, blk, blk, blk, blk, blk, blk, blk, blk, blk,
	}
)

type Snake struct {
	body      [768][2]int16
	length    int16
	direction int16
}

type SnakeGame struct {
	snake              Snake
	appleX, appleY     int16
	poisonX, poisonY   int16
	status             uint8
	frame, delay       int
	sparkleX, sparkleY int16
	eatCounter         int
}

var splashed = false
var scoreStr string
var snakeSoundEnabled = false // Sound mode selection (default OFF)

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
		appleX:  5,
		appleY:  5,
		poisonX: 5,
		poisonY: 5,
		status:  GameSplash,
		delay:   120,
	}
}

func (g *SnakeGame) Splash() {
	if !splashed {
		g.splash()
		splashed = true
	}
}

func (g *SnakeGame) Start() {
	display.FillScreen(blk)

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
	splashed = false
	g.status = GameOver
}

func (g *SnakeGame) showGameOver() {
	display.FillScreen(blk)

	// Display "GAME OVER" in large text
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 20, 80, "GAME OVER", colors[RED])

	// Display the score (snake length is the score)
	scoreText := "Score: " + strconv.Itoa(int(g.snake.length))
	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, scoreText)
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 130, scoreText, colors[YELLOW])

	// Display instructions
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 180, "Press A to retry", colors[WHITE])
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 210, "Press B to exit", colors[TEXT])
}

func (g *SnakeGame) splash() {
	display.FillScreen(blk)

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
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 0, int16(-6+i*11), line+"\n", grn)
	}

	scoreStr = strconv.Itoa(int(g.snake.length))
	tinyfont.WriteLineRotated(&display, &freesans.Regular12pt7b, 300, 200, "SCORE: "+scoreStr, colors[TEXT], tinyfont.ROTATION_270)

	// Sound mode menu (selected=0 is Sound OFF, selected=1 is Sound ON)
	selected := 0 // Default to Sound OFF
	if snakeSoundEnabled {
		selected = 1
	}
	g.drawSoundMenu(selected)

	btnUpPressed := false
	btnDownPressed := false
	btnAPressed := true // Start pressed to avoid immediate selection

	for {
		// Detect button press on transition from released to pressed
		if !btnUp.Get() {
			if !btnUpPressed && selected > 0 {
				selected--
				g.drawSoundMenu(selected)
			}
			btnUpPressed = true
		} else {
			btnUpPressed = false
		}

		if !btnDown.Get() {
			if !btnDownPressed && selected < 1 {
				selected++
				g.drawSoundMenu(selected)
			}
			btnDownPressed = true
		} else {
			btnDownPressed = false
		}

		if !btnA.Get() {
			if !btnAPressed {
				snakeSoundEnabled = (selected == 1) // selected=0 is Sound OFF, selected=1 is Sound ON
				break
			}
		} else {
			btnAPressed = false
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (g *SnakeGame) drawSoundMenu(selected int) {
	// Clear menu area
	//display.FillRectangle(30, 115, 260, 80, blk)

	// Draw menu options (selected=0 is Sound OFF, selected=1 is Sound ON)
	soundOnColor := colors[WHITE]
	soundOffColor := colors[WHITE]
	if selected == 0 {
		soundOffColor = colors[YELLOW]
	} else {
		soundOnColor = colors[YELLOW]
	}

	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 140, "Sound OFF", soundOffColor)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 170, "Sound ON", soundOnColor)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 30, 210, "Press A to start", colors[RED])
}

// flashNeopixels briefly flashes both Neopixels with the given color
func flashNeopixels(c color.RGBA) {
	off := color.RGBA{0, 0, 0, 255}
	ledColors := []color.RGBA{c, c}
	leds.WriteColors(ledColors)
	time.Sleep(100 * time.Millisecond)
	// Turn off LEDs - write twice like leds.go does for reliability
	ledColors[0] = off
	ledColors[1] = off
	leds.WriteColors(ledColors)
	time.Sleep(50 * time.Millisecond)
	leds.WriteColors(ledColors)
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
	for g.collisionWithSnake(g.appleX, g.appleY) || (g.appleX == g.sparkleX && g.appleY == g.sparkleY && g.eatCounter > 0) {
		g.appleX = int16(rand.Int31n(WIDTHBLOCKS))
		g.appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	}
	g.drawApple(g.appleX, g.appleY)
	// Poison apple
	g.poisonX = int16(rand.Int31n(WIDTHBLOCKS))
	g.poisonY = int16(rand.Int31n(HEIGHTBLOCKS))
	for g.collisionWithSnake(g.poisonX, g.poisonY) || (g.poisonX == g.sparkleX && g.poisonY == g.sparkleY && g.eatCounter > 0) {
		g.poisonX = int16(rand.Int31n(WIDTHBLOCKS))
		g.poisonY = int16(rand.Int31n(HEIGHTBLOCKS))
	}
	g.drawPoisonApple(g.poisonX, g.poisonY)
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
		g.Over()

		return
	}

	// draw head
	snakeColor := g.getSnakeColor()
	g.drawSnakePartial(x, y, snakeColor)
	if x == g.appleX && y == g.appleY {
		g.snake.length++
		// Clear any previous unfinished sparkle animation before starting a new one
		if g.eatCounter > 0 {
			g.clearSparkles(g.sparkleX, g.sparkleY)
		}
		// Trigger sparkle animation and color flash at apple position
		g.sparkleX = g.appleX
		g.sparkleY = g.appleY
		g.eatCounter = 8 // 8 frames = 800ms for both sparkles and color flash
		// remove poison apple
		g.drawSnakePartial(g.poisonX, g.poisonY, colors[BLACK])
		// create new apple and poison apple
		g.createApple()
		if snakeSoundEnabled {
			go tone(1500)
		} else {
			// TODO: fix Neopixel issue - LEDs sometimes stay on
			// go flashNeopixels(color.RGBA{0, 255, 0, 255}) // Green for good apple
		}
	} else if x == g.poisonX && y == g.poisonY {
		// Clear any previous unfinished sparkle animation
		if g.eatCounter > 0 {
			g.clearSparkles(g.sparkleX, g.sparkleY)
			g.eatCounter = 0
		}
		// Check if snake is too short to survive eating poison
		if g.snake.length <= 1 {
			g.Over()
			return
		}
		// remove tail (safe now since length > 1)
		g.drawSnakePartial(g.snake.body[g.snake.length-2][0], g.snake.body[g.snake.length-2][1], colors[BLACK])
		g.drawSnakePartial(g.snake.body[g.snake.length-1][0], g.snake.body[g.snake.length-1][1], colors[BLACK])
		g.snake.length--
		// remove regular apple
		g.drawSnakePartial(g.appleX, g.appleY, colors[BLACK])
		// create new apple and poison apple
		g.createApple()
		if snakeSoundEnabled {
			go tone(500)
		} else {
			// TODO: fix Neopixel issue - LEDs sometimes stay on
			// go flashNeopixels(color.RGBA{255, 0, 0, 255}) // Red for poison apple
		}
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

func (g *SnakeGame) drawPoisonApple(x, y int16) {
	display.FillRectangleWithBuffer(10*x, 10*y, 10, 10, poisonAppleBuf)
}

func (g *SnakeGame) getSnakeColor() color.RGBA {
	if g.eatCounter > 0 {
		// Alternate between normal green and yellow every frame
		if g.eatCounter%2 == 0 {
			return colors[YELLOW]
		}
	}
	return colors[SNAKE] // normal green
}

func (g *SnakeGame) drawSnake() {
	snakeColor := g.getSnakeColor()
	for i := int16(0); i < g.snake.length; i++ {
		g.drawSnakePartial(g.snake.body[i][0], g.snake.body[i][1], snakeColor)
	}
}

func (g *SnakeGame) drawSnakePartial(x, y int16, c color.RGBA) {
	display.FillRectangle(10*x, 10*y, 9, 9, c)
}

func (g *SnakeGame) drawSparkles(x, y int16, frame int) {
	// Define sparkle positions around the apple (relative offsets in pixels)
	sparkleOffsets := [][2]int16{
		{-2, -2}, {5, -2}, {12, -2}, // top row
		{-2, 5}, {12, 5}, // middle row (left and right)
		{-2, 12}, {5, 12}, {12, 12}, // bottom row
	}

	// Cycle through different sparkle colors based on frame
	var sparkleColor color.RGBA
	switch frame % 4 {
	case 0:
		sparkleColor = colors[WHITE]
	case 1:
		sparkleColor = color.RGBA{255, 255, 0, 255} // yellow
	case 2:
		sparkleColor = colors[ORANGE]
	case 3:
		sparkleColor = color.RGBA{255, 200, 200, 255} // light pink
	}

	baseX := 10 * x
	baseY := 10 * y

	// Draw sparkles at different positions based on frame for animation effect
	for i, offset := range sparkleOffsets {
		px := baseX + offset[0]
		py := baseY + offset[1]
		// Make sparkles appear/disappear in a pattern
		if (frame+i)%2 == 0 {
			// Draw a small 2x2 sparkle
			display.FillRectangle(px, py, 2, 2, sparkleColor)
		} else {
			// Clear the sparkle when it's not visible
			display.FillRectangle(px, py, 2, 2, colors[BLACK])
		}
	}
}

func (g *SnakeGame) clearSparkles(x, y int16) {
	// Clear all sparkle positions
	sparkleOffsets := [][2]int16{
		{-2, -2}, {5, -2}, {12, -2}, // top row
		{-2, 5}, {12, 5}, // middle row (left and right)
		{-2, 12}, {5, 12}, {12, 12}, // bottom row
	}

	baseX := 10 * x
	baseY := 10 * y

	for _, offset := range sparkleOffsets {
		px := baseX + offset[0]
		py := baseY + offset[1]
		display.FillRectangle(px, py, 2, 2, colors[BLACK])
	}
}

func (g *SnakeGame) updateEatAnimation() {
	if g.eatCounter > 0 {
		// Update sparkles
		g.drawSparkles(g.sparkleX, g.sparkleY, g.eatCounter)
		// Update snake color flash
		g.drawSnake()
		g.eatCounter--
		// Clear sparkles and reset snake color when animation ends
		if g.eatCounter == 0 {
			g.clearSparkles(g.sparkleX, g.sparkleY)
			g.drawSnake()
		}
	}
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
		// Update eat animation (sparkles and color flash) every frame
		g.updateEatAnimation()
		break
	case GameQuit:
	case GameOver:
		if !splashed {
			g.showGameOver()
			splashed = true
		}

		if !btnA.Get() {
			g.Start()
		}
		if !btnB.Get() {
			g.status = GameQuit
		}
	}
}
