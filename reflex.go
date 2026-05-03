package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	ReflexWait = iota
	ReflexReady
	ReflexGo
	ReflexResult
)

type ReflexGame struct {
	state       int
	targetColor int
	score       int
	round       int
	startTime   time.Time
	reactionMs  int64
	bestTime    int64
	splashShown bool
}

func NewReflexGame() *ReflexGame {
	return &ReflexGame{
		state:       ReflexWait,
		bestTime:    999999,
		splashShown: false,
	}
}

func (r *ReflexGame) Loop() {
	r.state = ReflexWait
	r.score = 0
	r.round = 0
	r.bestTime = 999999
	r.splashShown = false

	for {
		switch r.state {
		case ReflexWait:
			r.showInstructions()
			if !btnA.Get() {
				r.state = ReflexReady
				time.Sleep(300 * time.Millisecond)
			}
			if !btnB.Get() {
				return
			}

		case ReflexReady:
			r.showGetReady()
			// Random delay between 1-3 seconds
			delay := time.Duration(1000+rand.Int31n(2000)) * time.Millisecond
			time.Sleep(delay)
			r.state = ReflexGo

		case ReflexGo:
			r.round++
			r.targetColor = int(rand.Int31n(4)) // 0=Up, 1=Down, 2=Left, 3=Right
			r.startTime = time.Now()
			r.showTarget()

			// Wait for correct button press
			pressed := false
			for !pressed {
				if !btnB.Get() {
					return
				}
				switch r.targetColor {
				case 0:
					if !btnUp.Get() {
						pressed = true
					}
				case 1:
					if !btnDown.Get() {
						pressed = true
					}
				case 2:
					if !btnLeft.Get() {
						pressed = true
					}
				case 3:
					if !btnRight.Get() {
						pressed = true
					}
				}
				time.Sleep(10 * time.Millisecond)
			}

			r.reactionMs = time.Since(r.startTime).Milliseconds()
			if r.reactionMs < r.bestTime {
				r.bestTime = r.reactionMs
			}
			r.score += int(1000 - r.reactionMs)
			if r.score < 0 {
				r.score = 0
			}
			r.state = ReflexResult

		case ReflexResult:
			r.showResult()
			r.setNeoPixelsForResult()
			time.Sleep(2 * time.Second)
			r.clearNeoPixels()
			if r.round >= 5 {
				r.showFinalScore()
				time.Sleep(4 * time.Second)
				r.state = ReflexWait
				r.score = 0
				r.round = 0
				r.bestTime = 999999
				r.splashShown = false
			} else {
				r.state = ReflexReady
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *ReflexGame) showInstructions() {
	// Only draw the splash screen once
	if !r.splashShown {
		display.FillScreen(colors[BLACK])
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 40, 50, "REFLEX", colors[YELLOW])
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 40, 80, "TESTER", colors[YELLOW])
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 20, 130, "Press the matching", colors[WHITE])
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 20, 150, "direction button", colors[WHITE])
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 20, 170, "when arrow appears!", colors[WHITE])
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 60, 210, "[A] START", colors[SNAKE])
		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 60, 230, "[B] EXIT", colors[RED])
		r.splashShown = true
	}
}

func (r *ReflexGame) showGetReady() {
	// Clear the entire screen for GET READY message
	display.FillRectangle(0, 0, WIDTH, HEIGHT, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 30, 120, "GET READY!", colors[ORANGE])
}

func (r *ReflexGame) showTarget() {
	// Clear only the main content area, not the entire screen
	display.FillRectangle(0, 0, WIDTH, HEIGHT, colors[BLACK])

	// Show round number
	roundText := "Round " + strconv.Itoa(r.round) + "/5"
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 20, roundText, colors[TEXT])

	// Draw arrow based on direction
	centerX := int16(WIDTH / 2)
	centerY := int16(HEIGHT / 2)

	switch r.targetColor {
	case 0: // Up
		r.drawArrowUp(centerX, centerY, colors[PURPLE])
	case 1: // Down
		r.drawArrowDown(centerX, centerY, colors[SNAKE])
	case 2: // Left
		r.drawArrowLeft(centerX, centerY, colors[ORANGE])
	case 3: // Right
		r.drawArrowRight(centerX, centerY, colors[RED])
	}
}

func (r *ReflexGame) showResult() {
	// Clear only the center area where the circle was
	display.FillRectangle(0, 30, WIDTH, 180, colors[BLACK])

	// Show reaction time
	timeText := strconv.FormatInt(r.reactionMs, 10) + " ms"
	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, timeText)
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 80, timeText, colors[YELLOW])

	// Show feedback
	var feedback string
	var feedbackColor int
	if r.reactionMs < 250 {
		feedback = "AMAZING!"
		feedbackColor = SNAKE
	} else if r.reactionMs < 400 {
		feedback = "GREAT!"
		feedbackColor = SNAKE
	} else if r.reactionMs < 550 {
		feedback = "GOOD"
		feedbackColor = ORANGE
	} else if r.reactionMs < 750 {
		feedback = "OK"
		feedbackColor = ORANGE
	} else {
		feedback = "SLOW"
		feedbackColor = RED
	}

	w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, feedback)
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 120, feedback, colors[feedbackColor])

	// Show score
	scoreText := "Score: " + strconv.Itoa(r.score)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 80, 160, scoreText, colors[WHITE])
}

func (r *ReflexGame) showFinalScore() {
	// Clear the entire screen for final score (this is shown once at the end)
	display.FillRectangle(0, 0, WIDTH, HEIGHT, colors[BLACK])

	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 50, 50, "GAME OVER!", colors[YELLOW])

	scoreText := "Final Score: " + strconv.Itoa(r.score)
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 50, 100, scoreText, colors[WHITE])

	bestText := "Best Time: " + strconv.FormatInt(r.bestTime, 10) + " ms"
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 50, 130, bestText, colors[SNAKE])

	// Show rating
	avgTime := (5000 - int64(r.score)) / 5
	var rating string
	if avgTime < 300 {
		rating = "LIGHTNING FAST!"
	} else if avgTime < 450 {
		rating = "SUPER QUICK!"
	} else if avgTime < 600 {
		rating = "PRETTY GOOD!"
	} else {
		rating = "KEEP PRACTICING!"
	}

	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 30, 180, rating, colors[ORANGE])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 60, 220, "Press A to retry", colors[TEXT])
}

// Arrow drawing functions
func (r *ReflexGame) drawArrowUp(cx, cy int16, c color.RGBA) {
	// Draw arrow pointing up
	// Triangle head
	tinydraw.FilledTriangle(&display, cx, cy-50, cx-30, cy-10, cx+30, cy-10, c)
	// Shaft
	display.FillRectangle(cx-15, cy-10, 30, 60, c)
}

func (r *ReflexGame) drawArrowDown(cx, cy int16, c color.RGBA) {
	// Draw arrow pointing down
	// Shaft
	display.FillRectangle(cx-15, cy-50, 30, 60, c)
	// Triangle head
	tinydraw.FilledTriangle(&display, cx, cy+50, cx-30, cy+10, cx+30, cy+10, c)
}

func (r *ReflexGame) drawArrowLeft(cx, cy int16, c color.RGBA) {
	// Draw arrow pointing left
	// Triangle head
	tinydraw.FilledTriangle(&display, cx-50, cy, cx-10, cy-30, cx-10, cy+30, c)
	// Shaft
	display.FillRectangle(cx-10, cy-15, 60, 30, c)
}

func (r *ReflexGame) drawArrowRight(cx, cy int16, c color.RGBA) {
	// Draw arrow pointing right
	// Shaft
	display.FillRectangle(cx-50, cy-15, 60, 30, c)
	// Triangle head
	tinydraw.FilledTriangle(&display, cx+50, cy, cx+10, cy-30, cx+10, cy+30, c)
}

// NeoPixel feedback functions
func (r *ReflexGame) setNeoPixelsForResult() {
	ledColors := make([]color.RGBA, 2)
	var ledColor color.RGBA

	// Set color based on reaction time
	if r.reactionMs < 250 {
		// Amazing - bright green
		ledColor = color.RGBA{0, 255, 0, 255}
	} else if r.reactionMs < 400 {
		// Great - green
		ledColor = color.RGBA{0, 200, 0, 255}
	} else if r.reactionMs < 550 {
		// Good - yellow-green
		ledColor = color.RGBA{150, 200, 0, 255}
	} else if r.reactionMs < 750 {
		// OK - orange
		ledColor = colors[ORANGE]
	} else {
		// Slow - red
		ledColor = colors[RED]
	}

	ledColors[0] = ledColor
	ledColors[1] = ledColor
	leds.WriteColors(ledColors)
}

func (r *ReflexGame) clearNeoPixels() {
	ledColors := make([]color.RGBA, 2)
	ledColors[0] = color.RGBA{0, 0, 0, 255}
	ledColors[1] = color.RGBA{0, 0, 0, 255}
	leds.WriteColors(ledColors)
}
