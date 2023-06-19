package main

import (
	"image/color"
	"machine"
	"strings"
	"time"

	"tinygo.org/x/tinyfont/proggy"

	"machine/usb/hid/keyboard"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyterm"
)

var (
	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	btnA     = machine.BUTTON_A
	btnB     = machine.BUTTON_B
	btnUp    = machine.BUTTON_UP
	btnLeft  = machine.BUTTON_LEFT
	btnDown  = machine.BUTTON_DOWN
	btnRight = machine.BUTTON_RIGHT

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	terminal = tinyterm.NewTerminal(&display)

	font = &proggy.TinySZ8pt7b
)

var (
	shifted  bool
	lastKey  string
	lastTime time.Time
)

var logo = `
 ______   __        __    ______    __  __    ______ 
/\  ___\ /\ \      /\ \  /\  ___\  /\ \_\ \  /\__  _\
\ \  __\ \ \ \____ \ \ \ \ \ \__ \ \ \  __ \ \/_/\ \/
 \ \_\    \ \_____\ \ \_\ \ \_____\ \ \_\ \_\   \ \_\
  \/_/     \/_____/  \/_/  \/_____/  \/_/\/_/    \/_/
 ______     ______     _____     ______     ______
/\  == \   /\  __ \   /\  __-.  /\  ___\   /\  ___\
\ \  __<   \ \  __ \  \ \ \/\ \ \ \ \__ \  \ \  __\
 \ \_____\  \ \_\ \_\  \ \____-  \ \_____\  \ \_____\
  \/_____/   \/_/\/_/   \/____/   \/_____/   \/_____/
`

func main() {
	go handleDisplay()

	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		// hold down button B to shift to access second set of arrow commands
		if !btnB.Get() {
			shifted = true
		} else {
			shifted = false
		}

		// takeoff / land
		if !btnA.Get() {
			handleShiftedKey("[", "]")
		}

		// front flip
		// handleKey("t")

		if !btnLeft.Get() {
			handleShiftedKey("j", "a")
		}

		if !btnUp.Get() {
			handleShiftedKey("i", "w")
		}

		if !btnDown.Get() {
			handleShiftedKey("k", "s")
		}

		if !btnRight.Get() {
			handleShiftedKey("l", "d")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func handleShiftedKey(key1, key2 string) {
	if shifted {
		handleKey(key1)
		return
	}
	handleKey(key2)
}

func handleKey(key string) {
	// simple debounce
	if key == lastKey && time.Since(lastTime) < 150*time.Millisecond {
		return
	}

	kb := keyboard.New()
	kb.Write([]byte(key))

	lastKey, lastTime = key, time.Now()
}

func handleDisplay() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
		Width:    240,
	})

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})

	display.FillScreen(black)

	showSplash()

	input := make([]byte, 64)
	i := 0

	for {
		if machine.Serial.Buffered() > 0 {
			data, _ := machine.Serial.ReadByte()

			switch data {
			case 13:
				// return key
				terminal.Write([]byte("\r\n"))
				terminal.Write(input[:i])
				i = 0
			default:
				input[i] = data
				i++
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func showSplash() {
	for _, line := range strings.Split(strings.TrimSuffix(logo, "\n"), "\n") {
		terminal.Write([]byte(line + "\n"))
	}
}
