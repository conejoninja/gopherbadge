package main

import (
	"image/color"
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"
	"time"

	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"

	"tinygo.org/x/drivers/pca9548a"
	"tinygo.org/x/drivers/vl6180x"
)

const DEVICES = 5
const MIDICHANNEL = 10

func main() {
	machine.InitADC()
	usb.Product = "Dave GOhl"

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})
	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE
	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})
	display.FillScreen(color.RGBA{255, 255, 255, 255})
	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, "Dave")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (320-int16(w32))/2, 100, "Dave", color.RGBA{153, 51, 255, 255})
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, "GOhl")
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (320-int16(w32))/2, 180, "GOhl", color.RGBA{153, 51, 255, 255})

	sensor := machine.ADC{machine.ADC0}
	sensor.Configure(machine.ADCConfig{})

	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.I2C0_SCL_PIN, SDA: machine.I2C0_SDA_PIN})
	mux := pca9548a.New(machine.I2C0, pca9548a.Address)
	if !mux.IsConnected() {
		println("NO DEVICE DETECTED")
		return
	}

	port := mux.GetPortState()
	println("GET PORT", port)
	mux.DisablePort()
	mux.SetPort(0)
	port = mux.GetPortState()
	println("GET PORT", port)
	mux.SetPort(1)
	port = mux.GetPortState()
	println("GET PORT", port)

	dSensors := make([]vl6180x.Device, DEVICES)

	for i := uint8(0); i < DEVICES; i++ {
		mux.SetPort(i)
		time.Sleep(10 * time.Millisecond)
		dSensors[i] = vl6180x.New(machine.I2C0)
		connected := dSensors[i].Connected()
		if !connected {
			println("VL6180X device not found")
			return
		}
		println("VL6180X device found")
		dSensors[i].Configure(true)
	}
	time.Sleep(2000 * time.Millisecond)

	notes := []int{36, 50, 55, 58, 39}

	var speed uint16
	for {
		speed = sensor.Get()
		for i := uint8(0); i < DEVICES; i++ {
			mux.SetPort(i)
			if dSensors[i].Read() < 90 {
				midi.Midi.NoteOn(0, MIDICHANNEL, midi.Note(notes[i]), 64)
			}
		}
		time.Sleep(time.Duration(speed/100) * time.Millisecond)
	}
}
