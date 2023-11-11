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
	golab              = "./cmd/assets/golab.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo)")
	filepath := flag.String("filepath", "", "Fullpath of the image, only when -conf=custom")
	flag.Parse()

	if *conf == "custom" {
		if *filepath == "" {
			fmt.Println("-filepath can not be empty if -conf=custom")
			return
		}
		logos.Resize(*filepath)
		logos.GenerateLogoRGBAFile(*filepath)
		return
	}
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
		"golab":       golab,
	}
}
