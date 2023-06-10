package main

import (
	"machine/usb/hid/keyboard"
	"time"
)

func main() {

	time.Sleep(2000 * time.Millisecond)

	kb := keyboard.Port()

	// Use key super to open a text editor
	kb.Down(keyboard.KeyLeftGUI)
	time.Sleep(1000 * time.Millisecond)
	kb.Up(keyboard.KeyLeftGUI)
	time.Sleep(500 * time.Millisecond)
	kb.Write([]byte("text"))
	time.Sleep(1500 * time.Millisecond)
	kb.Press(keyboard.KeyEnter)
	time.Sleep(1000 * time.Millisecond)
	// Let them know we are hacking them MWAHAHAHA!
	kb.Write([]byte("Please wait while you are being hacked"))

	time.Sleep(2000 * time.Millisecond)

	// Execute command "xdg-open https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	// hid/keyboard currently only support US layout so it needs to be converted somehow
	// following "command" was tested in ISO International Spanish layout only
	kb.Down(keyboard.KeyF2)
	kb.Down(keyboard.KeyLeftAlt)
	time.Sleep(1000 * time.Millisecond)
	kb.Up(keyboard.KeyF2)
	kb.Up(keyboard.KeyLeftAlt)
	time.Sleep(1000 * time.Millisecond)
	kb.Write([]byte("xdg"))
	kb.Press(keyboard.KeypadMinus)
	kb.Write([]byte("open https>"))
	kb.Press(keyboard.KeypadSlash)
	kb.Press(keyboard.KeypadSlash)
	kb.Write([]byte("www.youtube.com"))
	kb.Press(keyboard.KeypadSlash)
	kb.Write([]byte("watch_v)dQw4w9WgXcQ"))
	kb.Press(keyboard.KeyEnter)

	// Turn the volume up
	time.Sleep(500 * time.Millisecond)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
	kb.Press(keyboard.KeyMediaVolumeInc)
}
