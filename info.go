package main

import (
	"machine"
	"runtime"
	"strconv"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func Info() {
	// Scan I2C once at startup (slow operation)
	i2cDevices := scanI2C()

	// Draw static elements once
	display.FillScreen(colors[BLACK])
	lineHeight := int16(22)
	valueX := int16(105) // X position where dynamic values start

	// Title
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 25, "System Status", colors[YELLOW])

	// Draw static labels once
	y := int16(50)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, y, "CPU Temp:", colors[WHITE])
	y += lineHeight
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, y, "Heap:", colors[WHITE])
	y += lineHeight

	// I2C device info (static, scanned once at startup)
	i2cStr := "I2C: " + strconv.Itoa(len(i2cDevices)) + " device(s)"
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, y, i2cStr, colors[WHITE])
	y += lineHeight

	// List detected I2C addresses
	if len(i2cDevices) > 0 {
		addrStr := "  Addr: "
		for i, addr := range i2cDevices {
			if i > 0 {
				addrStr += ", "
			}
			addrStr += "0x" + strconv.FormatInt(int64(addr), 16)
			if len(addrStr) > 35 { // Truncate if too long
				addrStr += "..."
				break
			}
		}
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, y, addrStr, colors[TEXT])
	}

	// Instructions
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 230, "Press any button to exit", colors[TEXT])

	for {
		y = int16(50)

		// CPU Temperature - clear only value area and redraw
		display.FillRectangle(valueX, y-15, 215, 20, colors[BLACK])
		tempMilliC := machine.ReadTemperature()
		tempC := tempMilliC / 1000
		tempFrac := (tempMilliC % 1000) / 100
		if tempFrac < 0 {
			tempFrac = -tempFrac
		}
		tempStr := strconv.Itoa(int(tempC)) + "." + strconv.Itoa(int(tempFrac)) + " C"
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, valueX, y, tempStr, colors[WHITE])
		y += lineHeight

		// Memory stats - clear only value area and redraw
		display.FillRectangle(valueX-40, y-15, 255, 20, colors[BLACK])
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		heapStr := formatBytes(memStats.HeapAlloc) + " / " + formatBytes(memStats.Sys)
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, valueX-40, y, heapStr, colors[WHITE])

		time.Sleep(500 * time.Millisecond)
		if !btnA.Get() || !btnB.Get() || !btnUp.Get() || !btnLeft.Get() || !btnRight.Get() || !btnDown.Get() {
			break
		}
	}
}

// scanI2C scans the I2C bus for connected devices
func scanI2C() []uint8 {
	var devices []uint8
	buf := make([]byte, 1)

	// Scan common I2C addresses (0x08 to 0x77)
	for addr := uint8(0x08); addr <= 0x77; addr++ {
		err := machine.I2C0.Tx(uint16(addr), nil, buf)
		if err == nil {
			devices = append(devices, addr)
		}
	}
	return devices
}

// formatBytes formats bytes as human-readable string
func formatBytes(b uint64) string {
	if b < 1024 {
		return strconv.FormatUint(b, 10) + "B"
	}
	kb := b / 1024
	return strconv.FormatUint(kb, 10) + "KB"
}
