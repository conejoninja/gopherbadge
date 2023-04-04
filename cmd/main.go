package main

import (
	"flag"
	"fmt"

	"github.com/conejoninja/gopherbadge/cmd/logos"
)

const (
	tinygoLogo         = "./cmd/assets/tinygo.jpg"
	purpleHardwareLogo = "./cmd/assets/purpleHardware.jpg"
	gopherconEU        = "./cmd/assets/gopherconeu.jpg"
	gopherconIE        = "./cmd/assets/gopherconie.jpg"
	gopherconUK        = "./cmd/assets/gopherconuk.jpg"
	gopherconUS        = "./cmd/assets/gopherconus.jpg"
	wasmio             = "./cmd/assets/wasmio.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo)")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		fmt.Println("I do not have yet this conf in my catalog.")
		return
	}

	logos.GenerateLogoRGBAFile(logo)
}

func confs() map[string]string {
	return map[string]string{
		"tinygo":      tinygoLogo,
		"purple":      purpleHardwareLogo,
		"gopherconeu": gopherconEU,
		"gopherconie": gopherconIE,
		"gopherconuk": gopherconUK,
		"gopherconus": gopherconUS,
		"wasmio":      wasmio,
	}
}
