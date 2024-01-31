package main

import (
	"machine"
	"machine/usb"
	"machine/usb/adc/midi"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

const TICS = 5

func main() {
	usb.Product = "Eddie VAR Halen"

	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.I2C0_SCL_PIN, SDA: machine.I2C0_SDA_PIN})
	accel := lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()

	i := 0
	btns := []machine.Pin{machine.ADC0, machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7, machine.GPIO8}
	for i = 0; i < 6; i++ {
		btns[i].Configure(machine.PinConfig{Mode: machine.PinInput})
	}

	baseNote := 60
	oldBaseNote := 60
	midichannel := uint8(1)

	note := []bool{false, false, false, false, false, false}
	oldNote := []bool{false, false, false, false, false, false}

	zs := make([]int16, TICS)
	zz := 0
	zt := int16(0)

	x, _, z := accel.ReadRawAcceleration()
	tic := uint8(0)
	for {
		x, _, z = accel.ReadRawAcceleration()

		tic++
		if tic >= TICS {
			zs[zz] = z / 100
			zz++
			if zz >= TICS {
				zz = 0
			}
			zt = 0
			for i = 0; i < TICS; i++ {
				zt += zs[i]
			}
			zt = zt / TICS

			oldBaseNote = baseNote
			if zt < -80 {
				for i = 0; i < TICS; i++ {
					zs[i] = 0
				}
				baseNote++
				if baseNote > 90 {
					baseNote = 90
				}
			} else if zt > 80 {
				for i = 0; i < TICS; i++ {
					zs[i] = 0
				}
				baseNote--
				if baseNote < 37 {
					baseNote = 37
				}
			}
			tic = 0
		}

		for i = 0; i < 6; i++ {
			oldNote[i] = note[i]
			note[i] = btns[i].Get()
			if baseNote != oldBaseNote {
				midi.Midi.NoteOff(0, midichannel, midi.Note(oldBaseNote+i), 64)
			}
			if note[i] != oldNote[i] || baseNote != oldBaseNote {
				if note[i] {
					midi.Midi.NoteOn(0, midichannel, midi.Note(baseNote+i), 64)
				} else {
					midi.Midi.NoteOff(0, midichannel, midi.Note(oldBaseNote+i), 64)
				}
			}
		}

		pitch := uint16((x / 4) + 8192)

		midi.Midi.PitchBend(0, midichannel, pitch)
		time.Sleep(5 * time.Millisecond)
	}
}
