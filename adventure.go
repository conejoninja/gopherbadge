package main

import (
	"strconv"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"

	qrcode "github.com/skip2/go-qrcode"
)

func adventure() {
	quit := false
	s := 0
	score := 0
	money := 0
	costume := 0
	pancake := false
	talk := false
	opinion := 0
	selected := int16(0)
	released := true
	for {
		println(s)
		scene(s)
		selected = 0
		score += s
		if s == 5 {
			costume = 1
		} else if s == 3 {
			talk = true
		} else if s == 6 {
			pancake = true
		} else if s == 8 {
			money += 689
			costume = 2
		} else if s == 17 {
			money += 5
		} else if s == 20 {
			opinion = 57
		} else if s == 21 {
			opinion = 37
		} else if s == 22 {
			opinion = 21
		} else if s == 25 {
			myNameIs(YourName)
		} else if s == 27 {
			if costume == 0 {
				s = 30
			} else if costume == 1 {
				s = 28
			} else {
				s = 29
			}
			continue
		} else if s == 32 {
			s = 34
			if talk {
				s = 33
			}
			continue
		} else if s == 41 {
			showGameStats(score, money, costume, opinion, pancake, talk)
		} else if s == 42 {
			showGameQR()
		} else if s == 44 {
			return
		}

		tinydraw.FilledCircle(&display, 10, HEIGHT-28+10*selected, 3, colors[WHITE])
		released = true
		for {

			if released && !btnUp.Get() && selected > 0 {
				tinydraw.FilledCircle(&display, 10, HEIGHT-28+10*selected, 3, colors[BLACK])
				selected--
				tinydraw.FilledCircle(&display, 10, HEIGHT-28+10*selected, 3, colors[WHITE])
			}

			if released && !btnDown.Get() && selected < 2 {
				tinydraw.FilledCircle(&display, 10, HEIGHT-28+10*selected, 3, colors[BLACK])
				selected++
				tinydraw.FilledCircle(&display, 10, HEIGHT-28+10*selected, 3, colors[WHITE])
			}

			if released && !btnA.Get() {
				if selected == 0 {
					s = sceneData[s].sceneA
					if s == 1 {
						money += 385
					} else if s == 2 {
						money += 335
					}
					break
				} else if selected == 1 {
					s = sceneData[s].sceneB
					if s == 1 {
						money += 600
					} else if s == 2 {
						money += 550
					}
					break
				} else {
					s = sceneData[s].sceneC
					if s == 1 {
						money += 1050
					} else if s == 2 {
						money += 970
					}
					break
				}
			}

			if btnA.Get() && btnUp.Get() && btnDown.Get() {
				released = true
			} else {
				released = false
			}

			if !btnB.Get() {
				return
			}
			time.Sleep(200 * time.Millisecond)

		}
		if quit {
			break
		}
	}
}

func scene(s int) {
	display.FillScreen(colors[WHITE])

	ss := splitBefore(sceneData[s].description)
	for i := int16(0); i < int16(len(ss)); i++ {
		tinyfont.WriteLine(&display, &freemono.Regular9pt7b, 6, 6+i*16, ss[i], colors[BLACK])
	}

	display.FillRectangle(0, HEIGHT-33, WIDTH, 33, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 20, HEIGHT-24, sceneData[s].optionA, colors[WHITE])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 20, HEIGHT-14, sceneData[s].optionB, colors[WHITE])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 20, HEIGHT-4, sceneData[s].optionC, colors[WHITE])

	tinydraw.Circle(&display, 10, HEIGHT-28, 4, colors[WHITE])
	tinydraw.Circle(&display, 10, HEIGHT-18, 4, colors[WHITE])
	tinydraw.Circle(&display, 10, HEIGHT-8, 4, colors[WHITE])

	display.Display()
}

func splitBefore(str string) []string {
	l := len(str)
	a := 0
	s := make([]string, 1)
	for {
		if l <= 28 {
			s = append(s, str[a:a+l])
			break
		} else {
			for i := 28; i > 0; i-- {
				if string(str[a+i]) == " " {
					s = append(s, str[a:a+i])
					a = a + i + 1
					l = l - i - 1
					break
				}
			}
		}
	}
	return s
}

func showGameStats(score, money, costume, opinion int, pancake, talk bool) {

	c := "SCORE: " + strconv.Itoa(score)
	w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, c)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 26, c, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 50, "Money spent: "+strconv.Itoa(money)+" EUR", colors[BLACK])

	c = "YOURSELF"
	if costume == 1 {
		c = "UNICORN PIJAMA"
	} else if costume == 2 {
		c = "DELUXE TUXEDO"
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 60, "You were dressed as: "+c, colors[BLACK])

	c = "NO pancakes"
	if pancake {
		c = "37 pancakes"
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 70, "Your breakfast: "+c, colors[BLACK])

	if talk {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 80, "You DID give a talk (and it was awesome)", colors[BLACK])
	} else {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 80, "You did NOT give a talk (sad noises)", colors[BLACK])
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 90, "Your opinion is "+strconv.Itoa(opinion)+"% popular", colors[BLACK])

	display.Display()
}

func showGameQR() {

	// https://tinyurl.com/gceuADV
	qr, err := qrcode.New("https://tinyurl.com/gceuADV", qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor)
	by := (HEIGHT - size*factor) / 2
	display.FillScreen(colors[WHITE])
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[0])
			} else {
				display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, colors[1])
			}
		}
	}

	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 30, "SCAN &", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 50, "SHARE", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 70, "YOUR", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 90, "SCORE", colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 110, "ONLINE", colors[BLACK])

	display.Display()
}
