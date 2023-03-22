package logos

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func GenerateLogoRGBAFile(filepath string) {
	colors := generateLogoRGBA(filepath)
	colorsStr := convertToString(colors)
	generateFile(colorsStr)
}

func generateLogoRGBA(filepath string) []color.RGBA {
	file, _ := os.Open(filepath)
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal("failed to decode image file", err)
	}

	colors := make([]color.RGBA, 0)

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			colors = append(colors, color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(255)})
		}
	}

	return colors
}

func convertToString(colors []color.RGBA) string {
	str := ""
	for _, c := range colors {
		str += fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
	}

	return str
}

func generateFile(colorsStr string) {
	err := os.WriteFile("logo.bin", []byte(colorsStr), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
