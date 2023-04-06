package main

import (
	"machine"
	"machine/usb/hid/mouse"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

const DEADZONE = 15

func main() {

	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.I2C0_SCL_PIN, SDA: machine.I2C0_SDA_PIN})
	accel := lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()

	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})

	mouseDevice := mouse.Port()

	x, y, _ := accel.ReadRawAcceleration()
	for {
		x, y, _ = accel.ReadRawAcceleration()
		x = x / 250
		y = y / 250

		// If the inclination is too little, do not move the mouse
		if x < DEADZONE && x > -DEADZONE {
			x = 0
		}
		if y < DEADZONE && y > -DEADZONE {
			y = 0
		}

		mouseDevice.Move(-int(x), -int(y))

		if !btnA.Get() {
			mouseDevice.Click(mouse.Left)
		}
		if !btnB.Get() {
			mouseDevice.Click(mouse.Right)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
