package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

// Maze cell types
const (
	pmEmpty = 0
	pmWall  = 1
	pmDot   = 2
	pmPower = 3
)

// Directions
const (
	pmUp    = 0
	pmDown  = 1
	pmLeft  = 2
	pmRight = 3
)

const (
	pmW          = 32
	pmH          = 24
	pmCell       = 10
	pmGhostCount = 4
	pmFrightLen  = 40 // frames of fright mode (~3.2s at 80ms/frame)

	pmPointsDot   = 10
	pmPointsPower = 50
	pmPointsGhost = 200
)

var (
	pmColorWall   = color.RGBA{33, 33, 222, 255}
	pmColorDot    = color.RGBA{255, 255, 255, 255}
	pmColorPower  = color.RGBA{255, 255, 200, 255}
	pmColorPac    = color.RGBA{255, 255, 0, 255}
	pmColorFright = color.RGBA{33, 33, 200, 255}
	pmColorBg     = color.RGBA{0, 0, 0, 255}
	pmGhostColor  = [4]color.RGBA{
		{255, 0, 0, 255},     // red (Blinky)
		{255, 184, 255, 255}, // pink (Pinky)
		{0, 255, 255, 255},   // cyan (Inky)
		{255, 184, 82, 255},  // orange (Clyde)
	}
)

type PacGhost struct {
	x      int16
	y      int16
	dir    int16
	fright int
	startX int16
	startY int16
}

const (
	pmDiffEasy   = 0
	pmDiffNormal = 1
	pmDiffHard   = 2
)

var pmHighScore int  // persists across games within a power cycle
var pmDifficulty int // pmDiffEasy, pmDiffNormal, or pmDiffHard

var pmPowerPositions = [4][2]int16{
	{1, 3}, {30, 3}, {1, 20}, {30, 20},
}

type PacmanGame struct {
	maze              [pmH][pmW]uint8
	pacX              int16
	pacY              int16
	dir               int16
	nextDir           int16
	ghosts            [pmGhostCount]PacGhost
	score             int
	dots              int
	totalDots         int
	powerRespawnCount int
	status            uint8
	frame             int
	won               bool
}

func NewPacmanGame() *PacmanGame {
	return &PacmanGame{
		status: GameSplash,
	}
}

func (g *PacmanGame) Loop() {
	g.status = GameSplash
	g.showSplash()
	if g.status == GameQuit {
		return
	}
	g.startGame()
	drawn := false
	for {
		switch g.status {
		case GamePlay:
			g.update()
		case GameOver:
			if !drawn {
				g.showGameOver()
				drawn = true
			}
			if !btnA.Get() {
				// Restart: show splash again for difficulty selection
				g.status = GameSplash
				g.showSplash()
				if g.status == GameQuit {
					return
				}
				g.startGame()
				drawn = false
			}
			if !btnB.Get() {
				return
			}
		case GameQuit:
			return
		}
		time.Sleep(80 * time.Millisecond)
	}
}

func (g *PacmanGame) showSplash() {
	display.FillScreen(pmColorBg)
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 55, 50, "PAC-MAN", pmColorPac)

	if pmHighScore > 0 {
		hsText := "Best: " + strconv.Itoa(pmHighScore)
		w32, _ := tinyfont.LineWidth(&freesans.Regular12pt7b, hsText)
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, (WIDTH-int16(w32))/2, 85, hsText, colors[YELLOW])
	}

	// Difficulty menu
	selected := pmDifficulty
	g.drawDifficultyMenu(selected)

	btnUpPressed := false
	btnDownPressed := false
	btnAPressed := true // avoid immediate selection

	for {
		if !btnUp.Get() {
			if !btnUpPressed && selected > 0 {
				selected--
				g.drawDifficultyMenu(selected)
			}
			btnUpPressed = true
		} else {
			btnUpPressed = false
		}

		if !btnDown.Get() {
			if !btnDownPressed && selected < pmDiffHard {
				selected++
				g.drawDifficultyMenu(selected)
			}
			btnDownPressed = true
		} else {
			btnDownPressed = false
		}

		if !btnA.Get() {
			if !btnAPressed {
				pmDifficulty = selected
				return
			}
		} else {
			btnAPressed = false
		}

		if !btnB.Get() {
			g.status = GameQuit
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (g *PacmanGame) drawDifficultyMenu(selected int) {
	labels := [3]string{"Easy Mode", "Normal Mode", "Hard Mode"}
	for i := 0; i < 3; i++ {
		c := colors[WHITE]
		if i == selected {
			c = colors[YELLOW]
		}
		tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, int16(125+i*28), labels[i], c)
	}
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 30, 215, "Press A to start", colors[SNAKE])
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 30, 238, "[B] EXIT", colors[RED])
}

func (g *PacmanGame) showGameOver() {
	// Update high score
	if g.score > pmHighScore {
		pmHighScore = g.score
	}

	display.FillScreen(pmColorBg)
	if g.won {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 40, 50, "YOU WIN!", colors[SNAKE])
	} else {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 20, 50, "GAME OVER", colors[RED])
	}
	scoreText := "Score: " + strconv.Itoa(g.score)
	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, scoreText)
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 100, scoreText, colors[YELLOW])

	hsText := "Best: " + strconv.Itoa(pmHighScore)
	w32, _ = tinyfont.LineWidth(&freesans.Regular12pt7b, hsText)
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, (WIDTH-int16(w32))/2, 135, hsText, colors[WHITE])

	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 190, "Press A to retry", colors[WHITE])
	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 50, 220, "Press B to exit", colors[TEXT])
}

func (g *PacmanGame) startGame() {
	g.score = 0
	g.dots = 0
	g.won = false
	g.frame = 0
	g.dir = pmLeft
	g.nextDir = pmLeft
	g.powerRespawnCount = 0
	g.initMaze()
	g.totalDots = g.dots

	// Pac-man start position: center-bottom area
	g.pacX = 15
	g.pacY = 17
	if g.maze[g.pacY][g.pacX] == pmDot {
		g.dots--
	}
	g.maze[g.pacY][g.pacX] = pmEmpty

	// Ghost start positions: around ghost house
	gx := [4]int16{14, 17, 14, 17}
	gy := [4]int16{9, 9, 13, 13}
	gd := [4]int16{pmLeft, pmRight, pmLeft, pmRight}
	for i := 0; i < pmGhostCount; i++ {
		g.ghosts[i] = PacGhost{
			x: gx[i], y: gy[i],
			dir:    gd[i],
			startX: gx[i], startY: gy[i],
		}
		// Clear dot under ghost start
		if g.maze[gy[i]][gx[i]] == pmDot {
			g.dots--
			g.maze[gy[i]][gx[i]] = pmEmpty
		}
	}

	// Draw the full maze once at start
	g.drawFullMaze()
	g.drawPacman()
	for i := 0; i < pmGhostCount; i++ {
		g.drawGhost(i)
	}
	g.status = GamePlay
}

func (g *PacmanGame) initMaze() {
	g.dots = 0

	// Fill with dots
	for y := 0; y < pmH; y++ {
		for x := 0; x < pmW; x++ {
			g.maze[y][x] = pmDot
			g.dots++
		}
	}

	// Helpers
	wall := func(x, y, w, h int) {
		for dy := 0; dy < h; dy++ {
			for dx := 0; dx < w; dx++ {
				if g.maze[y+dy][x+dx] == pmDot {
					g.dots--
				}
				g.maze[y+dy][x+dx] = pmWall
			}
		}
	}
	empty := func(x, y, w, h int) {
		for dy := 0; dy < h; dy++ {
			for dx := 0; dx < w; dx++ {
				if g.maze[y+dy][x+dx] == pmDot {
					g.dots--
				}
				g.maze[y+dy][x+dx] = pmEmpty
			}
		}
	}

	// Border
	wall(0, 0, pmW, 1)
	wall(0, pmH-1, pmW, 1)
	wall(0, 0, 1, pmH)
	wall(pmW-1, 0, 1, pmH)

	// Top section
	// Row 2-3 blocks
	wall(2, 2, 4, 2)
	wall(7, 2, 6, 2)
	wall(14, 2, 4, 2)
	wall(19, 2, 6, 2)
	wall(26, 2, 4, 2)

	// Row 5 blocks
	wall(2, 5, 4, 1)
	wall(7, 5, 2, 3)
	wall(10, 5, 5, 1)
	wall(17, 5, 5, 1)
	wall(23, 5, 2, 3)
	wall(26, 5, 4, 1)

	// Row 7 blocks
	wall(2, 7, 4, 1)
	wall(10, 7, 5, 1)
	wall(17, 7, 5, 1)
	wall(26, 7, 4, 1)

	// Row 8-9 blocks
	wall(2, 8, 4, 2)
	wall(7, 8, 2, 2)
	wall(23, 8, 2, 2)
	wall(26, 8, 4, 2)

	// Ghost house: rows 10-12, cols 12-19
	wall(12, 10, 8, 1)  // top
	wall(12, 12, 8, 1)  // bottom
	wall(12, 10, 1, 3)  // left
	wall(19, 10, 1, 3)  // right
	empty(13, 11, 6, 1) // interior
	empty(15, 10, 2, 1) // entrance

	// Bottom section (mirrors top)
	// Row 14-15 blocks
	wall(2, 14, 4, 2)
	wall(7, 14, 2, 2)
	wall(23, 14, 2, 2)
	wall(26, 14, 4, 2)

	// Row 16 blocks
	wall(2, 16, 4, 1)
	wall(10, 16, 5, 1)
	wall(17, 16, 5, 1)
	wall(26, 16, 4, 1)

	// Row 17 blocks
	wall(7, 17, 2, 1)
	wall(23, 17, 2, 1)

	// Row 18 blocks
	wall(2, 18, 4, 1)
	wall(7, 18, 2, 3)
	wall(10, 18, 5, 1)
	wall(17, 18, 5, 1)
	wall(23, 18, 2, 3)
	wall(26, 18, 4, 1)

	// Row 20-21 blocks
	wall(2, 20, 4, 2)
	wall(7, 20, 6, 2)
	wall(14, 20, 4, 2)
	wall(19, 20, 6, 2)
	wall(26, 20, 4, 2)

	// Power pellets in corners
	for _, p := range pmPowerPositions {
		g.maze[p[1]][p[0]] = pmPower
	}
}

func (g *PacmanGame) spawnPowerPellet() {
	// Collect indices of power pellet positions that have been eaten
	var eaten [4]int
	n := 0
	for i, p := range pmPowerPositions {
		if g.maze[p[1]][p[0]] == pmEmpty {
			eaten[n] = i
			n++
		}
	}
	if n == 0 {
		return
	}
	idx := eaten[rand.Int31n(int32(n))]
	p := pmPowerPositions[idx]
	g.maze[p[1]][p[0]] = pmPower
	g.dots++
	g.totalDots++
	g.drawCell(p[0], p[1])
}

func (g *PacmanGame) drawFullMaze() {
	display.FillScreen(pmColorBg)
	for y := int16(0); y < pmH; y++ {
		for x := int16(0); x < pmW; x++ {
			g.drawCell(x, y)
		}
	}
}

func (g *PacmanGame) drawCell(x, y int16) {
	px := x * pmCell
	py := y * pmCell
	switch g.maze[y][x] {
	case pmWall:
		display.FillRectangle(px, py, pmCell, pmCell, pmColorWall)
	case pmDot:
		display.FillRectangle(px, py, pmCell, pmCell, pmColorBg)
		display.FillRectangle(px+4, py+4, 2, 2, pmColorDot)
	case pmPower:
		display.FillRectangle(px, py, pmCell, pmCell, pmColorBg)
		display.FillRectangle(px+2, py+2, 6, 6, pmColorPower)
	default: // pmEmpty
		display.FillRectangle(px, py, pmCell, pmCell, pmColorBg)
	}
}

func (g *PacmanGame) drawPacman() {
	px := g.pacX * pmCell
	py := g.pacY * pmCell
	display.FillRectangle(px+1, py+1, pmCell-2, pmCell-2, pmColorPac)
	// Mouth: small gap in the direction of movement
	if g.frame%4 < 2 {
		switch g.dir {
		case pmRight:
			display.FillRectangle(px+7, py+4, 3, 2, pmColorBg)
		case pmLeft:
			display.FillRectangle(px, py+4, 3, 2, pmColorBg)
		case pmUp:
			display.FillRectangle(px+4, py, 2, 3, pmColorBg)
		case pmDown:
			display.FillRectangle(px+4, py+7, 2, 3, pmColorBg)
		}
	}
}

func (g *PacmanGame) drawGhost(idx int) {
	ghost := &g.ghosts[idx]
	px := ghost.x * pmCell
	py := ghost.y * pmCell
	c := pmGhostColor[idx]
	if ghost.fright > 0 {
		c = pmColorFright
	}
	display.FillRectangle(px+1, py+1, pmCell-2, pmCell-2, c)
	// Eyes: two small white dots
	display.FillRectangle(px+2, py+3, 2, 2, pmColorDot)
	display.FillRectangle(px+6, py+3, 2, 2, pmColorDot)
}

func (g *PacmanGame) update() {
	// Input
	switch {
	case !btnB.Get():
		g.status = GameQuit
		return
	case !btnRight.Get():
		g.nextDir = pmRight
	case !btnLeft.Get():
		g.nextDir = pmLeft
	case !btnUp.Get():
		g.nextDir = pmUp
	case !btnDown.Get():
		g.nextDir = pmDown
	}

	g.frame++

	// Try desired direction first
	if g.canMove(g.pacX, g.pacY, g.nextDir) {
		g.dir = g.nextDir
	}
	if g.canMove(g.pacX, g.pacY, g.dir) {
		oldX, oldY := g.pacX, g.pacY
		g.pacX += pmDX(g.dir)
		g.pacY += pmDY(g.dir)

		// Eat dot or power pellet
		cell := g.maze[g.pacY][g.pacX]
		if cell == pmDot {
			g.maze[g.pacY][g.pacX] = pmEmpty
			points := pmPointsDot
			// Dots are worth twice more in hard mode
			if pmDifficulty == pmDiffHard {
				points *= 2
			}
			g.score += points
			g.dots--
		} else if cell == pmPower {
			g.maze[g.pacY][g.pacX] = pmEmpty
			g.score += pmPointsPower
			g.dots--
			for i := range g.ghosts {
				g.ghosts[i].fright = pmFrightLen
			}
		}

		// Bonus power pellets (easy/normal only): one at 1/2, one at 3/4
		if g.powerRespawnCount < 2 && pmDifficulty != pmDiffHard {
			eaten := g.totalDots - g.dots
			threshold := g.totalDots / 2
			if g.powerRespawnCount == 1 {
				threshold = g.totalDots * 3 / 4
			}
			if eaten >= threshold {
				g.spawnPowerPellet()
				g.powerRespawnCount++
			}
		}

		// Redraw old cell
		g.drawCell(oldX, oldY)

		// Check win
		if g.dots <= 0 {
			g.won = true
			g.status = GameOver
			return
		}

		// Check collision with ghosts after pac-man moves
		for i := 0; i < pmGhostCount; i++ {
			if g.ghosts[i].x == g.pacX && g.ghosts[i].y == g.pacY {
				if g.ghosts[i].fright > 0 {
					g.eatGhost(i)
				} else {
					g.status = GameOver
					return
				}
			}
		}
	}

	// Draw pac-man at current position (for mouth animation)
	g.drawPacman()

	// Move ghosts (speed depends on difficulty)
	// Easy: every 3rd frame (~33%), Normal: every 2nd (~50%), Hard: 3 of 4 (~75%)
	moveGhosts := false
	switch pmDifficulty {
	case pmDiffEasy:
		moveGhosts = g.frame%3 == 0
	case pmDiffNormal:
		moveGhosts = g.frame%2 == 0
	case pmDiffHard:
		moveGhosts = g.frame%4 != 0
	}
	if moveGhosts {
		for i := 0; i < pmGhostCount; i++ {
			oldX, oldY := g.ghosts[i].x, g.ghosts[i].y
			g.moveGhost(i)
			// Redraw old cell
			g.drawCell(oldX, oldY)
			// Draw ghost at new position
			g.drawGhost(i)

			// Decrement fright
			if g.ghosts[i].fright > 0 {
				g.ghosts[i].fright--
			}

			// Check collision
			if g.ghosts[i].x == g.pacX && g.ghosts[i].y == g.pacY {
				if g.ghosts[i].fright > 0 {
					g.eatGhost(i)
				} else {
					g.status = GameOver
					return
				}
			}
		}
	}
}

func (g *PacmanGame) eatGhost(idx int) {
	g.score += pmPointsGhost
	g.ghosts[idx].x = g.ghosts[idx].startX
	g.ghosts[idx].y = g.ghosts[idx].startY
	g.ghosts[idx].fright = 0
	g.drawGhost(idx)
}

func (g *PacmanGame) moveGhost(idx int) {
	ghost := &g.ghosts[idx]

	// Gather valid moves (no reversals)
	reverse := pmReverse(ghost.dir)
	type opt struct {
		dir  int16
		x, y int16
		dist int16
	}
	var opts [4]opt
	n := 0

	for d := int16(0); d < 4; d++ {
		if d == reverse {
			continue
		}
		nx := ghost.x + pmDX(d)
		ny := ghost.y + pmDY(d)
		if nx < 0 || nx >= pmW || ny < 0 || ny >= pmH {
			continue
		}
		if g.maze[ny][nx] == pmWall {
			continue
		}
		dist := abs16(nx-g.pacX) + abs16(ny-g.pacY)
		opts[n] = opt{d, nx, ny, dist}
		n++
	}

	if n == 0 {
		// Dead end: reverse
		d := reverse
		nx := ghost.x + pmDX(d)
		ny := ghost.y + pmDY(d)
		if nx >= 0 && nx < pmW && ny >= 0 && ny < pmH && g.maze[ny][nx] != pmWall {
			ghost.x = nx
			ghost.y = ny
			ghost.dir = d
		}
		return
	}

	if ghost.fright > 0 {
		// Frightened: random
		c := rand.Int31n(int32(n))
		ghost.x = opts[c].x
		ghost.y = opts[c].y
		ghost.dir = opts[c].dir
	} else {
		// Chase: pick closest to pac-man, with 25% random chance
		best := 0
		for i := 1; i < n; i++ {
			if opts[i].dist < opts[best].dist {
				best = i
			}
		}
		if n > 1 && rand.Int31n(4) == 0 {
			c := rand.Int31n(int32(n))
			best = int(c)
		}
		ghost.x = opts[best].x
		ghost.y = opts[best].y
		ghost.dir = opts[best].dir
	}
}

func (g *PacmanGame) canMove(x, y, dir int16) bool {
	nx := x + pmDX(dir)
	ny := y + pmDY(dir)
	if nx < 0 || nx >= pmW || ny < 0 || ny >= pmH {
		return false
	}
	return g.maze[ny][nx] != pmWall
}

func pmDX(dir int16) int16 {
	switch dir {
	case pmLeft:
		return -1
	case pmRight:
		return 1
	}
	return 0
}

func pmDY(dir int16) int16 {
	switch dir {
	case pmUp:
		return -1
	case pmDown:
		return 1
	}
	return 0
}

func pmReverse(dir int16) int16 {
	switch dir {
	case pmUp:
		return pmDown
	case pmDown:
		return pmUp
	case pmLeft:
		return pmRight
	case pmRight:
		return pmLeft
	}
	return -1
}

func abs16(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}
