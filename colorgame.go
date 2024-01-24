package main

import (
	"math/rand"
	"time"
)

const (
	CGWIDTH  = 20
	CGHEIGHT = 15
	CGSIZE   = 16
	CGS      = 9
)

var (
	colorgame [][]bool

	p1x, p1y, p2x, p2y, tx int16
	v1x, v1y, v2x, v2y, ty int16
)

func ColorGame() {
	display.FillScreen(colors[ORANGE])

	colorgame = make([][]bool, CGWIDTH)

	p1x = int16(rand.Int31n(100)) + 30
	p1y = int16(rand.Int31n(200)) + 20
	p2x = int16(rand.Int31n(100)) + 190
	p2y = int16(rand.Int31n(200)) + 20

	v1x = 2
	v1y = 3
	v2x = -3
	v2y = 2

	for i := int16(0); i < CGWIDTH; i++ {
		colorgame[i] = make([]bool, CGHEIGHT)
		if i >= CGWIDTH/2 {
			for j := int16(0); j < CGHEIGHT; j++ {
				colorgame[i][j] = true
				display.FillRectangle(i*CGSIZE, j*CGSIZE, CGSIZE, CGSIZE, colors[PURPLE])
			}
		}
	}

	display.FillRectangle(p1x-CGS, p1y-CGS, CGSIZE, CGSIZE, colors[PURPLE])
	display.FillRectangle(p2x-CGS, p2y-CGS, CGSIZE, CGSIZE, colors[ORANGE])

	for {
		display.FillRectangle(p1x-CGS, p1y-CGS, CGSIZE, CGSIZE, colors[ORANGE])
		display.FillRectangle(p2x-CGS, p2y-CGS, CGSIZE, CGSIZE, colors[PURPLE])

		p1x += v1x
		p1y += v1y
		p2x += v2x
		p2y += v2y

		if (p1x < CGS && v1x < 0) || (p1x > (320-CGS) && v1x > 0) {
			v1x = -v1x
		}
		if (p2x < CGS && v2x < 0) || (p2x > (320-CGS) && v2x > 0) {
			v2x = -v2x
		}
		if (p1y < CGS && v1y < 0) || (p1y > (240-CGS) && v1y > 0) {
			v1y = -v1y
		}
		if (p2y < CGS && v2y < 0) || (p2y > (240-CGS) && v2y > 0) {
			v2y = -v2y
		}

		tx = (p1x - CGS) / CGSIZE
		ty = p1y / CGSIZE
		if tx < 0 {
			tx = 0
		}
		if colorgame[tx][ty] {
			colorgame[tx][ty] = false
			if v1x < 0 {
				v1x = -v1x
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[ORANGE])
		}
		tx = (p1x + CGS) / CGSIZE
		if tx >= CGWIDTH {
			tx = CGWIDTH - 1
		}
		if colorgame[tx][ty] {
			colorgame[tx][ty] = false
			if v1x > 0 {
				v1x = -v1x
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[ORANGE])
		}
		tx = p1x / CGSIZE
		ty = (p1y - CGS) / CGSIZE
		if ty < 0 {
			ty = 0
		}
		if colorgame[tx][ty] {
			colorgame[tx][ty] = false
			if v1y < 0 {
				v1y = -v1y
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[ORANGE])
		}
		ty = (p1y + CGS) / CGSIZE
		if ty >= CGHEIGHT {
			ty = CGHEIGHT - 1
		}
		if colorgame[tx][ty] {
			colorgame[tx][ty] = false
			if v1y > 0 {
				v1y = -v1y
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[ORANGE])
		}

		// P2
		tx = (p2x - CGS) / CGSIZE
		ty = p2y / CGSIZE
		if tx < 0 {
			tx = 0
		}
		if !colorgame[tx][ty] {
			colorgame[tx][ty] = true
			if v2x < 0 {
				v2x = -v2x
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[PURPLE])
		}
		tx = (p2x + CGS) / CGSIZE
		if tx >= CGWIDTH {
			tx = CGWIDTH - 2
		}
		if !colorgame[tx][ty] {
			colorgame[tx][ty] = true
			if v2x > 0 {
				v2x = -v2x
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[PURPLE])
		}
		tx = p2x / CGSIZE
		ty = (p2y - CGS) / CGSIZE
		if ty < 0 {
			ty = 0
		}
		if !colorgame[tx][ty] {
			colorgame[tx][ty] = true
			if v2y < 0 {
				v2y = -v2y
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[PURPLE])
		}
		ty = (p2y + CGS) / CGSIZE
		if ty >= CGHEIGHT {
			ty = CGHEIGHT - 2
		}
		if !colorgame[tx][ty] {
			colorgame[tx][ty] = true
			if v2y > 0 {
				v2y = -v2y
			}
			display.FillRectangle(tx*CGSIZE, ty*CGSIZE, CGSIZE, CGSIZE, colors[PURPLE])
		}

		display.FillRectangle(p1x-CGS, p1y-CGS, CGSIZE, CGSIZE, colors[PURPLE])
		display.FillRectangle(p2x-CGS, p2y-CGS, CGSIZE, CGSIZE, colors[ORANGE])

		if !btnB.Get() {
			break
		}

		time.Sleep(10 * time.Millisecond)

	}

}
