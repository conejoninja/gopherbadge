package main

import (
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

func schedule(day int, hour int) {
	quit := false
	for {
		showDay(day, hour)
		for {
			time.Sleep(200 * time.Millisecond)
			if !btnDown.Get() {
				hour++
				if hour > len(scheduleData[day].talks)-4 {
					hour = len(scheduleData[day].talks) - 4
				} else {
					break
				}
			}
			if !btnUp.Get() {
				hour--
				if hour < 0 {
					hour = 0
				} else {
					break
				}
			}
			if !btnLeft.Get() {
				day--
				if day < 0 {
					day = 0
				} else {
					hour = 0
					break
				}
			}
			if !btnRight.Get() {
				day++
				if day > len(scheduleData)-1 {
					day = len(scheduleData) - 1
				} else {
					hour = 0
					break
				}
			}
			if !btnB.Get() {
				quit = true
				break
			}
		}
		if quit {
			break
		}
	}
}

func showDay(day int, hour int) {
	display.FillScreen(colors[WHITE])

	tinydraw.Line(&display, 0, 76, WIDTH, 76, colors[BLACK])
	tinydraw.Line(&display, 0, 126, WIDTH, 126, colors[BLACK])
	tinydraw.Line(&display, 0, 176, WIDTH, 176, colors[BLACK])
	tinydraw.Line(&display, 50, 0, 50, HEIGHT, colors[BLACK])

	display.FillRectangle(0, 0, WIDTH, 26, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 8, 18, scheduleData[day].title, colors[WHITE])

	display.FillRectangle(0, HEIGHT-14, WIDTH, 14, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 0, HEIGHT-4, " [arrows] TO NAVIGATE                     [B] TO EXIT", colors[WHITE])

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 46, scheduleData[day].talks[hour].startHour, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 69, scheduleData[day].talks[hour].endHour, colors[BLACK])

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 40, scheduleData[day].talks[hour].line1, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 53, scheduleData[day].talks[hour].line2, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 66, scheduleData[day].talks[hour].line3, colors[BLACK])

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 96, scheduleData[day].talks[hour+1].startHour, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 119, scheduleData[day].talks[hour+1].endHour, colors[BLACK])

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 90, scheduleData[day].talks[hour+1].line1, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 103, scheduleData[day].talks[hour+1].line2, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 116, scheduleData[day].talks[hour+1].line3, colors[BLACK])

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 146, scheduleData[day].talks[hour+2].startHour, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 169, scheduleData[day].talks[hour+2].endHour, colors[BLACK])

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 140, scheduleData[day].talks[hour+2].line1, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 153, scheduleData[day].talks[hour+2].line2, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 166, scheduleData[day].talks[hour+2].line3, colors[BLACK])

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 196, scheduleData[day].talks[hour+3].startHour, colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 219, scheduleData[day].talks[hour+3].endHour, colors[BLACK])

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 190, scheduleData[day].talks[hour+3].line1, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 203, scheduleData[day].talks[hour+3].line2, colors[BLACK])
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 216, scheduleData[day].talks[hour+3].line3, colors[BLACK])

	display.Display()
}
