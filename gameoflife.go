package main

import (
	"image/color"
	"time"

	"github.com/acifani/vita/lib/game"
)

var (
	gamebuffer []byte

	universe *game.Universe

	width      uint32 = 53
	height     uint32 = 40
	population        = 20
	cellSize   int16  = 6

	wh = colors[WHITE]

	cellBuf = []color.RGBA{
		wh, wh, wh, wh, wh, wh,
		wh, wh, bk, bk, wh, wh,
		wh, bk, bk, bk, bk, wh,
		wh, bk, bk, bk, bk, wh,
		wh, wh, bk, bk, wh, wh,
		wh, wh, wh, wh, wh, wh,
	}
)

func GameOfLife() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	gamebuffer = make([]byte, height*width)
	universe = game.NewUniverse(height, width)
	universe.Randomize(population)
	universe.Read(gamebuffer)

	x, y, z := accel.ReadRawAcceleration()

	for {
		drawGrid()
		display.Display()
		universe.Read(gamebuffer)

		universe.Tick()

		x, y, z = accel.ReadRawAcceleration()
		if x < (-31000) || x > 31000 || y < (-31000) || y > 31000 || z < (-31000) || z > 31000 {
			universe.Reset()
			universe.Randomize(population)
		}
		if !btnB.Get() {
			break
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func drawGrid() {
	var rows, cols uint32

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := universe.GetIndex(rows, cols)

			switch {
			case universe.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case universe.Cell(idx) == game.Alive:
				display.FillRectangleWithBuffer(1+cellSize*int16(cols), cellSize*int16(rows), cellSize, cellSize, cellBuf)
			default: // game.Dead
				display.FillRectangle(1+cellSize*int16(cols), cellSize*int16(rows), cellSize, cellSize, colors[WHITE])
			}

		}
	}
}
